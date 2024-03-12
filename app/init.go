package app

import (
	"github.com/shenqingawei/framework/mysql_"
	"github.com/shenqingawei/framework/nacos"
	"github.com/shenqingawei/framework/nacos_listen"
)

func Init(viperPath string, types ...string) error {
	var err error
	err = nacos.InitNaocs(viperPath) //todo:nacos 配置连接
	if err != nil {
		return err
	}
	for _, v := range types {
		switch v {
		case "mysql":
			err = mysql_.InitMysql() //todo:连接 mysql
		case "auto_migrate_user":
			err = mysql_.AutoMigrateUser() //todo:自动迁移
		}
	}
	err = nacos_listen.NacosListen(viperPath) //todo:动态监听 nacos 配置
	return err
}
