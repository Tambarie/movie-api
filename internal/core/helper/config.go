package helper

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
)

type ConfigStruct struct {
	AppName             string `mapstructure:"app_name"`
	PostgresDatabaseURL string `mapstructure:"postgres_database_url"`
	ServiceAddress      string `mapstructure:"service_address"`
	ServicePort         string `mapstructure:"service_port"`
	ServiceMode         string `mapstructure:"service_mode"`
	DBType              string `mapstructure:"db_type"`
	DBHost              string `mapstructure:"postgres_db_host"`
	PostgresUsername    string `mapstructure:"postgres_db_name"`
	PostgresUser        string `mapstructure:"postgres_db_user"`
	PostgresDBPassword  string `mapstructure:"postgres_db_password"`
	PostgresDBPort      string `mapstructure:"postgres_db_port"`
	PostgresTimezone    string `mapstructure:"postgres_db_timezone"`
	PostgresDBMode      string `mapstructure:"postgres_db_mode"`
	ServiceName         string `mapstructure:"service_name"`
	RedisHost           string `mapstructure:"redis_host"`
	RedisPort           string `mapstructure:"redis_port"`
	LogDir              string `mapstructure:"log_dir"`
	LogFile             string `mapstructure:"log_file"`
	ExternalConfigPath  string `mapstructure:"external_config_path"`
}

var (
	postgres_database_url string
	dbtype                string
	postgresdb_pass       string
	service_address       string
	service_port          string
	postgresdb_mode       string
	postgresdb_host       string
	service_mode          string
	postgresdb_name       string
	postgresdb_user       string
	postgresdb_port       string
	postgresdb_timezone   string

	redis_host         string
	redis_port         string
	externalConfigPath string
)

func LoadConfig() (string, string, string, string, string, string, string, string, string, string, string, string, string, string, string) {
	flag.StringVar(&dbtype, "dbtype", Config.DBType, "application db type")
	flag.StringVar(&postgres_database_url, "postgres_database_url", Config.DBType, "application db type")
	flag.StringVar(&postgresdb_pass, "postgresdb_pass", Config.PostgresDBPassword, "application password")
	flag.StringVar(&service_address, "service_address", Config.ServiceAddress, "local host")
	flag.StringVar(&service_port, "service_port", Config.ServicePort, "application ports")
	flag.StringVar(&service_mode, "service_mode", Config.ServiceMode, "application mode, either dev or production")
	flag.StringVar(&postgresdb_host, "postgres_DBHost", Config.DBHost, "database host")
	flag.StringVar(&postgresdb_mode, "postgresMode", Config.PostgresDBMode, "database mode")
	flag.StringVar(&postgresdb_name, "postgresDBName", Config.PostgresUsername, "database name")
	flag.StringVar(&postgresdb_user, "postgresDBUser", Config.PostgresUser, "database user")
	flag.StringVar(&postgresdb_port, "postgresDBPort", Config.PostgresDBPort, "database port")
	flag.StringVar(&postgresdb_timezone, "postgresDBTimezone", Config.PostgresTimezone, "database timezone")
	flag.StringVar(&redis_host, "redis_host", Config.RedisHost, "redis host")
	flag.StringVar(&redis_port, "redis_port", Config.RedisPort, "redis port")
	flag.StringVar(&externalConfigPath, "external_config_path", Config.ExternalConfigPath, "external config path")
	flag.Parse()
	for i, value := range flag.Args() {
		os.Args[i] = value
	}
	return dbtype, postgres_database_url, postgresdb_pass, service_address, service_port, service_mode, postgresdb_host, postgresdb_mode, postgresdb_name, postgresdb_user, postgresdb_port, postgresdb_timezone, redis_host, redis_port, externalConfigPath
}

func LoadEnv(path string) (config ConfigStruct, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("movie-service")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return ConfigStruct{}, err
	}
	err = viper.Unmarshal(&config)
	return
}

func ReturnConfig() ConfigStruct {
	config, err := LoadEnv(".")
	if err != nil {
		log.Println(err)
	}
	if config.ExternalConfigPath != "" {
		viper.Reset()
		config, err = LoadEnv(config.ExternalConfigPath)
		if err != nil {
			log.Println(err)
		}
	}
	return config
}

var Config = ReturnConfig()
