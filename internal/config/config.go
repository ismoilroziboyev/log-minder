package config

type Config struct {
	MongoURI string `env:"MONGO_URI"`
	MongoDB  string `env:"MONGO_DB"`
	HttpHost string `env:"HTTP_HOST"`
	HttpPort int    `env:"HTTP_PORT"`
}
