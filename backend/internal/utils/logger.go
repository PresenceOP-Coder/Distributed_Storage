package utils

import (
	"encoding/json"
	"log"
)

func Infof(format string, args ...any) {
	log.Printf(format, args...)
}

func Errorf(format string, args ...any) {
	log.Printf(format, args...)
}

func LogJSON(fields map[string]any) {
	payload, err := json.Marshal(fields)
	if err != nil {
		log.Printf("[ERROR] event=log_json err=%v", err)
		return
	}
	log.Printf("%s", payload)
}
