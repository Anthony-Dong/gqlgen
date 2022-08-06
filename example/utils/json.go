package utils

import (
	"encoding/json"
	"log"
	"time"
)

func ToJsonString(v interface{}) string {
	if v == nil {
		return ""
	}
	data, _ := json.Marshal(v)
	return string(data)
}

func Info(format string, v ...interface{}) {
	log.Printf(format+"\n", v...)
}

func TimePtr(t time.Time) *time.Time {
	return &t
}
