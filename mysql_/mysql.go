package mysql_

import (
	"fmt"
	"github.com/shenqingawei/framework/nacos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMysql() error {
	var err error
	username := nacos.ConfigPz.Mysql.Username
	password := nacos.ConfigPz.Mysql.Password
	connectionType := nacos.ConfigPz.Mysql.ConnectionType
	host := nacos.ConfigPz.Mysql.Host
	port := nacos.ConfigPz.Mysql.Port
	databaseName := nacos.ConfigPz.Mysql.DatabaseName
	connectionParameters := nacos.ConfigPz.Mysql.ConnectionParameters

	dsn := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?%v&parseTime=True&loc=Local",
		username, password, connectionType, host, port, databaseName, connectionParameters)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
func AutoMigrateUser() error {
	return Db.AutoMigrate(new(User))
}

type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(11)"`
	Sex   int64  `gorm:"type:tinyint"`
	Phone string `gorm:"type:char(11)"`
}
