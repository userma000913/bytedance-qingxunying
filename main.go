package main

import (
	"dousheng-backend/dao/mysql"
	"dousheng-backend/dao/redis"
	"dousheng-backend/logger"
	"dousheng-backend/router"
	"dousheng-backend/setting"
	"dousheng-backend/util"
	"fmt"
)

func main() {
	// 配置初始化
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 数据库初始化
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	// redis初始化
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	if err := logger.Init(setting.Conf.LogConfig, "dev"); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 路由初始化
	r := router.InitRouter()
	//oss的bucket初始化
	err := util.GetBuck()
	if err != nil {
		fmt.Printf("bucket get failed, err:%v\n", err)
		return
	}
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
