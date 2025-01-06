package configs

type (
	Config struct {
		Service  Service  `mapstructure:"service"`
		Database Database `mapstructure:"database"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretJWT string `mapsctructure:"secretJWT"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}
)
