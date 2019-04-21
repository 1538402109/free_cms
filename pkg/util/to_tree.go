package util

import (
	"encoding/json"
	"fmt"
	"free_cms/models"
)

type Cat struct {
	Id       int
	Pid      int
	Name     string
	Children []*Cat `orm:"-"`
}


// var list []BookType
func GetTree(model *[]interface{}){

	//models2 := reflect.New(reflect.TypeOf(model))
	models.Db.Find(model)
	//t := fmt.Sprintf("%T",model)
	//fmt.Println(reflect.TypeOf(model))
/*	 for k,v := range model{
	 	fmt.Println(k,v)
	 }*/
	//fmt.Println(reflect.ValueOf(model).Elem().FieldByName("Id"))
	//typeOfCat  := reflect.TypeOf(model).Elem()
	//fmt.Println(typeOfCat)

	//fmt.Printf("name:'%v' kind:'%v'\n",typeOfCat.Name(), typeOfCat.Kind())
}


func Tree(list []*Cat) string {
	data := buildData(list)
	result := makeTreeCore(0, data)
	body, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func buildData(list []*Cat) map[int]map[int]*Cat {
	var data map[int]map[int]*Cat = make(map[int]map[int]*Cat)
	for _, v := range list {
		id := v.Id
		fid := v.Pid
		if _, ok := data[fid]; !ok {
			data[fid] = make(map[int]*Cat)
		}
		data[fid][id] = v
	}
	return data
}



func makeTreeCore(index int, data map[int]map[int]*Cat) []*Cat {
	tmp := make([]*Cat, 0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.Children = makeTreeCore(id, data)
		}
		tmp = append(tmp, item)
	}
	return tmp
}
