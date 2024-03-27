package model

import "time"

type UserRole struct {
	Id     int64 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Userid int64 `gorm:"column:userid;type:bigint(20)" json:"userid"`
	Roleid int64 `gorm:"column:roleid;type:bigint(20)" json:"roleid"`
}

func (m *UserRole) TableName() string {
	return "userrole"
}

type User struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Uid         int64     `gorm:"column:uid;type:bigint(20)" json:"uid" `
	Name        string    `gorm:"column:name;type:varchar(255)" json:"name" form:"name"`
	Password    string    `gorm:"column:password;type:varchar(255)" json:"password" form:"password"`
	Phone       string    `gorm:"column:phone;type:varchar(255);comment:联系方式" json:"phone" form:"phone"`
	RoleId      int64     `gorm:"column:role_id;type:varchar(255)" json:"role_id"`
	CreatedTime time.Time `gorm:"column:created_time;type:datetime" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:datetime" json:"updated_time"`
}

func (m *User) TableName() string {
	return "user"
}

type RolePermission struct {
	Id           int64 `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
	Roleid       int64 `gorm:"column:roleid;type:bigint(20)" json:"roleid"`
	Permissionid int64 `gorm:"column:permissionid;type:bigint(20)" json:"permissionid"`
}

func (m *RolePermission) TableName() string {
	return "rolepermission"
}

type Role struct {
	Id   int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
}

func (m *Role) TableName() string {
	return "role"
}

type Permission struct {
	Id           int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Name         string `gorm:"column:name;type:varchar(255)" json:"name"`
	Permissionid string `gorm:"column:Permissionid;type:varchar(255)" json:"Permissionid"`
}

func (m *Permission) TableName() string {
	return "permission"
}

type Item struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:项目ID" json:"id" form:"id"`
	Uid         string    `gorm:"column:uid;type:bigint(20)" json:"uid" form:"uid"`
	Name        string    `gorm:"column:name;type:varchar(255);comment:项目名称" json:"name" form:"name"`
	Description string    `gorm:"column:description;type:varchar(255);comment:项目描述" json:"description" form:"description"`
	Budget      int64     `gorm:"column:budget;type:varchar(255);comment:预算" json:"budget" form:"budget"`
	Publisher   string    `gorm:"column:publisher;type:varchar(255);comment:发布人" json:"publisher" form:"publisher"`
	FileURL     string    `gorm:"column:file_url;type:varchar(255);comment:文件地址" json:"file_url" form:"file_url"`
	Status      int64     `gorm:"column:status;type:varchar(255);comment:状态  0.待审批1.通过 2.未通过" json:"status" form:"status"`
	Because     string    `gorm:"column:because;type:varchar(255);comment:原因" json:"because" form:"because"`
	CreatedTime time.Time `gorm:"column:created_time;type:datetime;comment:创建时间" json:"created_time" form:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:datetime;comment:更新时间" json:"updated_time" form:"updated_time"`
}

func (m *Item) TableName() string {
	return "item"
}

type CUser struct {
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	Password2 string `json:"password2" form:"password2"`
	Phone     string `json:"phone" form:"phone"`
}
type StatusName struct {
	Id         int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	StatusId   int    `gorm:"column:status_id;type:int(11)" json:"status_id"`
	StatusName string `gorm:"column:status_name;type:varchar(255)" json:"status_name"`
}

func (m *StatusName) TableName() string {
	return "status_name"
}
