package d
//普通json格式
func ReturnJson(code int,msg string,data ...interface{})(jsonData map[string]interface{}){
	jsonData = make(map[string]interface{}, 3)
	jsonData["code"] = code
	jsonData["msg"] = msg
	if len(data) > 0 && data[0] !=nil{
		jsonData["data"] = data[0]
	}
	return
}
//layui 后台返回需要的json格式
func LayuiJson(code int,msg string,data ...interface{})(jsonData map[string]interface{}){
	jsonData = make(map[string]interface{}, 3)
	jsonData["code"] = code
	jsonData["msg"] = msg
	if len(data) > 1 {
		jsonData["count"] = data[1]
	}
	if len(data)>0{
		jsonData["data"] = data[0]
	}
	return
}

//bootstrap table 返回json
func TableJson(data interface{},col... interface{})(jsonData map[string]interface{}){
	jsonData = make(map[string]interface{}, 3)
	jsonData["rows"] = data
	if len(col)>0{
		jsonData["offset"] = col[0]
		jsonData["limit"] = col[1]
		jsonData["total"] = col[2]
	}
	return
}