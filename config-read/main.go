package main

import (
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/config"
)

type keyConfig struct {
	Subkey  string `json:"subkey"`
	Subkey1 int    `json:"subkey1"`
	Subkey2 string `json:"subkey2"`
}

type conf struct {
	Key keyConfig `json:"key"`
}

func main() {
	// get a value
	go func() {
		for {
			time.Sleep(time.Second)

			fmt.Println(config.Get("key", "subkey").String(""))

		}
	}()

	// run the service
	service.Run()
}
