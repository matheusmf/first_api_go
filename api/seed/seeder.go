package seed

import (
	"log"

	"first-api-go/api/models"
	"first-api-go/api/utils/helper"

	"github.com/jinzhu/gorm"
)

var users = []models.User{

	models.User{
		ID:       helper.GetIdFromString("9c5151fd-8711-443c-a4f3-b977e5eae1e6"),
		Nickname: "QA",
		Email:    "qa@gmail.com",
		Password: "password",
	},
	models.User{
		ID:       helper.GetIdFromString("1502f3c2-c093-4292-b928-965fdc68d0d4"),
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		ID:          helper.GetIdFromString("951919bc-03fe-42a6-827c-cebffaec82d2"),
		Title:       "Title 1",
		Description: "Description 1",
		Content:     "Hello world 1",
	},
	models.Post{
		ID:          helper.GetIdFromString("77327ea0-7317-4356-bbaa-53c0833d01fc"),
		Title:       "Title 2",
		Description: "Description 2",
		Content:     "Hello world 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*
		err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
