package config

import(
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// 'SetParameter' is a map of string for now
// https://docs.mongodb.com/manual/reference/configuration-options/
type Config struct {
	SetParameter		map[string]string	`yaml:"setParameter,omitempty"`
	AuditLog		*AuditLog		`yaml:"auditLog,omitempty"`
	Net			*Net			`yaml:"net,omitempty"`
	OperationProfiling	*OperationProfiling	`yaml:"operationProfiling,omitempty"`
	ProcessManagement	*ProcessManagement	`yaml:"processManagement,omitempty"`
	Replication		*Replication		`yaml:"replication,omitempty"`
	Security		*Security		`yaml:"security,omitempty"`
	Sharding		*Sharding		`yaml:"sharding,omitempty"`
	Snmp			*Snmp			`yaml:"snmp,omitempty"`
	Storage			*Storage		`yaml:"storage,omitempty"`
	SystemLog		*SystemLog		`yaml:"systemLog,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#auditlog-options
type AuditLog struct {
	Destination	string	`yaml:"destination,omitempty"`
	Filter		string	`yaml:"filter,omitempty"`
	Format		string	`yaml:"format,omitempty"`
	Path		string	`yaml:"path,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#operationprofiling-options
type OperationProfiling struct {
	Mode			string	`yaml:"mode,omitempty"`
	SlowOpThresholdMs	int	`yaml:"slowOpThresholdMs,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#processmanagement-options
type ProcessManagement struct {
	Fork		bool	`yaml:"fork,omitempty"`
	PidFilePath	string	`yaml:"pidFilePath,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#replication-options
type Replication struct {
	EnableMajorityReadConcern	bool	`yaml:"enableMajorityReadConcern,omitempty"`
	LocalPingThresholdMs		int	`yaml:"localPingThresholdMs,omitempty"`
	OplogSizeMB			int	`yaml:"oplogSizeMb,omitempty"`
	ReplSetName			string	`yaml:"replSetName,omitempty"`
	SecondaryIndexPrefetch		string	`yaml:"secondaryIndexPrefetch,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#sharding-options
type Sharding struct {
	ArchiveMovedChunks	bool	`yaml:"archiveMovedChunks"`
	ClusterRole		string	`yaml:"clusterRole,omitempty"`
	ConfigDB		string	`yaml:"configDB,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#snmp-options
type Snmp struct {
	Master		bool	`yaml:"master,omitempty"`
	Subagent	bool	`yaml:"subagent,omitempty"`
}

func New() *Config {
	return &Config{}
}

func Load(filePath string) (*Config, error) {
	config := &Config{}
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(raw, config)
	return config, err
}

func (config *Config) Write(filePath string) error {
	yaml, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, yaml, 0644)
}
