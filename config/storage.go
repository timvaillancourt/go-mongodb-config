package config

// https://docs.mongodb.com/manual/reference/configuration-options/#storage-options
type Storage struct {
	DbPath		string			`yaml:"dbPath,omitempty"`
	DirectoryPerDB	bool			`yaml:"directorPerDB,omitempty"`
	Engine		string			`yaml:"engine,omitempty"`
	IndexBuildRetry	bool			`yaml:"indexBuildRetry,omitempty"`
	RepairPath	string			`yaml:"repairPath,omitempty"`
	SyncPeriodSecs	int			`yaml:"syncPeriodSecs,omitempty"`
	Journal		*StorageJournal		`yaml:"journal,omitempty"`
	MMAPv1		*StorageMMAPv1		`yaml:"mmapv1,omitempty"`
	RocksDB		*StorageRocksDB		`yaml:"rocksdb,omitempty"`
	InMemory	*StorageWiredTiger	`yaml:"inMemory,omitempty"`
	WiredTiger	*StorageWiredTiger	`yaml:"wiredTiger,omitempty"`
}

type StorageJournal struct {
	CommitIntervalMs	int	`yaml:"commitIntervalMs,omitempty"`
	Enabled			bool	`yaml:"enabled,omitempty"`
}

type StorageRocksDB struct {
	CacheSizeGB		int	`yaml:"cacheSizeGB,omitempty"`
	Compression		string	`yaml:"compression,omitempty"`
	Counters		bool	`yaml:"counters,omitempty"`
	CrashSafeCounters	bool	`yaml:"crashSafeCounters,omitempty"`
	MaxWriteMBPerSec	int	`yaml:"maxWriteMBPerSec,omitempty"`
	SingleDeleteIndex	bool	`yaml:"singleDeleteIndex,omitempty"`
}

type StorageMMAPv1Quota struct {
	Enforced	bool	`yaml:"enforced,omitempty"`
	MaxFilesPerDB	int	`yaml:"maxFilesPerDB,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#storage-mmapv1-options
type StorageMMAPv1 struct {
	NsSize			int			`yaml:"nsSize,omitempty"`
	PreallocDataFiles	bool			`yaml:"preallocDataFiles,omitempty"`
	Quota			*StorageMMAPv1Quota	`yaml:"quota,omitempty"`
	Smallfiles		bool			`yaml:"smallfiles,omitempty"`
}

type StorageWiredTigerEngineConfig struct {
	CacheSizeGB		int	`yaml:"cacheSizeGB,omitempty"`
	DirectoryForIndexes	bool	`yaml:"directoryForIndexes,omitempty"`
	JournalCompressor	string	`yaml:"journalCompressor,omitempty"`
	StatisticsLogDelaySecs	int	`yaml:"statisticsLogDelaySecs,omitempty"`
}

type StorageWiredTigerCollectionConfig struct {
	BlockCompression	bool	`yaml:"blockCompression,omitempty"`
}

type StorageWiredTigerIndexConfig struct {
	PrefixCompression	bool	`yaml:"prefixCompression,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#storage-wiredtiger-options
type StorageWiredTiger struct {
	EngineConfig		*StorageWiredTigerEngineConfig		`yaml:"engineConfig,omitempty"`
	CollectionConfig	*StorageWiredTigerCollectionConfig	`yaml:"collectionConfig,omitempty"`
	IndexConfig		*StorageWiredTigerIndexConfig		`yaml:"indexConfig,omitempty"`
}
