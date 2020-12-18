package lib

import (
	"sync"
	"time"

	"github.com/songdony/vicuna-redis/go_redis"
)

var NewsCachePool *sync.Pool

func init(){
	NewsCachePool=&sync.Pool{
		New: func() interface{} {
			return go_redis.NewSimpleCache(go_redis.NewStringOperation(),time.Second*100,
				go_redis.Serilizer_JSON,go_redis.NewCrossPolicy("^news\\d{1,5}$",time.Second*30))   // 指定序列化是 JSON
		},
	}
}

func NewsCache() *go_redis.SimpleCache{
	return NewsCachePool.Get().(*go_redis.SimpleCache)
}

func ReleaseNewsCache(cache *go_redis.SimpleCache){
	NewsCachePool.Put(cache)
}