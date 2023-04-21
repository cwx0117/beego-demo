package test

import (
	"demo/models"
	"demo/utils/dbUtils"
	"fmt"
	"github.com/smartystreets/assertions"
	"testing"
)

// var modelMap = make(map[string]interface{})
var modelSlice []interface{}

func init() {
	GenerateTable()
}

func GenerateTable() {
	dsn := dbUtils.GetDsn("root", "123456", "127.0.0.1", 3306, "beego-demo")

	modelSlice = append(modelSlice, models.Org{}, models.User{})
	db, err := dbUtils.AutoMigrate(dsn, modelSlice)

	org := models.Org{
		Province: "广东省",
		City:     "深圳市",
		County:   "南山区",
		AreaCode: "44030500000000000000000000000000000000000000000001",
	}

	user := models.User{
		Org:   org,
		Name:  "John",
		Pwd:   "password",
		Email: "john@example1.com",
		Phone: "13800138001",
	}

	err = db.Create(&user).Error
	if err != nil {
		// 处理错误
		panic(err)
	}

}

func TestM(t *testing.T) {
	fmt.Println("程序启动，测试加载init方法顺序")

}

func TestAutoMigrate(t *testing.T) {
	dsn := dbUtils.GetDsn("root", "123456", "127.0.0.1", 3306, "beego-demo")
	_, err := dbUtils.AutoMigrate(dsn, modelSlice)
	if err != nil {
		panic(err)
	}
}

func TestGetConnPool(t *testing.T) {
	source := "root:123456@tcp(127.0.0.1:3306)/beego-demo?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := dbUtils.GetDsn("root", "123456", "127.0.0.1", 3306, "beego-demo")
	fmt.Println(assertions.ShouldEqual(dsn, source))
	pool, err := dbUtils.GetConnPool(dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println(pool)
}

func TestNewDbObj(t *testing.T) {
	var p = dbUtils.NewDbObj("root", "123456", "127.0.0.1", 3306, `beego-demo`)
	fmt.Println(p.Dsn)
}
