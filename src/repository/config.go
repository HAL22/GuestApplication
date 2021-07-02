package repository

import (
	"fmt"
	"log"

	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func getDataBaseConnectionWithTables(MyUser string, Password string, Host string, Port int, Db string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.Debug().DropTableIfExists(&model.Guest{})
	db.Debug().DropTableIfExists(&model.GuestArrivals{})
	db.Debug().DropTableIfExists(&model.Table{})
	db.Debug().AutoMigrate(&model.Guest{}, &model.GuestArrivals{}, &model.Table{})
	return db
}
