package config

// https://docs.mongodb.com/manual/reference/configuration-options/#net-http-options
type NetHttp struct {
	Enabled              bool `yaml:"enabled,omitempty" json:"enabled,omitempty" bson:"enabled,omitempty"`
	JSONPEnabled         bool `yaml:"JSONPEnabled,omitempty" json:"JSONPEnabled,omitempty" bson:"JSONPEnabled,omitempty"`
	RESTInterfaceEnabled bool `yaml:"RESTInterfaceEnabled,omitempty" json:"RESTInterfaceEnabled,omitempty" bson:"RESTInterfaceEnabled,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#net-unixdomainsocket-options
type NetUnixDomainSocket struct {
	Enabled         bool   `yaml:"enabled,omitempty" json:"enabled,omitempty" bson:"enabled,omitempty"`
	PathPrefix      string `yaml:"pathPrefix,omitempty" json:"pathPrefix,omitempty" bson:"pathPrefix,omitempty"`
	FilePermissions int    `yaml:"filePermissions,omitempty" json:"filePermissions,omitempty" bson:"filePermissions,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#net-ssl-options
type NetSSL struct {
	Mode                                string `yaml:"mode,omitempty" json:"mode,omitempty" bson:"mode,omitempty"`
	SSLOnNormalPorts                    bool   `yaml:"sslOnNormalPorts,omitempty" json:"sslOnNormalPorts,omitempty" bson:"sslOnNormalPorts,omitempty"`
	PEMKeyFile                          string `yaml:"pemKeyFile,omitempty" json:"pemKeyFile,omitempty" bson:"pemKeyFile,omitempty"`
	PEMKeyPassword                      string `yaml:"pemKeyPassword,omitempty" json:"pemKeyPassword,omitempty" bson:"pemKeyPassword,omitempty"`
	ClusterFile                         string `yaml:"clusterFile,omitempty" json:"clusterFile,omitempty" bson:"clusterFile,omitempty"`
	ClusterPassword                     string `yaml:"clusterPassword,omitempty" json:"clusterPassword,omitempty" bson:"clusterPassword,omitempty"`
	CAFile                              string `yaml:"CAFile,omitempty" json:"CAFile,omitempty" bson:"CAFile,omitempty"`
	CRLFile                             string `yaml:"CRLFile,omitempty" json:"CRLFile,omitempty" bson:"CRLFile,omitempty"`
	AllowConnectionsWithoutCertificates bool   `yaml:"allowConnectionsWithoutCertificates,omitempty" json:"allowConnectionsWithoutCertificates,omitempty" bson:"allowConnectionsWithoutCertificates,omitempty"`
	AllowInvalidCertificates            bool   `yaml:"allowInvalidCertificates,omitempty" json:"allowInvalidCertificates,omitempty" bson:"allowInvalidCertificates,omitempty"`
	AllowInvalidHostnames               bool   `yaml:"allowInvalidHostnames,omitempty" json:"allowInvalidHostnames,omitempty" bson:"allowInvalidHostnames,omitempty"`
	DisabledProtocols                   string `yaml:"disabledProtocols,omitempty" json:"disabledProtocols,omitempty" bson:"disabledProtocols,omitempty"`
	FIPSMode                            bool   `yaml:"FIPSMode,omitempty" json:"FIPSMode,omitempty" bson:"FIPSMode,omitempty"`
}

// https://docs.mongodb.com/manual/reference/configuration-options/#net-options
type Net struct {
	BindIp                 string               `yaml:"bindIp,omitempty" json:"bindIp,omitempty" bson:"bindIp,omitempty"`
	Http                   *NetHttp             `yaml:"http,omitempty" json:"http,omitempty" bson:"http,omitempty"`
	Ipv6                   bool                 `yaml:"ipv6,omitempty" json:"ipv6,omitempty" bson:"ipv6,omitempty"`
	MaxIncomingConnections int                  `yaml:"maxIncomingConnections,omitempty" json:"maxIncomingConnections,omitempty" bson:"maxIncomingConnections,omitempty"`
	Port                   int                  `yaml:"port,omitempty" json:"port,omitempty" bson:"port,omitempty"`
	UnixDomainSocket       *NetUnixDomainSocket `yaml:"unixDomainSocket,omitempty" json:"unixDomainSocket,omitempty" bson:"unixDomainSocket,omitempty"`
	SSL                    *NetSSL              `yaml:"ssl,omitempty" json:"ssl,omitempty" bson:"ssl,omitempty"`
	WireObjectCheck        bool                 `yaml:"wireObjectCheck,omitempty" json:"wireObjectCheck,omitempty" bson:"wireObjectCheck,omitempty"`
}
