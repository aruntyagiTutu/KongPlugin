package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

type Config struct {
	url string
}

func New() interface{} {
	return &Config{}
}

func (config Config) Access(kong *pdk.PDK) {
	//x := make(map[string][]string)
	//x["x-user-id"] = append(x["x-user-id"], config.url)
}

func main() {
	Version := "1.1"
	Priority := 999
	_ = server.StartServer(New, Version, Priority)
}
