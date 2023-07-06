package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	cfg *Config
)

// C 全局调用配置对象 单例模式
func C() *Config {
	if cfg == nil {
		panic("Load Config first")
	}
	return cfg
}

// LoadConfig 加载配置
func LoadConfig(path string) error {
	// 初始化config实例
	cfgObj := NewConfig()
	// 读取配置文件
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("配置文件读取失败:%s\n", err)
		return err
	}
	// 映射指结构体
	if err := yaml.Unmarshal(file, cfgObj); err != nil {
		fmt.Printf("配置映射失败:%s\n", err)
		return err
	}
	// 赋值
	cfg = cfgObj
	return nil
}
