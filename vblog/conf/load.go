package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

//	github.com/BurntSushi/toml go里使用比较广泛的toml格式解析库
//
// 查看该库基本用法 https://github.com/BurnSushi/toml
// Object 转化为 Toml 配置文件
func LoadFromFile(filepath string) error {
	c := DefaultConfig()
	if _, err := toml.DecodeFile(filepath, c); err != nil {
		return err
	}
	config = c
	return nil
}

// 读取环境变量 "github.com/caarlos0/env/v6"
// env -----> Object
func LoadFromEnv() error {
	// env.Parse
	c := DefaultConfig()
	// env Tag
	if err := env.Parse(c); err != nil {
		return err
	}
	config = c
	//c.MySQL.Host = os.Getenv("DATASOURCE_HOST")
	return nil
}
