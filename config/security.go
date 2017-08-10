package config

// https://docs.mongodb.com/manual/reference/configuration-options/#security-sasl-options
//type SecuritySasl struct {
//}

// https://docs.mongodb.com/manual/reference/configuration-options/#key-management-configuration-options
//type SecurityKmip struct {
//}

//type SecurityLdapBind struct {
//}

//type SecurityLdapAuthz struct {
//}

// https://docs.mongodb.com/manual/reference/configuration-options/#security-ldap-options
//type SecurityLdap struct {
//	Servers	string			`yaml:"servers,omitempty", json:"servers,omitempty"`
//	Authz	*SecurityLdapAuthz	`yaml:"authz,omitempty", json:"authz,omitempty"`
//	Bind	*SecurityLdapBind	`yaml:"bind,omitempty", json:"bind,omitempty"`
//}

// https://docs.mongodb.com/manual/reference/configuration-options/#security-options
type Security struct {
	Authorization string `yaml:"authorization,omitempty", json:"authorization,omitempty"`
	KeyFile       string `yaml:"keyFile,omitempty", json:"keyFile,omitempty"`
	//	Kmip		*SecurityKmip	`yaml:"kmip,omitempty", json:"kmip,omitempty"`
	//	Ldap		*SecurityLdap	`yaml:"ldap,omitempty", json:"ldap,omitempty"`
	//	Sasl		*SecuritySasl	`yaml:"sasl,omitempty", json:"sasl,omitempty"`
}
