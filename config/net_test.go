package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNet(t *testing.T) {
	c := loadConfig(t)
	assert.NotEmpty(t, c.Net.BindIp, "'net.bindIp' is an empty string")
	assert.NotZero(t, c.Net.Port, "'net.port' is not greater than 1")
}
