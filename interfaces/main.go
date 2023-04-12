package main

import (
	"interfaces/logger"

	"github.com/sirupsen/logrus"
)

type Store interface {
	Add(input []string)
	Get()
	DelLast()
	Lookup(search string)
}

type Storage struct {
	name []string
}

type Greeter interface {
	Greet() string
}

func (store *Storage) Add(input []string) {
	logrus.Debugf("Adding new entry to storage: %v", input)
	store.name = append(store.name, input...)
}

func (store *Storage) Get() {
	logrus.Infof("Storage content: %v", store.name)
}

func (store *Storage) DelLast() {
	store.name = store.name[:len(store.name)-1]
}

func (store *Storage) Lookup(search string) {

	for key, value := range store.name {
		if value == search {
			logrus.Infof("Match found at index: %v for lookup %v", key, value)
			return
		}
	}

	logrus.Infof("Unable to find records in storage for: %v", search)
}

func main() {

	logger.SetupLogger()

	var s Store = &Storage{}
	s.Add([]string{"Upertencio"})
	s.Add([]string{"Baticilo"})
	s.Add([]string{"Megicula"})
	s.Get()
	s.DelLast()
	s.Get()
	s.Lookup("Upertencio")

}
