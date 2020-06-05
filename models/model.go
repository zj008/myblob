package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"myblob/pkg/setting"
	"time"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"column(id)" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"update_at"`
}

func init()  {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlUser,
		setting.MysqlPassword,
		setting.MysqlHost,
		setting.MysqlDb,
	))
	if err != nil{
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
}