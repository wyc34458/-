package model

import "fmt"

func GetUser(name string) *User {
	var ret User
	if err := Conn.Raw("select * from user where name= ? limit 1", name).Scan(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return &ret
}
func CreateUser(user *User) error {
	if err := Conn.Create(user).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}
func CreateUserRole(userRole *UserRole) error {
	if err := Conn.Create(userRole).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}
func UpdateAdminUser(user User) error {
	// 执行 SQL 语句
	if err := Conn.Exec("UPDATE user SET role_id = ? WHERE id = ?", user.RoleId, user.Id).Error; err != nil {
		return err
	}
	return nil
}
func UpdateUser(user User) error {
	// 执行 SQL 语句
	if err := Conn.Exec("UPDATE user SET name = ?, password = ?, phone = ? WHERE id = ?", user.Name, user.Password, user.Phone, user.Id).Error; err != nil {
		return err
	}
	return nil
}
func UpdateUserRole(userRole *UserRole) error {
	// 执行 SQL 更新语句，将 userrole 表中与给定用户相关的角色ID更新为新的角色ID
	result := Conn.Exec("UPDATE userrole SET roleid = ? WHERE userid = ?", userRole.Roleid, userRole.Userid)
	if result.Error != nil {
		// 根据需要适当处理错误（例如，记录错误日志或返回自定义的错误信息）
		return result.Error
	}
	return nil
}
func GetPermissionByRoleID(roleid int64) string {
	var ret string
	// 执行 SQL 语句
	if err := Conn.Raw("SELECT permission.`name` FROM permission JOIN rolepermission ON permission.id = rolepermission.roleid WHERE roleid = ?", roleid).Scan(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}
