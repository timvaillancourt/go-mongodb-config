package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	c := loadConfig(t)
	s := c.Storage
	assert.Equal(t, "wiredTiger", s.Engine, "'storage.engine' is not wiredTiger")
	assert.Equal(t, "/dev/null", s.DbPath, "'storage.dbPath' is not /dev/null")
	assert.NotNil(t, s.Journal, "'storage.journal' is nil")
	assert.True(t, s.Journal.Enabled, "'storage.journal.enabled' is not true")
	assert.NotNil(t, s.WiredTiger, "'storage.wiredTiger' is nil/malformed")
	assert.NotNil(t, s.WiredTiger.EngineConfig, "'storage.wiredTiger' is nil/malformed")
	assert.Equal(t, 1, s.WiredTiger.EngineConfig.CacheSizeGB, "'storage.engineConfig.cacheSizeGB' is not 1")
	assert.NotNil(t, s.RocksDB, "'storage.rocksdb' is nil")
	assert.True(t, s.RocksDB.Counters, "'storage.rocksdb.counters' is false")
	assert.NotNil(t, s.MMAPv1, "'storage.mmapv1' is nil")
	assert.True(t, s.MMAPv1.PreallocDataFiles, "'storage.mmapv1.preallocDataFiles' is not true")
}
