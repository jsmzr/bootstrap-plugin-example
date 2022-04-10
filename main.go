package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-log/log"
	_ "github.com/jsmzr/bootstrap-plugin-config-yaml/yaml"
	_ "github.com/jsmzr/bootstrap-plugin-logrus/logrus"

	_ "github.com/jsmzr/bootstrap-plugin-redis/redis"
	// _ "github.com/jsmzr/bootstrap-plugin-redis/cluster"
	"github.com/jsmzr/bootstrap-plugin-redis/connection"

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
	clusterConfig, ok := config.Get("bootstrap.redis.cluster.addrs")
	if !ok {
		return
	}
	log.Info("get bootstrap.redis.cluster:%v", clusterConfig)
	var clusterOptions redis.ClusterOptions
	if err := config.Resolve("bootstrap.redis.cluster", &clusterOptions); err != nil {
		fmt.Println(err)
		return
	}
	log.Info("resolve: %v", clusterOptions)

	// conn := connection.GetClutserClient()
	conn := connection.GetClient()
	key := "test-1"
	value := "v1"
	if err := conn.Set(context.Background(), key, value, time.Hour).Err(); err != nil {
		log.Error("redis set faied", err)
		return
	}
	if result, err := conn.Get(context.Background(), key).Result(); err != nil {
		log.Error("redis get failed", err)
	} else {
		log.Info("redis get value: %v", result)
	}

}
