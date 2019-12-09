package main
import (
	"fmt"
	"github.com/astaxie/beego/orm"
)
var O orm.Ormer
func init(){
	orm.RegisterDataBase("default", "mysql", "root:116118@tcp(www.whatdoyoudo.club:3306)/dataset?charset=utf8", 30)
	// create table
	orm.RegisterModel(new(ModelInit))
	orm.RunSyncdb("default", false, true)
	O=orm.NewOrm()
}

func GroupImageIDGet(Group string)[]orm.ParamsList{
	var res  []orm.ParamsList
	_,err:=O.Raw("select distinct imageID from imageLabelTable where labelName = ? and confidence=1 limit 10;",Group).ValuesList(&res)
	if err != nil {
		fmt.Println(err)
	}
	return res
}
type UrlType struct {
	ImageID string
	Url string
}
func GroupUrlGet(GroupList []orm.ParamsList){

	var res orm.ParamsList
	fmt.Printf("%#v",len(GroupList))
	for i:=0;i<len(GroupList);i++{
		fmt.Printf("%s",GroupList[i][0])
		_,err:=O.Raw("select urlLink from urlTable where imageID=?",GroupList[i][0]).ValuesFlat(&res)
		if err != nil {
			fmt.Printf("%#v",err)
		}
		fmt.Printf("%#v",res)
	}
}