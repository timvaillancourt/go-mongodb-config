package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemLog(t *testing.T) {
	c := loadConfig(t)
	s := c.SystemLog
	assert.Equal(t, "file", s.Destination, "'systemLog.destination' is not 'file'")
	assert.True(t, s.LogAppend, "'systemLog.logAppend' is not true")
}
