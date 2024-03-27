package model

import (
	"bysj1/app/tools"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Items struct {
	Id          int64  `json:"id" form:"id"`
	UID         int64  `json:"uid" form:"uid"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Budget      int64  `json:"budget" form:"budget"`
	Publisher   string `json:"publisher" form:"publisher"`
	FileUrl     string `json:"file_url" form:"file_url"`
	StatusName  string `json:"statusName" form:"statusName"`
	Because     string `json:"because" form:"because"`
}

func GetItemCache(c context.Context) []Items {
	ret := make([]Items, 0)
	//先查Redis
	key := fmt.Sprintf("所有项目")
	str, err := Rdb.Get(c, key).Result()
	if len(str) > 0 && err == nil {
		fmt.Printf("不回溯数据库! \n")
		_ = json.Unmarshal([]byte(str), &ret)
		return ret
	}
	if err := Conn.Raw("SELECT item.id,item.uid, item.name,item.description, item.publisher, status_name.status_name,item.because FROM item JOIN status_name ON item.status = status_name.status_id").Scan(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	retStr, _ := json.Marshal(ret)
	err = Rdb.Set(c, key, retStr, 3600*time.Second).Err()
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}
func GetItemCaches(c context.Context) []Item {
	ret := make([]Item, 0)
	// 先查 Redis
	key := fmt.Sprintf("未审核")
	str, err := Rdb.Get(c, key).Result()
	if len(str) > 0 && err == nil {
		fmt.Printf("不回溯数据库! \n")
		_ = json.Unmarshal([]byte(str), &ret)
		return ret
	}
	// 回溯数据库
	query := fmt.Sprintf("SELECT * FROM item WHERE status = 0 ")
	if err := Conn.Raw(query).Scan(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	retStr, _ := json.Marshal(ret)
	err = Rdb.Set(c, key, retStr, 3600*time.Second).Err()
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}
func GetItemCachess(c context.Context) []Item {
	ret := make([]Item, 0)
	// 先查 Redis
	key := fmt.Sprintf("一次审核")
	str, err := Rdb.Get(c, key).Result()
	if len(str) > 0 && err == nil {
		fmt.Printf("不回溯数据库! \n")
		_ = json.Unmarshal([]byte(str), &ret)
		return ret
	}
	// 回溯数据库
	query := fmt.Sprintf("SELECT * FROM item WHERE status = 2  ")
	if err := Conn.Raw(query).Scan(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	retStr, _ := json.Marshal(ret)
	err = Rdb.Set(c, key, retStr, 3600*time.Second).Err()
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}
func GetItemCachesss(c *gin.Context) []Items {
	ret := make([]Items, 0)

	userNamestr, ok := c.Get("user")
	if !ok {
		c.JSON(401, tools.ECode{
			Code:    0,
			Message: "失败",
		})
	}
	userName := userNamestr.(*UserToken)
	// 2. 使用用户名在数据库中查询用户发布的所有项目
	if err := Conn.Raw("SELECT item.id,item.uid, item.name,item.description, item.publisher, status_name.status_name,item.because FROM item JOIN status_name ON item.status = status_name.status_id WHERE publisher = ?", userName.Name).Scan(&ret).Error; err != nil {
		fmt.Printf("err: %s", err.Error())
	}

	// 3. 缓存查询结果到 Redis
	key := fmt.Sprintf("个人项目:%s", userName)
	retStr, _ := json.Marshal(ret)
	err := Rdb.Set(c, key, retStr, 3600*time.Second).Err()
	if err != nil {
		fmt.Printf("err: %s", err.Error())
	}

	return ret
}

type StatusCount struct {
	StatusName string
	Count      int
}

func StatusNum(c context.Context) []StatusCount {
	ret := make([]StatusCount, 0)
	// 先查 Redis
	key := fmt.Sprintf("阶段总结")
	str, err := Rdb.Get(c, key).Result()
	if len(str) > 0 && err == nil {
		_ = json.Unmarshal([]byte(str), &ret)
		return ret
	}
	// 回溯数据库
	query := fmt.Sprintf("SELECT status_name.status_name, COUNT(*) AS count FROM item  JOIN status_name ON status_name.status_id=item.status GROUP BY status_name")
	if err := Conn.Raw(query).Scan(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	retStr, _ := json.Marshal(ret)
	err = Rdb.Set(c, key, retStr, 3600*time.Second).Err()
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	fmt.Println(ret)
	return ret
}
func Details(id string) *Items {
	var ret Items
	if err := Conn.Raw("SELECT item.id,item.file_url,item.budget,item.name,item.description, item.publisher, status_name.status_name,item.because FROM item JOIN status_name ON item.status = status_name.status_id WHERE item.id = ?", id).Scan(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return &ret
}
