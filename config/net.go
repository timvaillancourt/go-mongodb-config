package config

// https://docs.mongodb.com/manual/reference/configuration-options/#net-http-options
type NetHttp struct {
	Enabled              bool `yaml:"enabled,omitempty" json:"enabled,omitempty" json:"enabled,omitempty"`
	JSONPEnabled         bool `yaml:"JSONPEnabled,omitempty" json:"JSONPEnabled,omitempty"`
	RESTInterfaceEnabled bool `yaml:"RESTInterfaceEnabled,omitempty" json:"RESTInterfaceEnabled,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#net-unixdomainsocket-options
type NetUnixDomainSocket struct {
	Enabled         bool   `yaml:"enabled,omitempty" json:"enabled,omitempty" json:"enabled,omitempty"`
	PathPrefix      string `yaml:"pathPrefix,omitempty" json:"pathPrefix,omitempty"`
	FilePermissions int    `yaml:"filePermissions,omitempty" json:"filePermissions,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#net-ssl-options
type NetSSL struct {
	Mode                                string `yaml:"mode,omitempty" json:"mode,omitempty"`
	SSLOnNormalPorts                    bool   `yaml:"sslOnNormalPorts,omitempty" json:"sslOnNormalPorts,omitempty"`
	PEMKeyFile                          string `yaml:"pemKeyFile,omitempty" json:"pemKeyFile,omitempty"`
	PEMKeyPassword                      string `yaml:"pemKeyPassword,omitempty" json:"pemKeyPassword,omitempty"`
	ClusterFile                         string `yaml:"clusterFile,omitempty" json:"clusterFile,omitempty"`
	ClusterPassword                     string `yaml:"clusterPassword,omitempty" json:"clusterPassword,omitempty"`
	CAFile                              string `yaml:"CAFile,omitempty" json:"CAFile,omitempty"`
	CRLFile                             string `yaml:"CRLFile,omitempty" json:"CRLFile,omitempty"`
	AllowConnectionsWithoutCertificates bool   `yaml:"allowConnectionsWithoutCertificates,omitempty" json:"allowConnectionsWithoutCertificates,omitempty"`
	AllowInvalidCertificates            bool   `yaml:"allowInvalidCertificates,omitempty" json:"allowInvalidCertificates,omitempty"`
	AllowInvalidHostnames               bool   `yaml:"allowInvalidHostnames,omitempty" json:"allowInvalidHostnames,omitempty"`
	DisabledProtocols                   string `yaml:"disabledProtocols,omitempty" json:"disabledProtocols,omitempty"`
	FIPSMode                            bool   `yaml:"FIPSMode,omitempty" json:"FIPSMode,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#net-options
type Net struct {
	BindIp                 string               `yaml:"bindIp,omitempty" json:"bindIp,omitempty"`
	Http                   *NetHttp             `yaml:"http,omitempty" json:"http,omitempty"`
	Ipv6                   bool                 `yaml:"ipv6,omitempty" json:"ipv6,omitempty"`
	MaxIncomingConnections int                  `yaml:"maxIncomingConnections,omitempty" json:"maxIncomingConnections,omitempty"`
	Port                   int                  `yaml:"port,omitempty" json:"port,omitempty"`
	UnixDomainSocket       *NetUnixDomainSocket `yaml:"unixDomainSocket,omitempty" json:"unixDomainSocket,omitempty"`
	SSL                    *NetSSL              `yaml:"ssl,omitempty" json:"ssl,omitempty"`
	WireObjectCheck        bool                 `yaml:"wireObjectCheck,omitempty" json:"wireObjectCheck,omitempty"`
}
