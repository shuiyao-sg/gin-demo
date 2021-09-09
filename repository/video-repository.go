package repository

import (
	"demo/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("Fail to connect database")
	}
	err = db.AutoMigrate(&entity.Video{}, &entity.Person{})
	if err != nil {
		panic("Automigrate error")
	}
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	//err := db.connection.Close()
	//if err != nil {
	//	panic("Fail to close database")
	//}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}
func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}
func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}
func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}

