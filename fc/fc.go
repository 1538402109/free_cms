package fc

import (
	"fmt"
	"free_cms/models"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type TabelAttr struct {
	Field   string `gorm:"column:Field"`
	Type    string `gorm:"column:Type"`
	Null    string `gorm:"column:Null"`
	Key     string `gorm:"column:Key"`
	Default string `gorm:"column:Default"`
	Extra   string `gorm:"column:Extra"`
}

func Fc(tableName, modelPath, controllerPath, viewPath string) {
	tablePrefix := beego.AppConfig.String("tablePrefix")
	var tableAttr []TabelAttr
	models.Db.Raw("desc " + tablePrefix + tableName).Scan(&tableAttr)
	//获取数据
	if modelPath != "" {
		CreateModel(modelPath, tableName, tableAttr)
	}

	if controllerPath != "" {
		CreateController(controllerPath, tableName)
	}

	var viewData string
	if viewPath != "" {
		//viewData = createModel(tableAttr, *tableName, *modelPath, "")

		if err := os.MkdirAll(viewPath, 777); err != nil {
			log.Println("视图文件创建失败")
		}

		if err := ioutil.WriteFile(path.Join(modelPath, fmt.Sprintf("%s.go", tableName)), []byte(viewData), os.ModeType); err != nil {
			log.Println(err)
		}
	}

}

//字段类型转换
func getType(typeName string) (str string) {
	if strings.Index(typeName, "bigint") >= 0 || strings.Index(typeName, "int") >= 0 || strings.Index(typeName, "tinyint") >= 0 {
		return "int"
	}
	if strings.Index(typeName, "varchar") >= 0 || strings.Index(typeName, "char") >= 0 {
		return "string"
	}
	if strings.Index(typeName, "datetime") >= 0 || strings.Index(typeName, "time") >= 0 {
		return "time.Time"
	}
	return "string"
}

//字段转换大驼峰，小驼峰
func Hump(v, t string) (new string) {
	field := strings.Split(v, "_")
	if t == "min" {
		for k, v := range field {
			if k > 0 {
				field[k] = strings.Title(v)
			}
		}
	}
	if t == "max" {
		for k, v := range field {
			field[k] = strings.Title(v)
		}
	}
	new = strings.Join(field, "")
	return
}
