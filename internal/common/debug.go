package common

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func DebugStruct(value any) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		logrus.Errorf("Error marshalling config for debug: %v", err)
		return
	}
	logrus.Debugf("%s", string(data))
}