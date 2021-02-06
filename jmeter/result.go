package jmeter

import (
	"fmt"
	"time"
)

var rep = Report{MinTime: 9999999999}

type BaseInfo struct {
	appName  string
	method   string
	ip       string
	path     string        //不需要ip地址
	nowtime  *time.Time    //
	duration time.Duration //请求时长
	Code     int           //状态吗

}

type Report struct {
	Total        int //总数
	SuccCount    int //成功的数量
	FailureCount int //失败或者超时数量
	MaxTime      int //最长响应时间
	MinTime      int //最短响应时间
	TotalTime    int //总时长
}

func UpdateInfo(info BaseInfo) {
	rep.Total += 1
	if info.Code == 200 {
		rep.SuccCount += 1
	} else {
		rep.FailureCount += 1
	}
	if int(info.duration) > rep.MaxTime {
		rep.MaxTime = int(info.duration)
	}
	if int(info.duration) < rep.MinTime {
		rep.MinTime = int(info.duration)
	}
	rep.TotalTime += int(info.duration)
}

func PrintReport(time_second int) {
	if time_second <= 0 {
		fmt.Println("运行时长不能为0")
		return
	}
	fmt.Println("总 请 求 数:", rep.Total,
		"\n成功请求数 :", rep.SuccCount,
		"\n失败请求数 :", rep.FailureCount,
		"\nMax响应时长:", time.Duration(rep.MaxTime),
		"\nMin响应时长:", time.Duration(rep.MinTime),
		"\n总响应时长 :", time.Duration(rep.TotalTime),
		"\n平均响应时长:", time.Duration(rep.TotalTime/rep.Total),
		"\nQPS :     ", rep.Total/time_second) //QPS（TPS） = req/sec = 请求数/秒

}
