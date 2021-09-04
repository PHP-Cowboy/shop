package data_source

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type MysqlConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
	UserName string `json:"userName"`
	Pwd string `json:"pwd"`
	BaseName string `json:"baseName"`
	LogMode bool `json:"logMode"`
}

func LoadMysqlConf() *MysqlConf {
	conf := MysqlConf{}
	file,err := os.Open("conf/mysql.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data,readErr := ioutil.ReadAll(file)
	if readErr != nil {
		panic(readErr)
	}
	parseErr := json.Unmarshal(data,&conf)
	if parseErr != nil {
		fmt.Println(err)
		panic(err)
	}
	return &conf
}