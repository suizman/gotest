package system

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetHostname() string {
	hn, _ := os.Hostname()
	logrus.Debugf("Getting system hostname: %s", hn)
	return hn
}
