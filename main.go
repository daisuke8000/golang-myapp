package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"myapp/db"
	"os"
	"strconv"
)

func init() {
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
	db.Ow = db.Owner{rdbm, dbnm, user, pass}
}

func main() {
	//インスタンス初期化
	r := gin.Default()
	//src/templates
	r.LoadHTMLGlob("src/templates/*.html")
	//外部参照は大文字から
	db.DbInit()

	//index
	r.GET("/", func(c *gin.Context) {
		tasks := db.GetAll()
		c.HTML(200, "index.html", gin.H{
			"tasks": tasks,
		})
	})

	//create
	r.POST("/new", func(c *gin.Context) {
		title := c.PostForm("title")
		detail := c.PostForm("detail")
		name := c.PostForm("name")
		d := c.PostForm("day")
		day, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		db.Insert(title, detail, name, day)
		c.Redirect(302, "/")
	})

	//detail
	r.GET("/detail/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		task := db.GetOne(id)
		c.HTML(200, "detail.html", gin.H{
			"task": task,
		})
	})

	r.Run(":8080")
}
