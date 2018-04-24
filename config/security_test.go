package config

import (
	"testing"
)

func TestSecurity(t *testing.T) {
	c := loadConfig(t)
	s := c.Security
	if s.KeyFile != "/etc/mongod.key" {
		t.Errorf("'security.keyFile' is %s, not /etc/mongod.key", s.KeyFile)
	} else if s.Authorization != "enabled" {
		t.Errorf("'security.authorization' is %s, not enabled")
	}
}
