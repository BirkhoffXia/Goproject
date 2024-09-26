package conf

import (
	"encoding/json"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 这里不采用直接暴露变量的方式，比较好的方式，使用函数
var config *Config

// 这里可以补充逻辑 初始化结构体
func C() *Config {
	// sync.Lock
	if config == nil {
		// 给一个默认值
		config = DefaultConfig()
	}
	return config
}

func DefaultConfig() *Config {
	return &Config{
		Application: &Application{
			"127.0.0.1",
		},
		MySQL: &MySQL{
			Host:     "10.30.17.173", // Host:     "127.0.0.1",
			Port:     3306,
			DB:       "vblog",
			Username: "root",
			Password: "boyi@BJroot123", // Password: "123456",
			Debug:    true,
		},
	}
}

// db对象也是一个单例模式
type MySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`

	// 判断这个私有属性，来判断是否返回自己有的对象
	db *gorm.DB
	l  sync.Mutex
}

// 程序配置对象，启动时，会读取配置，并且为程序提供需要全局变量
// 把配置对象做出全局变量(单列模式)
type Config struct {
	Application *Application `json:"app" yaml:"app" toml:"app"`
	MySQL       *MySQL       `json:"mysql" yaml:"mysql" toml:"mysql"`
}

type Application struct {
	Domain string `json:"domain" yaml:"domain" toml:"domain" env:"APP_DOMAIN"`
}

// fmt.Stringger 如果想要自定义 对象fmt.PrintXXX() 打印的值
// String() string
// &{0xc00009c780} 转化为Json 好读取
func (c *Config) String() string {
	jd, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return fmt.Sprintf("%p", c)
	}
	return string(jd)
}

// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m *MySQL) DSN() string {
	// fmt.Println("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DB)
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}

// 通过配置就能通过一个DB实例
func (m *MySQL) GetDB() *gorm.
	DB {

	m.l.Lock()
	defer m.l.Unlock()

	if m.db == nil {
		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.db = db
		// 补充Debug配置
		if m.Debug {
			m.db = db.Debug()
		}
	}
	return m.db
}

func (c *Config) DB() *gorm.DB {
	return c.MySQL.GetDB()
}
