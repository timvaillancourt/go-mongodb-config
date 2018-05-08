package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecurity(t *testing.T) {
	c := loadConfig(t)
	assert.Equal(t, "/etc/mongod.key", c.Security.KeyFile, "'security.keyFile' is not /etc/mongod.key")
	assert.Equal(t, "enabled", c.Security.Authorization, "'security.authorization' is not enabled")
}
