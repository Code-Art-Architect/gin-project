package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

// Logrus 初始化一个记录器实例
var Logrus = logrus.New()

type LogConfig struct {
	LogDir   string `json:"log_dir"`
	LogLevel string `json:"log_level"`
}

func init() {
	config := LoadLogConfig()
	file, err := os.OpenFile(config.LogDir, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	Logrus.Out = file

	// 定义map，用于存储日志级别
	logLevelMap := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"panic": logrus.PanicLevel,
		"fatal": logrus.FatalLevel,
		"error": logrus.ErrorLevel,
		"warn": logrus.WarnLevel,
		"info": logrus.InfoLevel,
		"debug": logrus.DebugLevel,
	}

	Logrus.SetLevel(logLevelMap[config.LogLevel])
	Logrus.SetFormatter(&logrus.TextFormatter{})
}

func LoadLogConfig() *LogConfig {
	conf := LogConfig{}
	file, err := os.Open("/Users/element/GoProjects/gin-project/conf/log-config.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &conf); err != nil {
		panic(err)
	}

	return &conf
}
