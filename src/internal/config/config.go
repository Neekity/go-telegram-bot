package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Conf Config

type Config struct {
	BotToken        string                    `yaml:"botToken"`
	PhotoConfigs    map[string]PhotoConfig    `yaml:"photoConfigs"`
	SupportChatIds  []int64                   `yaml:"supportChatIds"`
	ResourceConfigs map[string]ResourceConfig `yaml:"resourceConfigs"`
}

type PhotoConfig struct {
	Url        string `yaml:"url"`
	PreviewUrl string `yaml:"previewUrl"`
}

type ResourceConfig struct {
	Url        string `yaml:"url"`
	PreviewUrl string `yaml:"previewUrl"`
}

type SubscribeUser struct {
	ChatId int64 `yaml:"chatId"`
}

func init() {
	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	if err := v.Unmarshal(&Conf); err != nil {
		log.Panic(err)
	}

	fmt.Println("CONFIG => ", Conf)
}
