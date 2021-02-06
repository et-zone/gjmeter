package jmeter

import (
	"sync"
	"time"

	"github.com/et-zone/httpclient"
)

type JmeterParam struct {
	ConnSize    int              `json:"member"`      //任务数
	Time_second int              `json:"time_second"` //持续时间
	Count       int              `json:"count"`       //执行总次数
	Method      string           `json:"method"`      //请求类型
	Url         string           `json:"url"`         //地址
	Param       httpclient.Param `json:"param"`       //请求参数
	Body        []byte           `json:"body"`        //请求体
}

type ConfigParam struct {
	ConnSize    int               `json:"member"`      //任务数
	Time_second int               `json:"time_second"` //持续时间
	Count       int               `json:"count"`       //执行总次数
	Method      string            `json:"method"`      //请求类型
	Url         string            `json:"url"`         //地址
	Header      map[string]string `json:"header"`      //请求参数
	Body        string            `json:"body"`        //请求体
}

func DoTask(jmeterParam JmeterParam) {

	if jmeterParam.Time_second == 0 {
		mutex := &sync.RWMutex{}
		c := jmeterParam.Count
		clients := initClients(jmeterParam.ConnSize, jmeterParam.Param)
		chs := make([]chan int, jmeterParam.ConnSize)
		for i, _ := range clients {
			chs[i] = make(chan int)

			go run(chs[i], mutex, clients[i], &c, jmeterParam.Method, jmeterParam.Url, jmeterParam.Body)
		}
		for _, ch := range chs {
			<-ch
		}
		closeClients(&clients)
	} else {

		clients := initClients(jmeterParam.ConnSize, jmeterParam.Param)
		chs := make([]chan int, jmeterParam.ConnSize)
		for i, cli := range clients {
			chs[i] = make(chan int)
			go runtime(chs[i], cli, jmeterParam.Time_second, jmeterParam.Method, jmeterParam.Url, jmeterParam.Body)
		}
		for _, ch := range chs {
			<-ch
		}
		closeClients(&clients)
	}

}

func runtime(ch chan int, cli *httpclient.Client, time_second int, method string, url string, body []byte) {
	t := time.Now().Add(time.Duration(time_second) * time.Second)
	for {
		if t.Before(time.Now()) {
			break
		}
		ctx := httpclient.NewContext()
		cli.Dao(ctx, method, url, body)
		baseinfo := BaseInfo{}
		_, _, _, _, _, baseinfo.duration, baseinfo.Code = ctx.GeteContextInfo()
		UpdateInfo(baseinfo)
	}
	ch <- 1
}

func run(ch chan int, mutex *sync.RWMutex, cli *httpclient.Client, count *int, method string, url string, body []byte) {

	for {
		ctx := httpclient.NewContext()
		mutex.Lock()
		if *count < 1 {
			break
		}
		*count = *count - 1
		cli.Dao(ctx, method, url, body)
		baseinfo := BaseInfo{}
		_, _, _, _, _, baseinfo.duration, baseinfo.Code = ctx.GeteContextInfo()
		UpdateInfo(baseinfo)
		mutex.Unlock()

	}
	mutex.Unlock()
	ch <- 1
}
