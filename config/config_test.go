package config

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func testConfigFile() string {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		return filepath.Join(filepath.Dir(filename), "mongod.conf")
	}
	return ""
}

func loadConfig(t *testing.T) *Config {
	bytes, err := ioutil.ReadFile(testConfigFile())
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	c := &Config{}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	return c
}

func TestConfig(t *testing.T) {
	c := loadConfig(t)
	if c == nil {
		assert.FailNow(t, "config is nil")
	}
	assert.NotNil(t, c.Net, "'net' is nil")
	assert.NotNil(t, c.OperationProfiling, "'operationProfiling' is nil")
	assert.NotNil(t, c.ProcessManagement, "'processManagement' is nil")
	assert.NotNil(t, c.Replication, "'replication' is nil")
	assert.NotNil(t, c.Security, "'security' is nil")
	assert.NotNil(t, c.Sharding, "'sharding' is nil")
	assert.NotNil(t, c.Storage, "'storage' is nil")
	assert.NotNil(t, c.SystemLog, "'systemLog' is nil")
}

func TestLoadConfig(t *testing.T) {
	config, err := Load(testConfigFile())
	assert.NoError(t, err, "Error running config.Load()")
	compareConfig := loadConfig(t)
	assert.Equal(t, compareConfig, config, "Test and config.Load() output differ")
}
