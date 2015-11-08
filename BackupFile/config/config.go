package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//配置文件的格式, 不用map[string]interface{}是因为用强类型对于简单的配置比较方便
type Myconfig struct {
	Copyfiles          []string //准备复制的文件
	OutputPath         string   //输出的路径
	Username, Password string   //测试
}

var LocalConfig Myconfig

func init() {
	filename := "config.json"
	read(filename) //读取配置文件
	//write(filename) //写入配置文件
}
func write(filename string) {
	//写入配置文件
	LocalConfig := new(Myconfig)
	conf := LocalConfig
	conf.Username = "wo cao zhinengtishi"
	conf.Password = "123"
	conf.Copyfiles = append(conf.Copyfiles, "//I53470/sharedoc/taochi.mdb")
	conf.Copyfiles = append(conf.Copyfiles, "//I53470/sharedoc/test.txt")
	conf.OutputPath = "backup"

	data, err := json.Marshal(conf)
	fout, err := os.Create("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()
	b, err := fout.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Write config file bytes:", b)
}
func read(filename string) {
	//读取配置文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("找不到配置文件,系统不能运行.")
		//log.Fatal(err)
	}
	err = json.Unmarshal(data, &LocalConfig)
	if err != nil {
		log.Fatal("无法解析配置文件的格式.")
		//log.Fatal(err)
	}
}
