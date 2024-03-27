package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rbcervilla/redisstore/v9"
)

var store *redisstore.RedisStore
var sessionName = "session-name"

func GetSession(c *gin.Context) map[interface{}]interface{} {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Printf("session:%+v\n", session.Values)
	return session.Values
}
func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Printf("session : %+v\n", session.Values)
	session.Values["name"] = ""
	session.Values["id"] = int64(0)
	session.Values["uid"] = int64(0)
	session.Values["roleid"] = int64(0)
	return session.Save(c.Request, c.Writer)
}
func SetSession(c *gin.Context, id int64, uid int64, name string, roleid int64) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["id"] = id
	session.Values["uid"] = uid
	session.Values["name"] = name
	session.Values["roleid"] = roleid
	session.Options.MaxAge = 3600
	return session.Save(c.Request, c.Writer)
}
