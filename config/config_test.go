package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var testLoadUri = "https://raw.githubusercontent.com/timvaillancourt/go-mongodb-config/master/config/test/mongod.conf"

func testConfigFile(t *testing.T) string {
	_, filename, _, ok := runtime.Caller(0)
	if ok && filename != "" {
		return filepath.Join(filepath.Dir(filename), "test", "mongod.conf")
	}
	assert.FailNow(t, "Could not find path to go file")
	return ""
}

func loadConfig(t *testing.T) *Config {
	bytes, err := ioutil.ReadFile(testConfigFile(t))
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

func TestNewConfig(t *testing.T) {
	assert.Equal(t, &Config{}, New(), "config.New() did not return a &Config{} struct")
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
	_, err := Load("/does/not/exist")
	assert.Error(t, err, "config.Load() should return an error for paths that do not exist")

	config, err := Load(testConfigFile(t))
	assert.NoError(t, err, "Error running config.Load()")
	compareConfig := loadConfig(t)
	assert.Equal(t, compareConfig, config, "Test and config.Load() output differ")
}

func TestLoadUri(t *testing.T) {
	_, err := LoadUri("https://thisshouldfail:90000")
	assert.Error(t, err, "config.LoadUri() should fail due to bad uri")

	config, err := LoadUri(testLoadUri)
	assert.NoError(t, err, "Error running config.LoadUri()")
	assert.NotNil(t, config, "config.LoadUri() returned nil")
	assert.NotNil(t, config.Storage, "'storage' is nil")
	assert.Equal(t, "/dev/null", config.Storage.DbPath, "'storage.dbPath' is not /dev/null")
}

func TestWrite(t *testing.T) {
	config := loadConfig(t)
	assert.NotNil(t, config, "cannot load test config")

	testFile := "/tmp/go-mongodb-config.test.yaml"
	assert.NoError(t, config.Write(testFile), "config.Write() got error")
	defer os.Remove(testFile)

	config, err := Load(testFile)
	assert.NoError(t, err, "Error running config.Load()")
	assert.NotNil(t, config, "config.Load() returned nil")
	assert.NotNil(t, config.Storage, "'storage' is nil")
	assert.Equal(t, "/dev/null", config.Storage.DbPath, "'storage.dbPath' is not /dev/null")
}
