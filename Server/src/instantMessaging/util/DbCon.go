package util

import (
	"encoding/json"
	_ "github.com/bmizerany/pq"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
)

var DBCon map[string]*sqlx.DB

func init(){
	path := GetCurrentDirectory()
	path = path + "dbCon.json"
	js,err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	var db []DbConData
	//解析配置文件
	json.Unmarshal(js,&db)
	//连接绑定
	establishConnection(db,&DBCon)
}

/**
建立连接，并且把他们绑定到连接对象上
 */
func establishConnection(db []DbConData,Db *map[string]*sqlx.DB){
	for i := 0; i< len(db); i++ {
		//将name做为键，DB连接做为值传给map对象
		(*Db)[db[i].Name] = db[i].InitConnection()
	}
}


/**
配置文件的对象
 */
type DbConData struct {
	Name string `json:"name"` //别名
	Data string `json:"data"` //数据库名称
	User string `json:"user"` //用户名
	Dbname string `json:"dbname"` //库名称
	Password string `json:"password"` //密码
	Port string `json:"port"` //端口
	Sslmode string `json:"sslmode"` //是否启用ssl连接
	Host string `json:"host"` //主机名
}

func(b *DbConData)InitConnection()(con *sqlx.DB){
	db,err := sqlx.Open(b.Data,"user="+b.User+" dbname="+b.Dbname+" password="+b.Password+" port="+b.Port+" sslmode="+b.Sslmode+" host="+b.Host)
	databaseError(err)
	return db
}

func databaseError(err error){
	if err != nil {
		log.Fatal("DB层错误:" + err.Error())
	}
}