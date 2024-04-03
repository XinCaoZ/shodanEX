package conf

import (
	"time"
)

// Config 结构体用于存储当前时间
type Config struct {
	Time time.Time
}

// NewConfig 函数用于创建一个包含当前时间的 Config 实例
func NewConfig() *Config {
	return &Config{
		Time: time.Now(), // 设置当前时间
	}
}

// Init 方法用于初始化配置
func Init() *Config {
	return NewConfig()
}

// PrintCurrentTime 函数用于打印 Config 实例中的当前时间
func PrintCurrentTime(config *Config) string {
	return config.Time.Format("this time is: " + time.RFC3339)
}
