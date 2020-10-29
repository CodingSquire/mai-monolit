package logger

type Config struct {
	Level     string `env:"LOGGER_LEVEL,default=debug"`
	Timestamp bool   `env:"LOGGER_TIMESTAMP,default=true"`
	Caller    bool   `env:"LOGGER_CALLER,default=true"`
	Pretty    bool   `env:"LOGGER_PRETTY,default=true"`
}
