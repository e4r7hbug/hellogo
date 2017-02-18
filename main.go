/*
Package hellogo represents a fun Go toy program.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
)

var (
	log = logrus.New()

	requestBad   = "https://this.will.not.work"
	requestNotOk = "https://httpbin.org/status/404"
	requestGood  = "https://ifconfig.co"
)

func init() {
	log.Level = logrus.DebugLevel
}

func main() {
	fmt.Printf("Hellow!\n")

	if log.Level < logrus.DebugLevel {
		log.WithField("level", log.Level).Info("Debug logging turned off.")
	}

	Get(requestBad)
	Get(requestNotOk)
	Get(requestGood)

	Library1Printer()
}

func Get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err).Error("Get failed.")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.WithFields(logrus.Fields{
			"Status":  resp.Status,
			"Request": resp.Request.URL,
		}).Warning("Request not OK.")
		return
	}

	log.WithFields(logrus.Fields{
		"Status":  resp.Status,
		"Request": resp.Request.URL,
	}).Info("Get succeeded.")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Panic("Body not good.")
	}
	log.Debug(fmt.Sprintf("Response body:\n%s", body))
}
