package repository

import (
	"fmt"
	"log"

	"github.com/GG_Backend_tech_challenge/src/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDataBaseConnectionWithTables(MyUser string, Password string, Host string, Port string, Db string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
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

func GetDataBaseConnectionWithTablesAndData(MyUser string, Password string, Host string, Port string, Db string) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", MyUser, Password, Host, Port, Db)
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
	table1 := model.Table{ID: 1, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
	table2 := model.Table{ID: 2, Capacity: 20, Sizeofguests: 0, Emptyseats: 20}
	table3 := model.Table{ID: 3, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
	table4 := model.Table{ID: 4, Capacity: 20, Sizeofguests: 0, Emptyseats: 20}
	table5 := model.Table{ID: 5, Capacity: 15, Sizeofguests: 0, Emptyseats: 15}
	table6 := model.Table{ID: 6, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
	table7 := model.Table{ID: 7, Capacity: 20, Sizeofguests: 0, Emptyseats: 20}
	table8 := model.Table{ID: 8, Capacity: 10, Sizeofguests: 0, Emptyseats: 10}
	table9 := model.Table{ID: 9, Capacity: 20, Sizeofguests: 0, Emptyseats: 20}
	table10 := model.Table{ID: 10, Capacity: 15, Sizeofguests: 0, Emptyseats: 15}
	db.Create(&table1)
	db.Create(&table2)
	db.Create(&table3)
	db.Create(&table4)
	db.Create(&table5)
	db.Create(&table6)
	db.Create(&table7)
	db.Create(&table8)
	db.Create(&table9)
	db.Create(&table10)
	return db
}
