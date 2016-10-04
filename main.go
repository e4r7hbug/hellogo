/*
Package hellogo represents a fun Go toy program.
*/
package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
)

var (
	log = logrus.New()
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	log.Level = logrus.DebugLevel
}

func main() {
	fmt.Printf("Hellow!\n")

	log.Debug("This is DEBUG.")
	log.Info("This is INFO.")
	log.Warning("This is WARNING.")
	log.Error("This is ERROR.")

	if false {
		log.Panic("This is PANIC.")
		log.Fatal("This is FATAL.")
	}

	Library1Printer()
}
