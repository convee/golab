package conf

import "github.com/jinzhu/configor"

var AppConfig = struct {
	DB struct {
		Host     string `default:"127.0.0.1"`
		Port     string `default:"3306"`
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
	}
	SiteName string `default:"SiteName"`
}{}

func InitConfig() {
	configor.Load(&AppConfig, "conf/config.toml")
}



