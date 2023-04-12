package main

import (
	"bufio"
	"os"

	"github.com/sirupsen/logrus"
)

// openAndRead opens a file located in /tmp directory.
// If the file does not exist the an error is prompted,
// otherwise the file is readed and printed on screen.
func openAndRead() {
	file, err := os.OpenFile("/tmp/tmp.file", os.O_RDONLY, 0700)

	if err != nil {
		logrus.Errorf("Error reading file: %v", err)
		file.Close()
	} else {
		logrus.Infof("Reading file: %v", file.Name())
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			logrus.Infof("line: %s\n", scanner.Text())
		}

		file.Close()
	}
}

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, FullTimestamp: true})
	openAndRead()

}
