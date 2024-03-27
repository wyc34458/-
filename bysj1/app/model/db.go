package model

import (
	"bysj1/app/tools"
	"context"
	"fmt"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB
var Rdb *redis.Client

func NewMysql() {
	my := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", tools.Configs.MySql.Username, tools.Configs.MySql.Password, tools.Configs.MySql.Host, tools.Configs.MySql.Database)
	fmt.Println(my)
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s\n", err)
		panic(err)
	}
	Conn = conn
}
func NewRdb() {
	rbd := redis.NewClient(&redis.Options{
		Addr:     "192.168.10.20:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Rdb = rbd
	store, _ = redisstore.NewRedisStore(context.TODO(), Rdb)
	return
}
func Close() {
	db, err := Conn.DB()
	if err != nil {
		fmt.Printf("Failed to close database connection: %s\n", err)
	} else {
		_ = db.Close()
	}

	if Rdb != nil {
		err := Rdb.Close()
		if err != nil {
			fmt.Printf("Failed to close Redis connection: %s\n", err)
		}
	}
}
