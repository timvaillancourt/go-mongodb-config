package config

import (
	"testing"
)

func TestNet(t *testing.T) {
	c := loadConfig(t)
	n := c.Net
	if n.BindIp == "" {
		t.Error("'net.bindIp' is an empty string")
	} else if n.Port < 1 {
		t.Error("'net.port' is not greater than 1")
	}
}
