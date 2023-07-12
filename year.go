package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type year map[time.Month]map[int][]event

var _ fmt.Stringer = year{}

func (y year) next() (time.Month, int, []event) {
	now := time.Now()
	nextYear := now.AddDate(1, 0, 0)
	for t := now; t.Before(nextYear); t = t.AddDate(0, 0, 1) {
		if events := y.at(t); len(events) > 0 {
			return t.Month(), t.Day(), events
		}
	}
	return 0, 0, nil
}

func (y year) at(t time.Time) []event {
	return y[t.Month()][t.Day()]
}

func (y year) today() []event {
	return y.at(time.Now())
}

func (e year) String() string {
	var sb strings.Builder

	for m := time.January; m <= time.December; m++ {
		month, ok := e[m]
		if !ok {
			continue
		}
		for d := 1; d <= 31; d++ {
			day, ok := month[d]
			if !ok {
				continue
			}
			for _, event := range day {
				sb.WriteString(fmt.Sprintf(formatListLine, d, m, event))
			}
		}
	}
	return sb.String()
}

func readYear(path string) (year, error) {
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
	return year{
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