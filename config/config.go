package config

import (
	"fmt"
)

type App struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// GetHost 拼接IP+端口
func (a *App) GetHost() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

type Logger struct {
	Level      string `yaml:"level"`
	FileName   string `yaml:"fileName"`
	MaxSize    int    `yaml:"maxSize"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
	Compress   bool   `yaml:"comPress"`
}

type Config struct {
	App        *App              `yaml:"app"`
	Log        *Logger           `yaml:"log"`
	Ding       map[string]string `yaml:"dingtoken"`
	Jvm        *Jvm              `yaml:"jvm"`
	PromLabels *PromLabels       `yaml:"promlabels"`
}

// NewApp app 默认配置参数
func NewApp() *App {
	return &App{
		Host: "127.0.0.1",
		Port: "8080",
	}
}

// NewConfig Config 默认参数
func NewConfig() *Config {
	return &Config{
		App:  NewApp(),
		Log:  NewLogger(),
		Ding: map[string]string{},
	}
}

// NewLogger 日志默认配置
func NewLogger() *Logger {
	return &Logger{
		Level:      "DEBUG",
		FileName:   "./data/logs/app.log",
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 10,
		Compress:   false,
	}
}

//type DingToken struct {
//	AppName map[string]string
//}

// Jvm 导出配置 结构体对象
type Jvm struct {
	DumpMin   int  `yaml:"dump_min"`
	DumpMax   int  `yaml:"dump_max"`
	DumpTsMin int  `yaml:"dump_ts_min"`
	DumpTsMax int  `yaml:"dump_ts_max"`
	IsDump    bool `yaml:"is_dump"`
}

type PromLabels struct {
	JvmLabel   string `yaml:"jvm_label"`
	EventLabel string `yaml:"event_label"`
}
