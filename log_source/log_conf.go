package log_source

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type LogConfig struct {
	LogDir string `json:"logDir"`
	LogLevel string `json:"logLevel"`
}

func LoadLogConfig() *LogConfig {
	logConfig := &LogConfig{}
	file,err := os.Open("conf/log_conf.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data,readErr := ioutil.ReadAll(file)
	if readErr != nil {
		panic(readErr)
	}
	fmt.Println(data)
	parseErr := json.Unmarshal(data,logConfig)
	if parseErr != nil {
		panic(parseErr)
	}
	fmt.Println(logConfig)
	return logConfig
}