package config

import (
	"io/ioutil"
	"testing"

	"github.com/go-test/deep"
	"gopkg.in/yaml.v2"
)

const testConfigFile = "./mongod.conf"

func loadConfig(t *testing.T) *Config {
	bytes, err := ioutil.ReadFile(testConfigFile)
	if err != nil {
		t.Errorf("Error loading config file %s: %s", testConfigFile, err)
		return nil
	}
	c := &Config{}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		t.Errorf("Error parsing config file %s: %s", testConfigFile, err)
		return nil
	}
	return c
}

func TestConfig(t *testing.T) {
	c := loadConfig(t)
	if c == nil {
		t.Error("config is nil")
		return
	} else if c.Net == nil {
		t.Error("'net' is nil")
	} else if c.OperationProfiling == nil {
		t.Error("'operationProfiling' is nil")
	} else if c.ProcessManagement == nil {
		t.Error("'processManagement' is nil")
	} else if c.Replication == nil {
		t.Error("'replication' is nil")
	} else if c.Security == nil {
		t.Error("'security' is nil")
	} else if c.Sharding == nil {
		t.Error("'sharding' is nil")
	} else if c.Storage == nil {
		t.Error("'storage' is nil")
	} else if c.SystemLog == nil {
		t.Error("'systemLog' is nil")
	}
}

func TestLoadConfig(t *testing.T) {
	config, err := Load(testConfigFile)
	if err != nil {
		t.Errorf("Error running config.Load(): %s", err)
		return
	}
	compareConfig := loadConfig(t)
	if diff := deep.Equal(config, compareConfig); diff != nil {
		t.Errorf("Test and config.Load() output differ: %s", diff)
	}
}
