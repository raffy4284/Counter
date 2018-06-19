package database

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBHandler struct {
	O  interface{}
	DB *gorm.DB
}

var instantiated *DBHandler
var once sync.Once

func GetDatabase() *DBHandler {
	if instantiated == nil {
		instantiated = new(DBHandler)
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			panic("failed to connect database")
		}
		instantiated.DB = db
	}
	return instantiated
}
