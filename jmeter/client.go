package jmeter

import (
	"github.com/et-zone/httpclient"
)

func initClients(size int, param httpclient.Param) (clients []*httpclient.Client) {
	clients = []*httpclient.Client{}
	for i := 0; i < size; i++ {
		cli := httpclient.InitDefaultClient()
		cli.Param = param
		clients = append(clients, cli)
	}
	return clients
}

func closeClients(clients *[]*httpclient.Client) {
	for _, cli := range *clients {
		cli.Close()
	}
}
