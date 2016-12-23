package config

type SystemLogComponent struct {
	Verbosity	int	`yaml:"verbosity,omitempty"`
}

type SystemLogComponentStorage struct {
	Journal		*SystemLogComponent	`yaml:"journal,omitempty"`
	Verbosity	int			`yaml:"verbosity,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#systemlog-component-options
type SystemLogComponents struct {
	AccessControl	*SystemLogComponent		`yaml:"accessControl,omitempty"`
	Command		*SystemLogComponent		`yaml:"command,omitempty"`
	Control		*SystemLogComponent		`yaml:"control,omitempty"`
	Ftdc		*SystemLogComponent		`yaml:"ftdc,omitempty"`
	Geo		*SystemLogComponent		`yaml:"geo,omitempty"`
	Index		*SystemLogComponent		`yaml:"index,omitempty"`
	Network		*SystemLogComponent		`yaml:"network,omitempty"`
	Query		*SystemLogComponent		`yaml:"query,omitempty"`
	Replication	*SystemLogComponent		`yaml:"replication,omitempty"`
	Sharding	*SystemLogComponent		`yaml:"sharding,omitempty"`
	Storage		*SystemLogComponentStorage	`yaml:"storage,omitempty"`
	Write		*SystemLogComponent		`yaml:"write,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#systemlog-options
type SystemLog struct {
	Destination		string	`yaml:"destination,omitempty"`
	LogAppend		bool	`yaml:"logAppend,omitempty"`
	LogRotate		string	`yaml:"logRotate,omitempty"`
	Path			string	`yaml:"path,omitempty"`
	Quiet			bool	`yaml:"quiet,omitempty"`
	TimeStampFormat		string	`yaml:"timeStampFormat,omitempty"`
	TraceAllExceptions	bool	`yaml:"traceAllExceptions,omitempty"`
	SyslogFacility		string	`yaml:"syslogFacility,omitempty"`
	Verbosity		int	`yaml:"verbosity,omitempty"`
}
