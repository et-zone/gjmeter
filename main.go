package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/et-zone/gjmeter/jmeter"
	"github.com/et-zone/httpclient"
)

var path = "config.json"

func main() {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("open config err", err.Error())
		return
	}
	config := jmeter.ConfigParam{}
	err = json.Unmarshal(b, &config)
	if err != nil {
		fmt.Println("config err", err.Error())
		return
	}
	fmt.Println(config)
	p := httpclient.NewParam()
	for k, v := range config.Header {
		p.SetParam(k, v)
	}

	param := jmeter.JmeterParam{
		ConnSize:    config.ConnSize,
		Time_second: config.Time_second,
		Count:       config.Count,
		Method:      config.Method,
		Url:         config.Url,
		Param:       p,
	}
	jmeter.DoTask(param)

	jmeter.PrintReport(param.Time_second)
}

/*
	param := jmeter.JmeterParam{
		ConnSize:    10,
		Time_second: 2,
		Count:       100,
		Method:      "POST",
		Url:         "http://127.0.0.1:8888/ping",
		Param:       p,
	}
*/

/*
{
    "member":10,
    "time_second":5,
    "count":0,
    "method":"POST",
    "url":"http://127.0.0.1:8888/ping",
    "header":{},
    "body":""
}
*/
