package bancho

import "github.com/spf13/viper"

type Config struct {
	BindAddr string

	DbProvider   string
	DbConnection string

	RedisAddr string
	RedisPass string
	RedisDB   int
}

func GetConfig() (config Config, err error) {
	v := viper.New()
	v.SetConfigName("bancho")

	v.SetDefault("BindAddr", "127.0.0.1:6300")
	v.SetDefault("DbProvider", "sqlite3")
	v.SetDefault("DbConnection", "bancho.db")

	v.SetDefault("RedisAddr", "127.0.0.1:6379")
	v.SetDefault("RedisPass", "")
	v.SetDefault("RedisDB", "0")

	v.AddConfigPath(".")
	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&config)
	return
}
