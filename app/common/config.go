package common

import (
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

const DefaultConfigFileType = "toml"

var cfg Config

type ConfigOptions struct {
	viper *viper.Viper
	opts  Options
}

type Options struct {
	configFileName string
	configFileType string
	configFileDir  string
}

type Config struct {
	Core     Core
	Database Database
}

type Core struct {
	Name string
	Port int
	Host string
}

type Database struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

func InitConfig(env string) error {
	configDirPath := getConfigDirPath()
	configOptions := getConfigOptions(configDirPath, env)
	err := configOptions.loadConfigs(&cfg)
	if err != nil {
		return err
	}
	return nil
}

func GetConfig() Config {
	return cfg
}

func getConfigDirPath() string {
	workDir := os.Getenv("WORKDIR")
	configDirPath := ""
	if workDir == "" {
		_, thisFile, _, _ := runtime.Caller(1)
		configDirPath = path.Join(path.Dir(thisFile), "../../configs")
	} else {
		configDirPath = path.Join(workDir, "./configs")
	}
	return configDirPath
}

func getConfigOptions(configDirPath string, env string) *ConfigOptions {
	return &ConfigOptions{
		viper: viper.New(),
		opts: Options{
			configFileName: env,
			configFileType: DefaultConfigFileType,
			configFileDir:  configDirPath,
		},
	}
}

func (c *ConfigOptions) loadConfigs(config interface{}) error {
	c.viper.AddConfigPath(c.opts.configFileDir)
	c.viper.SetConfigType(c.opts.configFileType)
	c.viper.SetConfigName(c.opts.configFileName)

	err := c.viper.ReadInConfig()
	if err != nil {
		return err
	}

	return c.viper.Unmarshal(config)
}
