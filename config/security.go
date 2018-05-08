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
//	Servers	string			`yaml:"servers,omitempty" json:"servers,omitempty" bson:"servers,omitempty"`
//	Authz	*SecurityLdapAuthz	`yaml:"authz,omitempty" json:"authz,omitempty" bson:"authz,omitempty"`
//	Bind	*SecurityLdapBind	`yaml:"bind,omitempty" json:"bind,omitempty" bson:"bind,omitempty"`
//}

// https://docs.mongodb.com/manual/reference/configuration-options/#security-options
type Security struct {
	Authorization        string `yaml:"authorization,omitempty" json:"authorization,omitempty" bson:"authorization,omitempty"`
	ClusterAuthMode      string `yaml:"clusterAuthMode,omitempty" json:"clusterAuthMode,omitempty" bson:"clusterAuthMode,omitempty"`
	EnableEncryption     bool   `yaml:"enableEncryption,omitempty" json:"enableEncryption,omitempty" bson:"enableEncryption,omitempty"`
	EncryptionCipherMode string `yaml:"encryptionCipherMode,omitempty" json:"encryptionCipherMode,omitempty" bson:"encryptionCipherMode,omitempty"`
	EncryptionKeyFile    string `yaml:"encryptionKeyFile,omitempty" json:"encryptionKeyFile,omitempty" bson:"encryptionKeyFile,omitempty"`
	JavascriptEnabled    bool   `yaml:"javascriptEnabled,omitempty" json:"javascriptEnabled,omitempty" bson:"javascriptEnabled,omitempty"`
	KeyFile              string `yaml:"keyFile,omitempty" json:"keyFile,omitempty" bson:"keyFile,omitempty"`
	RedactClientLogData  bool   `yaml:"redactClientLogData,omitempty" json:"redactClientLogData,omitempty" bson:"redactClientLogData,omitempty"`
	TransitionToAuth     bool   `yaml:"transitionToAuth,omitempty" json:"transitionToAuth,omitempty" bson:"transitionToAuth,omitempty"`
	//	Kmip		*SecurityKmip	`yaml:"kmip,omitempty" json:"kmip,omitempty" bson:"kmip,omitempty"`
	//	Ldap		*SecurityLdap	`yaml:"ldap,omitempty" json:"ldap,omitempty" bson:"ldap,omitempty"`
	//	Sasl		*SecuritySasl	`yaml:"sasl,omitempty" json:"sasl,omitempty" bson:"sasl,omitempty"`
}
