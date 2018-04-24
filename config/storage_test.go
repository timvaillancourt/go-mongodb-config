package config

import (
	"testing"
)

func TestStorage(t *testing.T) {
	c := loadConfig(t)
	s := c.Storage
	if s.Engine != "wiredTiger" {
		t.Errorf("'storage.engine' is %s not wiredTiger", s.Engine)
	} else if s.DbPath != "/dev/null" {
		t.Errorf("'storage.dbPath' is %s not /dev/null", s.DbPath)
	} else if s.Journal == nil {
		t.Error("'storage.journal' is nil")
	} else if !s.Journal.Enabled {
		t.Error("'storage.journal.enabled' is not true")
	} else if s.WiredTiger == nil || s.WiredTiger.EngineConfig == nil {
		t.Error("'storage.wiredTiger' is nil/malformed")
	} else if s.WiredTiger.EngineConfig.CacheSizeGB != 1 {
		t.Errorf("'storage.engineConfig.cacheSizeGB' is not 1, got %d", s.WiredTiger.EngineConfig.CacheSizeGB)
	} else if s.RocksDB == nil {
		t.Error("'storage.rocksdb' is nil")
	} else if s.RocksDB.Counters != true {
		t.Error("'storage.rocksdb.counters' is false")
	} else if s.MMAPv1 == nil {
		t.Error("'storage.mmapv1' is nil")
	} else if s.MMAPv1.PreallocDataFiles != true {
		t.Error("'storage.mmapv1.preallocDataFiles' is not true")
	}
}
