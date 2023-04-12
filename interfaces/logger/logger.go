package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "info"
	}

	ll, err := logrus.ParseLevel(lvl)

	if err != nil {
		logrus.Errorf("Error: %v. Setting log level to DEBUG", err)
		ll = logrus.DebugLevel
	} else {
		logrus.Infof("Setting log level to: %v", lvl)
		logrus.SetLevel(ll)
	}

}
