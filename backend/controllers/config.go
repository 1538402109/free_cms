package controllers

import (
	"encoding/json"
	"free_cms/backend/models"
	models2 "free_cms/common/models"
	"log"
)

type ConfigController struct {
	CommonController
}

func (c *ConfigController) Index() {
	configModel := models.NewConfig()
	if c.Ctx.Input.IsPost(){
		var name map[string]string
		if err := c.Ctx.Input.Bind(&name,"name");err !=nil{
			c.JsonResult(201,"参数错误")
		}
		for k,v:=range name{
			if err := models2.Db.Model(configModel).Where("name=?",k).Update("value",v).Error;err !=nil{
				log.Println(err)
			}
		}
		c.JsonResult(200,"修改成功")
	}
	//分组
	config, _ := configModel.FindByName("configgroup")
	var configGroup map[string]string
	json.Unmarshal([]byte(config.Value), &configGroup)

	//数据
	configs, _ := models.NewConfig().FindAll()

	c.Data["configGroup"] = configGroup
	c.Data["configs"] = configs
	c.TplName = c.ADMIN_TPL + "config/index.html"
}
