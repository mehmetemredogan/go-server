package boot

import (
	"github.com/mehmetemredogan/go-server/internal/array"
	"io/ioutil"
	"log"
	"regexp"
)

func Reader(filepath string) {
	detectConfFileName	:= regexp.MustCompile(`configs/([a-z]+)\.json`)
	getConfFileName		:= detectConfFileName.FindStringSubmatch(filepath)

	confFileList		:= []string{"service"}

	search				:= array.InArray(confFileList, getConfFileName[1])
	if search == -1 {
		log.Fatalln("Configuration file not recognized!")
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln("The configuration file could not be read!")
	}

	switch getConfFileName[1] {
	case "service":
		ServiceConf(string(data))
	}
}