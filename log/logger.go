package log

import (
	"go.uber.org/zap"
	"encoding/json"
)

var logger *zap.Logger

const logConfig = `{
	  "level": "debug",
	  "development": false,
	  "encoding": "json",
	  "outputPaths": ["zap.log"],
	  "errorOutputPaths": ["err.log"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`

func NewLogger() *zap.Logger {
	rawJSON := []byte(logConfig)

	var (
		cfg zap.Config
		err error
	)
	if err = json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
