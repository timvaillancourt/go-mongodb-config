package config

// https://docs.mongodb.com/manual/reference/configuration-options/#net-http-options
type NetHttp struct {
	Enabled			bool	`yaml:"enabled,omitempty"`
	JSONPEnabled		bool	`yaml:"JSONPEnabled,omitempty"`
	RESTInterfaceEnabled	bool	`yaml:"RESTInterfaceEnabled,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#net-unixdomainsocket-options
type NetUnixDomainSocket struct {
	Enabled		bool	`yaml:"enabled,omitempty"`
	PathPrefix	string	`yaml:"pathPrefix,omitempty"`
	FilePermissions	int	`yaml:"filePermissions,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#net-ssl-options
type NetSSL struct {
	Mode					string	`yaml:"mode,omitempty"`
	SSLOnNormalPorts			bool	`yaml:"sslOnNormalPorts,omitempty"`
	PEMKeyFile				string	`yaml:"pemKeyFile,omitempty"`
	PEMKeyPassword				string	`yaml:"pemKeyPassword,omitempty"`
	ClusterFile				string	`yaml:"clusterFile,omitempty"`
	ClusterPassword				string	`yaml:"clusterPassword,omitempty"`
	CAFile					string	`yaml:"CAFile,omitempty"`
	CRLFile					string	`yaml:"CRLFile,omitempty"`
	AllowConnectionsWithoutCertificates	bool	`yaml:"allowConnectionsWithoutCertificates,omitempty"`
	AllowInvalidCertificates		bool	`yaml:"allowInvalidCertificates,omitempty"`
	AllowInvalidHostnames			bool	`yaml:"allowInvalidHostnames,omitempty"`
	DisabledProtocols			string	`yaml:"disabledProtocols,omitempty"`
	FIPSMode				bool	`yaml:"FIPSMode,omitempty"`

}

// https://docs.mongodb.com/manual/reference/configuration-options/#net-options
type Net struct {
	BindIp			string			`yaml:"bindIp,omitempty"`
	Http			*NetHttp		`yaml:"http,omitempty"`
	Ipv6			bool			`yaml:"ipv6,omitempty"`
	MaxIncomingConnections	int			`yaml:"maxIncomingConnections,omitempty"`
	Port			int			`yaml:"port,omitempty"`
	UnixDomainSocket	*NetUnixDomainSocket	`yaml:"unixDomainSocket,omitempty"`
	SSL			*NetSSL			`yaml:"ssl,omitempty"`
	WireObjectCheck		bool			`yaml:"wireObjectCheck,omitempty"`
}
