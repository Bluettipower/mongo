package mongo

// Config mongodb configuration parameters
type Config struct {
	URL          string
	DB           string
	Username     string
	Password     string
	Service      string
	IsReplicaSet bool
}

// NewConfigNonReplicaSet create mongodb configuration for a non-replicaSet
func NewConfigNonReplicaSet(url, db, username, password, service string) *Config {
	return &Config{
		URL:          url,
		DB:           db,
		Username:     username,
		Password:     password,
		Service:      service,
		IsReplicaSet: false,
	}
}

// NewConfigReplicaSet create mongodb configuration for a ReplicaSet
func NewConfigReplicaSet(url, db string) *Config {
	return &Config{
		URL:          url,
		DB:           db,
		IsReplicaSet: true,
	}
}
