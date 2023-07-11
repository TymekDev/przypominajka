package main

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type events map[time.Month]map[int][]event

func (e events) at(t time.Time) []event {
	return e[t.Month()][t.Day()]
}

func (e events) today() []event {
	return e.at(time.Now())
}

func readEvents(path string) (events, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config struct {
		January   map[int][]event `yaml:"january"`
		February  map[int][]event `yaml:"february"`
		March     map[int][]event `yaml:"march"`
		April     map[int][]event `yaml:"april"`
		May       map[int][]event `yaml:"may"`
		June      map[int][]event `yaml:"june"`
		July      map[int][]event `yaml:"july"`
		August    map[int][]event `yaml:"august"`
		September map[int][]event `yaml:"september"`
		October   map[int][]event `yaml:"october"`
		November  map[int][]event `yaml:"november"`
		December  map[int][]event `yaml:"december"`
	}
	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return events{
		time.January:   config.January,
		time.February:  config.February,
		time.March:     config.March,
		time.April:     config.April,
		time.May:       config.May,
		time.June:      config.June,
		time.July:      config.July,
		time.August:    config.August,
		time.September: config.September,
		time.October:   config.October,
		time.November:  config.November,
		time.December:  config.December,
	}, nil
}