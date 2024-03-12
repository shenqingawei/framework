package nacos_listen

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/shenqingawei/framework/mysql_"
	"github.com/shenqingawei/framework/nacos"
	"github.com/shenqingawei/framework/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NacosListen() error {
	c, err := nacos.ConnectionNacos()
	if err != nil {
		return err
	}
	return c.ListenConfig(vo.ConfigParam{
		DataId: viper.NacosConfig.Name,
		Group:  viper.NacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			db, err := mysql_.Db.DB()
			if err != nil {
				panic(err)
			}
			if db != nil {
				err := db.Close()
				if err != nil {
					panic(err)
				}
			}

			username := nacos.ConfigPz.Mysql.Username
			password := nacos.ConfigPz.Mysql.Password
			connectionType := nacos.ConfigPz.Mysql.ConnectionType
			host := nacos.ConfigPz.Mysql.Host
			port := nacos.ConfigPz.Mysql.Port
			databaseName := nacos.ConfigPz.Mysql.DatabaseName
			connectionParameters := nacos.ConfigPz.Mysql.ConnectionParameters

			dsn := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?%v&parseTime=True&loc=Local",
				username, password, connectionType, host, port, databaseName, connectionParameters)
			mysql_.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			err = mysql_.Db.AutoMigrate(new(mysql_.User))
			if err != nil {
				panic(err)
			}
		},
	})
}
