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
	"github.com/jsmzr/bootstrap-plugin-redis/connection"

	// _ "github.com/jsmzr/bootstrap-plugin-redis/redis"
	//_ "github.com/jsmzr/bootstrap-plugin-redis/cluster"
	// _ "github.com/jsmzr/bootstrap-plugin-redis/sentinel"

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
	logDemo()
	configDemo()
	// 需要配置本地的 redis 信息
	// redisDemo()
}

func logDemo() {
	log.Info("string value: %s", "123321")
	log.Info("int value: %d", 123456)
	log.Info("bool value: %v", true)
	log.Info("float value: %f", 3.14)
	log.Info("struct value: %v", ConfigResource{
		Config: "foo",
		Log:    "boo",
	})

	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")
}

func configDemo() {
	if v, ok := config.Get("bootstrap.example.int"); ok {
		log.Info("config int value: %d", v.Int())
	}
	if v, ok := config.Get("bootstrap.example.float"); ok {
		log.Info("config float value: %f", v.Float())
	}
	if v, ok := config.Get("bootstrap.example.string"); ok {
		log.Info("config float value: %s", v.String())
	}
	if v, ok := config.Get("bootstrap.example.bool"); ok {
		log.Info("config float value: %v", v.Bool())
	}

	var exampleConfig ConfigResource
	if err := config.Resolve("bootstrap.example", &exampleConfig); err != nil {
		log.Warn("can't resolve bootstrap.example")
	} else {
		log.Info("config struct value: %v", exampleConfig)
	}

	var clusterOptions redis.ClusterOptions
	if err := config.Resolve("bootstrap.redis.cluster", &clusterOptions); err != nil {
		log.Warn("can't resolve bootstrap.redis.cluster")
	} else {
		log.Info("resolve: %v", clusterOptions)
	}
}

func redisDemo() {
	// 需要获取集群连接则在 import 选择导入 github.com/jsmzr/bootstrap-plugin-redis/cluster
	// conn := connection.GetClutserClient()
	// 需要获取单机或哨兵下的连接则在 import 选择导入 github.com/jsmzr/bootstrap-plugin-redis/redis 或 sentinel
	conn := connection.GetClient()
	key := "test-1"
	value := "v1"
	if err := conn.Set(context.Background(), key, value, time.Hour).Err(); err != nil {
		log.Error("redis set faied", err)
		return
	}
	log.Info("redis set %s=%s", key, value)
	if result, err := conn.Get(context.Background(), key).Result(); err != nil {
		log.Error("redis get failed", err)
	} else {
		log.Info("redis get %s value: %s", key, result)
	}
}
