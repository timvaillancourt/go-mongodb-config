package config

import (
	"testing"
)

func TestSystemLog(t *testing.T) {
	c := loadConfig(t)
	s := c.SystemLog
	if s.Destination != "file" {
		t.Errorf("'systemLog.destination' is %s, not file", s.Destination)
	} else if s.LogAppend != true {
		t.Errorf("'systemLog.logAppend' is %s not true", s.LogAppend)
	}
}
