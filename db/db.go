package db

import (
	"github.com/jinzhu/gorm"
)

type Owner struct {
	DbName  string
	DbTable string
	DbUser  string
	DbPass  string
}

type Task struct {
	gorm.Model
	Title  string
	Detail string
	Name   string
	Day    int
}

var Ow Owner
var oneline string

//外部package参照はmethod大文字
func DbInit() {
	oneline = Ow.DbUser + ":" + Ow.DbPass + "@/" + Ow.DbTable + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(Ow.DbName, oneline)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Task{})
}

func GetAll() []Task {
	db, err := gorm.Open(Ow.DbName, oneline)
	if err != nil{
		panic(err.Error())
	}
	var tasks []Task
	db.Order("created_at desc").Find(&tasks)
	db.Close()
	return tasks
}

func GetOne(id int) Task {
	db,err := gorm.Open(Ow.DbName, oneline)
	if err != nil{
		panic(err.Error())
	}
	var task Task
	db.First(&task, id)
	db.Close()
	return task
}

func Insert(title string, detail string, name string, day int) {
	db, err := gorm.Open(Ow.DbName, oneline)
	if err != nil{
		panic(err.Error())
	}
	db.Create(&Task{Title: title, Detail: detail, Name: name, Day: day})
	defer db.Close()
}