package config

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

// 'SetParameter' is a map of string for now
// https://docs.mongodb.com/manual/reference/configuration-options/
type Config struct {
	SetParameter       map[string]string   `yaml:"setParameter,omitempty" json:"setParameter,omitempty" bson:"setParameter,omitempty"`
	AuditLog           *AuditLog           `yaml:"auditLog,omitempty" json:"auditLog,omitempty" bson:"auditLog,omitempty"`
	Net                *Net                `yaml:"net,omitempty" json:"net,omitempty" bson:"net,omitempty"`
	OperationProfiling *OperationProfiling `yaml:"operationProfiling,omitempty" json:"operationProfiling,omitempty" bson:"operationProfiling,omitempty"`
	ProcessManagement  *ProcessManagement  `yaml:"processManagement,omitempty" json:"processManagement,omitempty" bson:"processManagement,omitempty"`
	Replication        *Replication        `yaml:"replication,omitempty" json:"replication,omitempty" bson:"replication,omitempty"`
	Security           *Security           `yaml:"security,omitempty" json:"security,omitempty" bson:"security,omitempty"`
	Sharding           *Sharding           `yaml:"sharding,omitempty" json:"sharding,omitempty" bson:"sharding,omitempty"`
	Snmp               *Snmp               `yaml:"snmp,omitempty" json:"snmp,omitempty" bson:"snmp,omitempty"`
	Storage            *Storage            `yaml:"storage,omitempty" json:"storage,omitempty" bson:"storage,omitempty"`
	SystemLog          *SystemLog          `yaml:"systemLog,omitempty" json:"systemLog,omitempty" bson:"systemLog,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#auditlog-options
type AuditLog struct {
	Destination string `yaml:"destination,omitempty" json:"destination,omitempty" bson:"destination,omitempty"`
	Filter      string `yaml:"filter,omitempty" json:"filter,omitempty" bson:"filter,omitempty"`
	Format      string `yaml:"format,omitempty" json:"format,omitempty" bson:"format,omitempty"`
	Path        string `yaml:"path,omitempty" json:"path,omitempty" bson:"path,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#operationprofiling-options
type OperationProfiling struct {
	Mode              string `yaml:"mode,omitempty" json:"mode,omitempty" bson:"mode,omitempty"`
	SlowOpThresholdMs int    `yaml:"slowOpThresholdMs,omitempty" json:"slowOpThresholdMs,omitempty" bson:"slowOpThresholdMs,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#processmanagement-options
type ProcessManagement struct {
	Fork        bool   `yaml:"fork,omitempty" json:"fork,omitempty" bson:"fork,omitempty"`
	PidFilePath string `yaml:"pidFilePath,omitempty" json:"pidFilePath,omitempty" bson:"pidFilePath,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#replication-options
type Replication struct {
	EnableMajorityReadConcern bool   `yaml:"enableMajorityReadConcern,omitempty" json:"enableMajorityReadConcern,omitempty" bson:"enableMajorityReadConcern,omitempty"`
	LocalPingThresholdMs      int    `yaml:"localPingThresholdMs,omitempty" json:"localPingThresholdMs,omitempty" bson:"localPingThresholdMs,omitempty"`
	OplogSizeMB               int    `yaml:"oplogSizeMb,omitempty" json:"oplogSizeMb,omitempty" bson:"oplogSizeMb,omitempty"`
	ReplSetName               string `yaml:"replSetName,omitempty" json:"replSetName,omitempty" bson:"replSet,omitempty"`
	SecondaryIndexPrefetch    string `yaml:"secondaryIndexPrefetch,omitempty" json:"secondaryIndexPrefetch,omitempty" bson:"secondaryIndexPrefetch,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#sharding-options
type Sharding struct {
	ArchiveMovedChunks bool   `yaml:"archiveMovedChunks" json:"archiveMovedChunks" bson:"archiveMovedChunks"`
	ClusterRole        string `yaml:"clusterRole,omitempty" json:"clusterRole,omitempty" bson:"clusterRole,omitempty"`
	ConfigDB           string `yaml:"configDB,omitempty" json:"configDB,omitempty" bson:"configDB,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#snmp-options
type Snmp struct {
	Master   bool `yaml:"master,omitempty" json:"master,omitempty" bson:"master,omitempty"`
	Subagent bool `yaml:"subagent,omitempty" json:"subagent,omitempty" bson:"subagent,omitempty"`
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

func LoadUri(httpUri string) (*Config, error) {
	resp, err := http.Get(httpUri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	return config, err
}

func (config *Config) String() string {
	yaml, err := yaml.Marshal(config)
	if err != nil {
		return ""
	}
	return string(yaml)
}

func (config *Config) Write(filePath string) error {
	yaml, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, yaml, 0644)
}
