package logrusdemo

import (
	"os"

	"github.com/sirupsen/logrus"
)

func WriteLog() {
	// jsonLog()
	// logrusInstance()
}

func jsonLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("A walrus appears")
	logrus.WithField("animal", "walrus").Info("A walrus appears")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}

func logrusInstance() {
	var log = logrus.New()

	log.Out = os.Stdout

  log.SetLevel(logrus.DebugLevel)
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}

