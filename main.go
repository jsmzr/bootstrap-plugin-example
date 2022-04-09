package main

import (
	"fmt"

	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-log/log"
	_ "github.com/jsmzr/bootstrap-plugin-config-yaml/yaml"
	_ "github.com/jsmzr/bootstrap-plugin-logrus/logrus"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

type ConfigResource struct {
	Config string
	Log    string
}

func main() {
	err := plugin.PostProccess()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Info("pulgin init success %s", "123321")
	log.Info("pulgin init success %d", 123456)
	var configResource ConfigResource
	config.Resolve("bootstrap.example", &configResource)
	log.Info("get config resource %v", configResource)

}
