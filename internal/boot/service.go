package boot

import (
	"encoding/json"
	"os"
	"strconv"
)

type Service struct {
	General General `json:"general"`
	Network Network `json:"network"`
	SSL     SSL     `json:"ssl"`
}

type General struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Network struct {
	Domain  string `json:"domain"`
	Port    int    `json:"port"`
	Timeout int    `json:"timeout"`
}

type SSL struct {
	CRT string `json:"crt"`
	KEY string `json:"key"`
}

func ServiceConf(data string) {
	var settings Service

	_ = json.Unmarshal([]byte(data), &settings)

	_ = os.Setenv("SERVICE_NAME", settings.General.Name)
	_ = os.Setenv("SERVICE_DESC", settings.General.Desc)

	_ = os.Setenv("NETWORK_DOMAIN", settings.Network.Domain)
	_ = os.Setenv("NETWORK_PORT", strconv.Itoa(settings.Network.Port))
	_ = os.Setenv("NETWORK_TIMEOUT", strconv.Itoa(settings.Network.Timeout))

	_ = os.Setenv("CRT", settings.SSL.CRT)
	_ = os.Setenv("KEY", settings.SSL.KEY)
}