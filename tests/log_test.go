package tests

import (
	"encoding/json"
	"go.uber.org/zap"
	"testing"
)

func Test_logger(t *testing.T) {
	rawJSON := []byte(`{
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
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		t.Fatal(err.Error())
	}
	logger, err := cfg.Build()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")

	url := "loopring.org"
	for i := 1; i < 100000; i++ {
		logger.Info("saving number", zap.String("url", url), zap.Int("attempt", i))
	}
}
