package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

type Task struct {
	gorm.Model
	Title  string
	Detail string
	Name   string
	Day    int
}

type Owner struct {
	DbName  string
	DbTable string
	DbUser  string
	DbPass  string
}


var ow Owner

func dbInit() {
	oneline := ow.DbUser+":"+ow.DbPass+"@/"+ow.DbTable+"?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(ow.DbName, oneline)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Task{})
}

func init()  {
	//.env読み込み
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	// 構造体に格納
	rdbm := os.Getenv("GO_RDBMS")
	dbnm := os.Getenv("GO_DB_NAME")
	user := os.Getenv("GO_DB_USER")
	pass := os.Getenv("GO_DB_PASS")
	ow = Owner{rdbm, dbnm, user, pass}
}


func main() {
	//インスタンス初期化
	r := gin.Default()
	//templates
	r.LoadHTMLGlob("templates/*.html")
	//db-migaration
	dbInit()
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"TableName": ow.DbTable,"UserName": ow.DbUser})
	})

	r.Run(":8080")
}
