package main

import (
	"github.com/gin-gonic/gin"

	"github.com/songdony/vicuna-redis/lib"
)

func main()  {
	r := gin.New()
	r.Use(func(context *gin.Context) {
		defer func() {
			if e:= recover();e!=nil{
				context.JSON(400,gin.H{"message":e})
			}
		}()
		context.Next()
	})
	r.Handle("GET","/news/:id", func(context *gin.Context) {
		// 对象池获取
		newsCache := lib.NewsCache()
		defer lib.ReleaseNewsCache(newsCache)
		// key: news123 news101
		newsID := context.Param("id")
		newsCache.DBGetter = lib.NewsDBGetter(newsID)

		//context.Header("Content-type","application/json")
		//context.String(200,newsCache.GetCache("news"+newsID).(string))
		newsModel := lib.NewNewsModel()
		newsCache.GetCacheForObject("news"+newsID,newsModel)
		context.JSON(200,newsModel)
	})

	r.Run(":8080")

}