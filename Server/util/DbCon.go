package util

import (
	"encoding/json"
	_ "github.com/bmizerany/pq"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
)


func init(){
	path := GetCurrentDirectory()
	path = path + "dbCon.json"
	js,err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	var db []DbConData
	json.Unmarshal(js,&db)

}


/**
配置文件的对象
 */
type DbConData struct {
	Name string `json:"name"` //别名
	Data string `json:"data"` //数据库名称
	User string `json:"user"` //用户名
	Dbname string `json:"dbname"` //表名称
	Password string `json:"password"` //密码
	Port string `json:"port"` //端口
	Sslmode string `json:"sslmode"` //是否启用ssl连接
	Host string `json:"host"` //主机名
}

func(b *DbConData)InitConnection()(con map[string]*sqlx.DB,err error){
	sqlx.Open(b.Data,"user="+b.User+" dbname="+b.Dbname+" password="+b.Password+" port="+b.Port+" sslmode="+b.Sslmode+" host="+b.Host)
}