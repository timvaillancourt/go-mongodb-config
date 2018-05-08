package config

type SystemLogComponent struct {
	Verbosity int `yaml:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty" bson:"verbosity,omitempty" bson:"verbosity,omitempty" bson:"verbosity,omitempty"`
}

type SystemLogComponentStorage struct {
	Journal   *SystemLogComponent `yaml:"journal,omitempty" json:"journal,omitempty" bson:"journal,omitempty"`
	Verbosity int                 `yaml:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty" bson:"verbosity,omitempty" bson:"verbosity,omitempty" bson:"verbosity,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#systemlog-component-options
type SystemLogComponents struct {
	AccessControl *SystemLogComponent        `yaml:"accessControl,omitempty" json:"accessControl,omitempty" bson:"accessControl,omitempty"`
	Command       *SystemLogComponent        `yaml:"command,omitempty" json:"command,omitempty" bson:"command,omitempty"`
	Control       *SystemLogComponent        `yaml:"control,omitempty" json:"control,omitempty" bson:"control,omitempty"`
	Ftdc          *SystemLogComponent        `yaml:"ftdc,omitempty" json:"ftdc,omitempty" bson:"ftdc,omitempty"`
	Geo           *SystemLogComponent        `yaml:"geo,omitempty" json:"geo,omitempty" bson:"geo,omitempty"`
	Index         *SystemLogComponent        `yaml:"index,omitempty" json:"index,omitempty" bson:"index,omitempty"`
	Network       *SystemLogComponent        `yaml:"network,omitempty" json:"network,omitempty" bson:"network,omitempty"`
	Query         *SystemLogComponent        `yaml:"query,omitempty" json:"query,omitempty" bson:"query,omitempty"`
	Replication   *SystemLogComponent        `yaml:"replication,omitempty" json:"replication,omitempty" bson:"replication,omitempty"`
	Sharding      *SystemLogComponent        `yaml:"sharding,omitempty" json:"sharding,omitempty" bson:"sharding,omitempty"`
	Storage       *SystemLogComponentStorage `yaml:"storage,omitempty" json:"storage,omitempty" bson:"storage,omitempty"`
	Write         *SystemLogComponent        `yaml:"write,omitempty" json:"write,omitempty" bson:"write,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#systemlog-options
type SystemLog struct {
	Destination        string `yaml:"destination,omitempty" json:"destination,omitempty" bson:"destination,omitempty"`
	LogAppend          bool   `yaml:"logAppend,omitempty" json:"logAppend,omitempty" bson:"logAppend,omitempty"`
	LogRotate          string `yaml:"logRotate,omitempty" json:"logRotate,omitempty" bson:"logRotate,omitempty"`
	Path               string `yaml:"path,omitempty" json:"path,omitempty" bson:"path,omitempty"`
	Quiet              bool   `yaml:"quiet,omitempty" json:"quiet,omitempty" bson:"quiet,omitempty"`
	TimeStampFormat    string `yaml:"timeStampFormat,omitempty" json:"timeStampFormat,omitempty" bson:"timeStampFormat,omitempty"`
	TraceAllExceptions bool   `yaml:"traceAllExceptions,omitempty" json:"traceAllExceptions,omitempty" bson:"traceAllExceptions,omitempty"`
	SyslogFacility     string `yaml:"syslogFacility,omitempty" json:"syslogFacility,omitempty" bson:"syslogFacility,omitempty"`
	Verbosity          int    `yaml:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty" bson:"verbosity,omitempty" bson:"verbosity,omitempty" bson:"verbosity,omitempty"`
}
