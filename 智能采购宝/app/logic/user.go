package logic

import (
	"bysj1/app/model"
	"bysj1/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func UpdateAdminUser(context *gin.Context) {

	// 解析请求中的项目数据
	var updateduser model.User
	if err := context.ShouldBind(&updateduser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户数据"})
		return
	}
	// 调用更新项目的函数
	updateduser.UpdatedTime = time.Now()
	err := model.UpdateAdminUser(updateduser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新用户"})
		return
	}
	fmt.Println(updateduser)
	UserRole := model.UserRole{
		Userid: updateduser.Id,
		Roleid: updateduser.RoleId,
	}
	if err := model.UpdateUserRole(&UserRole); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10007,
			Message: "用户角色关联失败",
		})
		return
	}
	// 返回更新成功的信息
	context.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}

func UpdateUser(context *gin.Context) {
	// 解析请求中的项目数据
	var updateduser model.User
	if err := context.ShouldBindJSON(&updateduser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户数据"})
		return
	}
	// 调用更新项目的函数
	updateduser.UpdatedTime = time.Now()
	err := model.UpdateUser(updateduser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新用户"})
		return
	}
	// 返回更新成功的信息
	context.JSON(http.StatusOK, gin.H{"message": "用户更新成功"})

}
