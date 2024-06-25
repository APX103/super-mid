package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const DefaultCacheSize = 2048

// make them configurable
var defaultLimitConfigs = []*LimitConfig{
	{
		LimitType: "server",
		QPS:       2,
		CacheSize: 0,
		Burst:     4,
	},
	{
		LimitType: "ip",
		QPS:       2,
		CacheSize: 0,
		Burst:     4,
	},
}

// BaseConfig base.yml
type BaseConfig struct {
	AppID               string            `yaml:"app_id"`
	AppSecret           string            `yaml:"app_secret"`
	EncryptKey          string            `yaml:"encrypt_key"`
	RepoGroupNameMap    map[string]string `yaml:"group_config,omitempty"`
	RepoGroupWebhookMap map[string]string `yaml:"ex_group_config,omitempty"`
	MongoConfig         MongoConfig       `yaml:"mongo_config"`
	RedisConfig         RedisConfig       `yaml:"redis_config"`
	LimitConfig         []*LimitConfig    `yaml:"limit_config,omitempty"`
	SSOConfig           SSOConfig         `yaml:"sso_config"`
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

type SSOConfig struct {
	AK  string `yaml:"ak"`
	SK  string `yaml:"sk"`
	JSK string `yaml:"jwt_secret_token"`
}

type LimitConfig struct {
	LimitType string `yaml:"limitType"`
	Burst     int    `yaml:"burst"`
	QPS       int    `yaml:"qps"`
	CacheSize int    `yaml:"cacheSize"`
}

func (c *LimitConfig) Validate() error {
	logrus.Debug("validate limit config")
	if c.QPS == 0 || c.Burst == 0 {
		return fmt.Errorf("LimitConfig Burst and QPS cannot be empty")
	}
	if c.QPS > c.Burst {
		return fmt.Errorf("LimitConfig QPS(%d) must less than Burst(%d)", c.QPS, c.Burst)
	}
	if c.CacheSize == 0 {
		c.CacheSize = DefaultCacheSize
	}
	return nil
}

func NewBaseConfig() *BaseConfig {
	logrus.Debug(" [Fx] BaseConfig Init ")
	conf := BaseConfig{}
	conf = conf.GetConfig()
	if len(conf.LimitConfig) == 0 {
		logrus.Debug("cannot get limit config from config, using default.")
		conf.LimitConfig = defaultLimitConfigs
	}
	return &conf
}

func (i *BaseConfig) GetConfig() BaseConfig {
	logrus.Debug(" Getting config ")
	conf := &BaseConfig{}
	env := os.Getenv("LARKBOT_ENV")
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
		var err error
		conf.AppID = os.Getenv("LarkbotEnvAppID")
		conf.AppSecret = os.Getenv("LarkbotEnvAppSecret")
		conf.EncryptKey = os.Getenv("LarkbotEnvEncryptKey")
		conf.MongoConfig.Url = os.Getenv("LarkbotEnvMongoConfigUrl")
		conf.MongoConfig.Port = os.Getenv("LarkbotEnvMongoConfigPort")
		conf.MongoConfig.Username = os.Getenv("LarkbotEnvMongoConfigUsername")
		conf.MongoConfig.Password = os.Getenv("LarkbotEnvMongoConfigPassword")
		conf.MongoConfig.Database = os.Getenv("LarkbotEnvMongoConfigDatabase")
		conf.RedisConfig.Url = os.Getenv("LarkbotEnvRedisConfigUrl")
		conf.RedisConfig.Port = os.Getenv("LarkbotEnvRedisConfigPort")
		conf.RedisConfig.Password = os.Getenv("LarkbotEnvRedisConfigPassword")
		conf.RedisConfig.Database, err = strconv.Atoi(os.Getenv("LarkbotEnvRedisConfigDatabase"))
		conf.SSOConfig.AK = os.Getenv("LarkbotEnvSSOConfigAK")
		conf.SSOConfig.SK = os.Getenv("LarkbotEnvSSOConfigSK")
		conf.SSOConfig.JSK = os.Getenv("LarkbotEnvSSOConfigJSK")
		if err != nil {
			panic("Can not get config file")
		}
		return *conf
	}
}
