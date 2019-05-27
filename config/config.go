package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// Config 对应配置文件结构
type Config struct {
	Listen    string              `toml:"listen"`
	DBServers map[string]DBServer `toml:"dbservers"`
}

// UnmarshalConfig 解析toml配置
func UnmarshalConfig(tomlfile string) (*Config, error) {
	c := &Config{}
	if _, err := toml.DecodeFile(tomlfile, c); err != nil {
		return c, err
	}
	return c, nil
}

// DBServerConf 获取数据库配置
func (c Config) DBServerConf(key string) (DBServer, bool) {
	s, ok := c.DBServers[key]
	return s, ok
}

// GetListenAddr 监听地址
func (c Config) GetListenAddr() string {
	return c.Listen
}

// DBServer 表示DB服务器配置
type DBServer struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DBName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// ConnectString 表示连接数据库的字符串
func (m DBServer) ConnectString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.Host, m.Port, m.User, m.Password, m.DBName)
}

//LoadJSON 读取json
func LoadJSON(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
