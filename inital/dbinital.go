package inital

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lzhphantom/go-spider/model"
	_ "mysql"
)

var DataBase *gorm.DB

func DBInit() {
	db, err := gorm.Open("mysql", "root:root@/go-spider?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("DB connect failed:", err)
		return
	}
	fmt.Println("database connect success")
	db.LogMode(true)

	db.CreateTable(&model.ChoooRankListInfo{})

	DataBase = db
}
