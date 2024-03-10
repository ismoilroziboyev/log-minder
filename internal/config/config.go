package config

type Config struct {
	MongoURI      string `env:"MONGO_URI"`
	MongoDB       string `env:"MONGO_DB"`
	HttpHost      string `env:"HTTP_HOST,default=0.0.0.0"`
	HttpPort      int    `env:"HTTP_PORT,default=35468"`
	AdminUser     string `env:"ADMIN_USER"`
	AdminPassword string `env:"ADMIN_PASSWORD"`
}
