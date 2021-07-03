package repository

import (
	"fmt"
	"log"

	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDataBaseConnectionWithTables(MyUser string, Password string, Host string, Port int, Db string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.Debug().DropTableIfExists(&model.Guest{})
	db.Debug().DropTableIfExists(&model.GuestArrivals{})
	db.Debug().DropTableIfExists(&model.Table{})
	db.Debug().AutoMigrate(&model.Guest{})
	db.Debug().AutoMigrate(&model.Table{})
	db.Debug().AutoMigrate(&model.GuestArrivals{})
	return db
}

func GetDataBaseConnectionWithTablesAndData(MyUser string, Password string, Host string, Port int, Db string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")
	db.Debug().DropTableIfExists(&model.Guest{})
	db.Debug().DropTableIfExists(&model.GuestArrivals{})
	db.Debug().DropTableIfExists(&model.Table{})
	db.Debug().AutoMigrate(&model.Guest{})
	db.Debug().AutoMigrate(&model.Table{})
	db.Debug().AutoMigrate(&model.GuestArrivals{})
	table1 := model.Table{ID: 2, Guests: make([]model.Guest, 0, 10), Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
	table2 := model.Table{ID: 4, Guests: make([]model.Guest, 0, 20), Capacity: 20, Sizeofguests: 0, Emptyseats: 20}
	table3 := model.Table{ID: 6, Guests: make([]model.Guest, 0, 10), Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
	table4 := model.Table{ID: 8, Guests: make([]model.Guest, 0, 20), Capacity: 20, Sizeofguests: 0, Emptyseats: 20}
	table5 := model.Table{ID: 10, Guests: make([]model.Guest, 0, 15), Capacity: 15, Sizeofguests: 0, Emptyseats: 15}
	db.Create(&table1)
	db.Create(&table2)
	db.Create(&table3)
	db.Create(&table4)
	db.Create(&table5)
	return db
}
