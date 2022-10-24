package main

import (
	"hotSwapLogger/system"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
}

func main() {

	router := gin.Default()

	router.GET("/set/log/:level", func(c *gin.Context) {
		level := c.Param("level")
		switch level {
		case "error":
			logrus.SetLevel(logrus.ErrorLevel)
		case "warn", "warning":
			logrus.SetLevel(logrus.WarnLevel)
		case "info":
			logrus.SetLevel(logrus.InfoLevel)
		case "debug":
			logrus.SetLevel(logrus.DebugLevel)
		}

		c.String(http.StatusOK, "Log level set to: %s", level)
	})

	router.GET("/get/log/level", func(c *gin.Context) {
		logrus.Error("This is error msg")
		logrus.Warn("This is warn msg")
		logrus.Info("This is info msg")
		logrus.Debug("This is debug msg")
		c.String(http.StatusOK, "Loglevel is: %s", logrus.GetLevel())
	})

	router.GET("/get/system/hostname", func(c *gin.Context) {
		c.String(http.StatusOK, "System hostname is: %s", system.GetHostname())
	})

	router.Run(":8080")

}
