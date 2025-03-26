package utils

import "encoding/json"

type Log struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Context   string `json:"context"`
}

func (l *Log) ConvertLogToByte() (logData []byte, err error) {
	logData, err = json.Marshal(l)
	if err != nil {
		return nil, err
	}

	return logData, nil
}
