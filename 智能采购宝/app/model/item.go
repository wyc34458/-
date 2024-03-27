package model

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
)

func AddItem(name Item) error {
	err := Conn.Raw("SELECT * FROM item WHERE name = ?", name.Name).Error
	// 如果存在相同的记录，则返回错误

	err = Conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("item").Create(&name).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
func DelItem(id int64) bool {
	if err := Conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("delete from item where id = ? limit 1", id).Error; err != nil {
			fmt.Printf("err:%s", err.Error())
			return err
		}
		return nil
	}); err != nil {
		fmt.Printf("err:%s", err.Error())
		return false
	}
	return true
}
func GetItemId(id int64) Item {
	var ret Item
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := Conn.Raw("select * from item where id =?", id).Scan(&ret).Error
		if err != nil {
			fmt.Printf("err:%s", err.Error())
		}
	}()
	wg.Wait()
	return ret
}
func DeleteProject(id int64) bool {
	if err := Conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("delete from item where id = ? limit 1", id).Error; err != nil {
			fmt.Printf("err:%s", err.Error())
			return err
		}
		return nil
	}); err != nil {
		fmt.Printf("err:%s", err.Error())
		return false
	}
	return true
}
func UpdateItem(id int64, item Item) error {
	// 执行 SQL 语句
	if err := Conn.Exec("UPDATE item SET name = ?, description = ?, budget = ?,  publisher = ? , file_url = ? ,status = ?  WHERE id = ?", item.Name, item.Description, item.Budget, item.Publisher, item.FileURL, 0, id).Error; err != nil {
		return err
	}
	return nil
}
func UpdateAdminItem(id int64, item Item) error {
	// 执行 SQL 语句
	if err := Conn.Exec("UPDATE item SET status = ?,because = ?  WHERE id = ?", item.Status, item.Because, id).Error; err != nil {
		return err
	}
	return nil
}
func SecondUpdateAdminItem(id int64, item Item) error {
	// 执行 SQL 语句
	if err := Conn.Exec("UPDATE item SET status = ?,because = ?  WHERE  status=? ", item.Status, item.Because, id).Error; err != nil {
		return err
	}
	return nil
}
