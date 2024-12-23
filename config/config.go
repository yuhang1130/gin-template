package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Server struct {
	Port  string
	Level string
}

type DataSource struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string `mapstructure:"db_name"`
	Config   string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DataBase int    `mapstructure:"data_base"`
	StoreDb  string `mapstructure:"store_db"`
}

type Jwt struct {
	Secret string
	TTL    int64
	Name   string
}

type Log struct {
	Level string
}

type AllConfig struct {
	Server     Server
	DataSource DataSource
	Redis      Redis
	Jwt        Jwt
	Log        Log
}

func InitLoadConfig() *AllConfig {
	// 读取环境变量
	env := os.Getenv("env")
	if env == "" {
		// 如果环境变量未设置，则使用命令行标志
		envStr := flag.String("env", "dev", "Environment: dev or prod")
		flag.Parse()
		env = *envStr
	}
	println("Environment:", env)
	// 设置配置文件的名称（不需要扩展名）
	viper.SetConfigName(fmt.Sprintf("gin-%s", env))

	// 设置配置文件的类型为yaml
	viper.SetConfigType("yaml")

	// 设置配置文件的路径（这里是相对于当前工作目录）
	viper.AddConfigPath("./config")
	// viper.AddConfigPath("/Users/yuhanghe/2024Codes/yuhang/gin-template/config") //go build不会自动包含或调整资源文件的位置

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 自动反序列化到结构
	// var configData *AllConfig
	configData := &AllConfig{}
	if err := viper.Unmarshal(configData); err != nil {
		fmt.Printf("Read config file to struct err: %s\n", err.Error())
		panic("Read config file to struct err")
	}
	fmt.Printf("Unmarshal file success. configData: %+v\n", configData)

	return configData
}

func (d *DataSource) Dsn() string {
	return d.UserName + ":" + d.Password + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.DBName + "?" + d.Config
}
