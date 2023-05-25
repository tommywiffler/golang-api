package db

import (
	"gorm.io/gorm"

	"golang-api/infra"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Debug().AutoMigrate(
		&infra.User{},
		&infra.Friend{},
	)
	if err != nil {
		log.Fatal(err)
	}

}

func DBLoad(db *gorm.DB) {
	DB.Create(&infra.User{ID: 1234, Name: "Matt", Email: "matt@email.com", Age: 25})
	DB.Create(&infra.User{ID: 5678, Name: "Mark", Email: "mark@email.com", Age: 30})
	DB.Create(&infra.User{ID: 1212, Name: "Luke", Email: "luke@email.com", Age: 20})
	DB.Create(&infra.User{ID: 3434, Name: "John", Email: "john@email.com", Age: 40})
	DB.Create(&infra.User{ID: 1111, Name: "Bill", Email: "bill@email.com", Age: 22})
	DB.Create(&infra.User{ID: 2222, Name: "Will", Email: "will@email.com", Age: 37})
	DB.Create(&infra.Friend{FriendID: 1234, UserID: 1212})
	DB.Create(&infra.Friend{FriendID: 5678, UserID: 3434})
}
