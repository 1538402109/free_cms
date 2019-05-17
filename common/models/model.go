package models

import (
	"fmt"
	"free_cms/common"
	"free_cms/pkg/util"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var Db *gorm.DB

type Model struct {
	Id        int       `json:"id" form:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" sql:"index"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = beego.AppConfig.String("dbType")
	dbName = beego.AppConfig.String("dbName")
	user = beego.AppConfig.String("user")
	password = beego.AppConfig.String("password")
	host = beego.AppConfig.String("host")
	tablePrefix = beego.AppConfig.String("tablePrefix")

	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	//Db.LogMode(true)

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer Db.Close()
}


func GetMysqlMsg()(mysqlMsg map[string]string){
	mysqlMsg = make(map[string]string)
	var version string
	if err := Db.Raw("select version()").Row().Scan(&version);err!=nil{
		log.Println(err)
	}
	mysqlMsg["version"] = version
	return
}

func (m *Books) AfterUpdate(scope *gorm.Scope) (err error) {
	Db.Create(&AdminLog{Route:common.Fc.Request.URL.String(),
		UserId:common.UserId,
		Ip:int(util.Ip2long(common.Fc.Input.IP())),
		Method:common.Fc.Request.Method,
		Description:fmt.Sprintf("%s修改了表%s id:%d的%s",common.UserId,scope.TableName(),m.Id,fmt.Sprintf("%+v",scope.Value)),
	})
	return
}

func (m *Books) AfterCreate(scope *gorm.Scope) (err error) {
	Db.Create(&AdminLog{Route:common.Fc.Request.URL.String(),
		UserId:common.UserId,
		Ip:int(util.Ip2long(common.Fc.Input.IP())),
		Method:common.Fc.Request.Method,
		Description:fmt.Sprintf("%s添加了表%s id:%d的%s",common.UserId,scope.TableName(),m.Id,fmt.Sprintf("%+v",scope.Value)),
	})
	return
}

func (m *Books) AfterDelete(scope *gorm.Scope) (err error) {
	Db.Create(&AdminLog{Route:common.Fc.Request.URL.String(),
		UserId:common.UserId,
		Ip:int(util.Ip2long(common.Fc.Input.IP())),
		Method:common.Fc.Request.Method,
		Description:fmt.Sprintf("%s删除了表%s id:%d",common.UserId,scope.TableName(),m.Id),
	})
	return
}