package config

type SystemLogComponent struct {
	Verbosity int `yaml:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty"`
}

type SystemLogComponentStorage struct {
	Journal   *SystemLogComponent `yaml:"journal,omitempty" json:"journal,omitempty"`
	Verbosity int                 `yaml:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#systemlog-component-options
type SystemLogComponents struct {
	AccessControl *SystemLogComponent        `yaml:"accessControl,omitempty" json:"accessControl,omitempty"`
	Command       *SystemLogComponent        `yaml:"command,omitempty" json:"command,omitempty"`
	Control       *SystemLogComponent        `yaml:"control,omitempty" json:"control,omitempty"`
	Ftdc          *SystemLogComponent        `yaml:"ftdc,omitempty" json:"ftdc,omitempty"`
	Geo           *SystemLogComponent        `yaml:"geo,omitempty" json:"geo,omitempty"`
	Index         *SystemLogComponent        `yaml:"index,omitempty" json:"index,omitempty"`
	Network       *SystemLogComponent        `yaml:"network,omitempty" json:"network,omitempty"`
	Query         *SystemLogComponent        `yaml:"query,omitempty" json:"query,omitempty"`
	Replication   *SystemLogComponent        `yaml:"replication,omitempty" json:"replication,omitempty"`
	Sharding      *SystemLogComponent        `yaml:"sharding,omitempty" json:"sharding,omitempty"`
	Storage       *SystemLogComponentStorage `yaml:"storage,omitempty" json:"storage,omitempty"`
	Write         *SystemLogComponent        `yaml:"write,omitempty" json:"write,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#systemlog-options
type SystemLog struct {
	Destination        string `yaml:"destination,omitempty" json:"destination,omitempty"`
	LogAppend          bool   `yaml:"logAppend,omitempty" json:"logAppend,omitempty"`
	LogRotate          string `yaml:"logRotate,omitempty" json:"logRotate,omitempty"`
	Path               string `yaml:"path,omitempty" json:"path,omitempty"`
	Quiet              bool   `yaml:"quiet,omitempty" json:"quiet,omitempty"`
	TimeStampFormat    string `yaml:"timeStampFormat,omitempty" json:"timeStampFormat,omitempty"`
	TraceAllExceptions bool   `yaml:"traceAllExceptions,omitempty" json:"traceAllExceptions,omitempty"`
	SyslogFacility     string `yaml:"syslogFacility,omitempty" json:"syslogFacility,omitempty"`
	Verbosity          int    `yaml:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty" json:"verbosity,omitempty"`
}
