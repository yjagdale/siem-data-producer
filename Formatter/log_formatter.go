package Formatter

import (
	log "github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/models/override_model"
	"math/rand"
	"strings"
	"time"
)

func FormatLog(line string) string {
	for key, value := range override_model.OverrideValues {
		if strings.Contains(line, key) {
			line = strings.ReplaceAll(line, key, getRandomValue(value))
		}
	}
	return line
}

func getRandomValue(vals []string) string {
	log.Debugln("Total items", len(vals))
	index := rand.Intn(len(vals))
	log.Debugln("Selecting index", index)
	if strings.HasPrefix(vals[index], "FUNCTION") {
		return getValueForSpecialFunction(vals[index])
	}
	return vals[index]
}

func getValueForSpecialFunction(logLine string) string {
	output := strings.Split(logLine, "::")
	if len(output) == 3 {
		switch output[1] {
		case "DATE":
			return time.Now().UTC().Format(output[2])
		case "IP":
			tmp, err := Hosts(output[2])
			if err == nil {
				return tmp[rand.Intn(len(tmp))]
			}
		}
	}
	return logLine
}
