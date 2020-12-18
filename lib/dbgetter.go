package lib

import (
	"log"

	"github.com/songdony/vicuna-redis/go_redis"
)

func NewsDBGetter(id string) go_redis.DBGetterFunc{

	return func() interface{} {
		log.Println("get from db")
		newsModel:=NewNewsModel()
		Gorm.Table("mynews").Where("id=?",id).Find(newsModel)
		return newsModel
	}

}
