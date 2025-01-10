package configs

type (
	Config struct {
		Service       Service       `mapstructure:"service"`
		Database      Database      `mapstructure:"database"`
		SpotifyConfig SpotifyConfig `mapstructure:"spotify"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretJWT string `mapsctructure:"secretJWT"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}

	SpotifyConfig struct {
		ClientID     string `mapstructure:"clientID"`
		ClientSecret string `mapstructure:"clientSecret"`
	}
)
