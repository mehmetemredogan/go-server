package main

import (
	"github.com/mehmetemredogan/go-server/internal/boot"
)

var (
	serviceConf = "configs/service.json"
)

func main() {
	config()

	boot.EnvLoader()

	boot.Starter()
}

func config() {
	boot.Reader(serviceConf)
}