package dbUtils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type DbObj struct {
	UserName string
	Pwd      string
	Ip       string
	Port     int
	DbName   string
	Dsn      string
}

func NewDbObj(userName string, pwd string, ip string, port int, dbName string) *DbObj {
	dsn := userName + ":" + pwd + "@tcp(" + ip + ":" + strconv.Itoa(port) + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return &DbObj{UserName: userName, Pwd: pwd, Ip: ip, Port: port, DbName: dbName, Dsn: dsn}
}

func GetDsn(userName string, pwd string, ip string, port int, dbName string) string {
	return userName + ":" + pwd + "@tcp(" + ip + ":" + strconv.Itoa(port) + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func GetConnPool(dsn string) (gorm.ConnPool, error) {
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db.ConnPool, nil
}
func AutoMigrate(dsn string, ms []interface{}) (*gorm.DB, error) {
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	for _, value := range ms {
		err := db.AutoMigrate(value)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
