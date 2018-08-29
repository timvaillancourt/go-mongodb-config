package config

// https://docs.mongodb.com/manual/reference/configuration-options/#storage-options
type Storage struct {
	DbPath          string             `yaml:"dbPath,omitempty" json:"dbPath,omitempty"`
	DirectoryPerDB  bool               `yaml:"directorPerDB,omitempty" json:"directorPerDB,omitempty"`
	Engine          string             `yaml:"engine,omitempty" json:"engine,omitempty"`
	IndexBuildRetry bool               `yaml:"indexBuildRetry,omitempty" json:"indexBuildRetry,omitempty"`
	RepairPath      string             `yaml:"repairPath,omitempty" json:"repairPath,omitempty"`
	SyncPeriodSecs  int                `yaml:"syncPeriodSecs,omitempty" json:"syncPeriodSecs,omitempty"`
	Journal         *StorageJournal    `yaml:"journal,omitempty" json:"journal,omitempty"`
	MMAPv1          *StorageMMAPv1     `yaml:"mmapv1,omitempty" json:"mmapv1,omitempty"`
	RocksDB         *StorageRocksDB    `yaml:"rocksdb,omitempty" json:"rocksdb,omitempty"`
	InMemory        *StorageWiredTiger `yaml:"inMemory,omitempty" json:"inMemory,omitempty"`
	WiredTiger      *StorageWiredTiger `yaml:"wiredTiger,omitempty" json:"wiredTiger,omitempty"`
}

type StorageJournal struct {
	CommitIntervalMs int  `yaml:"commitIntervalMs,omitempty" json:"commitIntervalMs,omitempty"`
	Enabled          bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`
}

type StorageRocksDB struct {
	CacheSizeGB       int    `yaml:"cacheSizeGB,omitempty" json:"cacheSizeGB,omitempty" json:"cacheSizeGB,omitempty"`
	Compression       string `yaml:"compression,omitempty" json:"compression,omitempty"`
	Counters          bool   `yaml:"counters,omitempty" json:"counters,omitempty"`
	CrashSafeCounters bool   `yaml:"crashSafeCounters,omitempty" json:"crashSafeCounters,omitempty"`
	MaxWriteMBPerSec  int    `yaml:"maxWriteMBPerSec,omitempty" json:"maxWriteMBPerSec,omitempty"`
	SingleDeleteIndex bool   `yaml:"singleDeleteIndex,omitempty" json:"singleDeleteIndex,omitempty"`
}

type StorageMMAPv1Quota struct {
	Enforced      bool `yaml:"enforced,omitempty" json:"enforced,omitempty"`
	MaxFilesPerDB int  `yaml:"maxFilesPerDB,omitempty" json:"maxFilesPerDB,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#storage-mmapv1-options
type StorageMMAPv1 struct {
	NsSize            int                 `yaml:"nsSize,omitempty" json:"nsSize,omitempty"`
	PreallocDataFiles bool                `yaml:"preallocDataFiles,omitempty" json:"preallocDataFiles,omitempty"`
	Quota             *StorageMMAPv1Quota `yaml:"quota,omitempty" json:"quota,omitempty"`
	Smallfiles        bool                `yaml:"smallfiles,omitempty" json:"smallfiles,omitempty"`
}

type StorageWiredTigerEngineConfig struct {
	CacheSizeGB            float64 `yaml:"cacheSizeGB,omitempty" json:"cacheSizeGB,omitempty" json:"cacheSizeGB,omitempty"`
	DirectoryForIndexes    bool    `yaml:"directoryForIndexes,omitempty" json:"directoryForIndexes,omitempty"`
	JournalCompressor      string  `yaml:"journalCompressor,omitempty" json:"journalCompressor,omitempty"`
	StatisticsLogDelaySecs int     `yaml:"statisticsLogDelaySecs,omitempty" json:"statisticsLogDelaySecs,omitempty"`
}

type StorageWiredTigerCollectionConfig struct {
	BlockCompression bool `yaml:"blockCompression,omitempty" json:"blockCompression,omitempty"`
}

type StorageWiredTigerIndexConfig struct {
	PrefixCompression bool `yaml:"prefixCompression,omitempty" json:"prefixCompression,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#storage-wiredtiger-options
type StorageWiredTiger struct {
	EngineConfig     *StorageWiredTigerEngineConfig     `yaml:"engineConfig,omitempty" json:"engineConfig,omitempty"`
	CollectionConfig *StorageWiredTigerCollectionConfig `yaml:"collectionConfig,omitempty" json:"collectionConfig,omitempty"`
	IndexConfig      *StorageWiredTigerIndexConfig      `yaml:"indexConfig,omitempty" json:"indexConfig,omitempty"`
}
