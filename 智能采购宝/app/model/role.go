package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func GetPermissionid(roleid int64) []int64 {
	var ret []RolePermission
	var res []int64
	if err := Conn.Raw("SELECT permissionid FROM rolepermission WHERE roleid = ?", roleid).Scan(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return []int64{0}
	}

	for _, permission := range ret {
		res = append(res, permission.Permissionid)
	}
	return res // 假设RolePermission结构体中权限ID字段名为PermissionID
}
func GetSecondApproverRoleID(roleid int64) int64 {
	var parentRoleID int64

	// 查询父级角色ID，假设父级角色ID存储在roles表的parent_role_id字段中
	err := Conn.Raw("SELECT parent_role_id FROM role WHERE role_id = ?", roleid).Scan(&parentRoleID).Error
	if err != nil {
		fmt.Printf("Error retrieving parent role ID: %s", err.Error())
		return 0
	}

	// 查询二级审批人角色ID，假设二级审批人角色ID存储在roles表的role_id字段中
	var secondApproverRoleID int64
	err = Conn.Raw("SELECT role_id FROM roles WHERE parent_role_id = ?", parentRoleID).Scan(&secondApproverRoleID).Error
	if err != nil {
		fmt.Printf("Error retrieving second approver role ID: %s", err.Error())
		return 0
	}

	return secondApproverRoleID
}
func GetPermission(permissionid []int64) []string {
	var res []string
	for _, i2 := range permissionid {
		var ret Permission
		Conn.Raw("SELECT name FROM permission WHERE permissionid = ?", i2).Scan(&ret)
		res = append(res, ret.Name)
	}
	return res
}
func AssignProjectToApprover(projectID int64, approverRoleID int64) error {
	// 查询项目是否存在
	project := Item{}
	err := Conn.First(&project, projectID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("project not found")
		}
		return fmt.Errorf("error retrieving project: %s", err.Error())
	}

	// 更新项目的审批人角色ID
	project.Status = approverRoleID
	err = Conn.Save(&project).Error
	if err != nil {
		return fmt.Errorf("error assigning project to approver: %s", err.Error())
	}

	return nil
}
