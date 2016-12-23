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
//	Servers	string			`yaml:"servers,omitempty"`
//	Authz	*SecurityLdapAuthz	`yaml:"authz,omitempty"`
//	Bind	*SecurityLdapBind	`yaml:"bind,omitempty"`
//}

// https://docs.mongodb.com/manual/reference/configuration-options/#security-options
type Security struct {
	Authorization	string		`yaml:"authorization,omitempty"`
	KeyFile		string		`yaml:"keyFile,omitempty"`
//	Kmip		*SecurityKmip	`yaml:"kmip,omitempty"`
//	Ldap		*SecurityLdap	`yaml:"ldap,omitempty"`
//	Sasl		*SecuritySasl	`yaml:"sasl,omitempty"`
}
