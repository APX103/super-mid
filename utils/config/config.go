package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// BaseConfig base.yml
type BaseConfig struct {
	AppID       string      `yaml:"app_id"`
	AppSecret   string      `yaml:"app_secret"`
	EncryptKey  string      `yaml:"encrypt_key"`
	MongoConfig MongoConfig `yaml:"mongo_config"`
}

type MongoConfig struct {
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"db"`
}

type RedisConfig struct {
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	Password string `yaml:"password,omitempty"`
	Database int    `yaml:"db,omitempty"`
}

func NewBaseConfig() *BaseConfig {
	logrus.Debug(" [Fx] BaseConfig Init ")
	conf := BaseConfig{}
	conf = conf.GetConfig()
	return &conf
}

func (i *BaseConfig) GetConfig() BaseConfig {
	logrus.Debug(" Getting config ")
	conf := &BaseConfig{}
	env := os.Getenv("SUPERMID_ENV")
	fmt.Println(env)
	if env == "" {
		env = "dev"
	}
	if env != "k8s" {
		yamlFile, err := os.ReadFile("config/" + env + ".yaml")
		if err != nil {
			logrus.Error(err)
			panic("Can not get config file")
		}

		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			panic(err)
		}
		logrus.Debug(" Got config ")
		return *conf
	} else {
		conf.AppID = os.Getenv("SUPERMIDEnvAppID")
		conf.AppSecret = os.Getenv("SUPERMIDEnvAppSecret")
		conf.EncryptKey = os.Getenv("SUPERMIDEnvEncryptKey")
		conf.MongoConfig.Url = os.Getenv("SUPERMIDEnvMongoConfigUrl")
		conf.MongoConfig.Port = os.Getenv("SUPERMIDEnvMongoConfigPort")
		conf.MongoConfig.Username = os.Getenv("SUPERMIDEnvMongoConfigUsername")
		conf.MongoConfig.Password = os.Getenv("SUPERMIDEnvMongoConfigPassword")
		conf.MongoConfig.Database = os.Getenv("SUPERMIDEnvMongoConfigDatabase")
		return *conf
	}
}
