package logic

import (
	"bysj1/app/model"
	"bysj1/app/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func GetItem(context *gin.Context) {
	ret := model.GetItemCache(context)
	context.JSON(http.StatusOK, tools.ECode{
		Data: ret,
	})
}
func GetItems(context *gin.Context) {
	ret := model.GetItemCaches(context)
	context.JSON(http.StatusOK, tools.ECode{
		Data: ret,
	})
}
func GetItemss(context *gin.Context) {
	ret := model.GetItemCachess(context)
	context.JSON(http.StatusOK, tools.ECode{
		Data: ret,
	})
}
func GetItemsss(context *gin.Context) {
	ret := model.GetItemCachesss(context)
	context.JSON(http.StatusOK, tools.ECode{
		Data: ret,
	})
}

func Details(context *gin.Context) {
	id := context.Param("id")
	details := model.Details(id)
	context.JSON(http.StatusOK, tools.ECode{
		Data: details,
	})

}

type ItemDetail struct {
	Name        string `json:"name"`
	UID         string `json:"uid"`
	Description string `json:"description"`
	// 其他项目详情字段
}

// 处理"查看详情"请求的API端点
//
//	func GetItemDetailHandler(c *gin.Context) {
//		uid := c.Param("uid") // 从URL中获取项目的UID
//		itemDetail := model.GetItemCachessss(c)
//		// 根据项目的UID筛选出对应的项目详情
//		var detail *ItemDetail
//		for _, item := range itemDetail {
//			if item.Uid == uid {
//				detail = &ItemDetail{
//					Name:        item.Name,
//					UID:         item.Uid,
//					Description: item.Description,
//					// 可根据需要添加其他项目详情字段
//				}
//				break
//			}
//		}
//
//		if detail != nil {
//			c.JSON(http.StatusOK, detail) // 返回项目详情
//		} else {
//			c.JSON(http.StatusNotFound, gin.H{"error": "项目详情未找到"}) // 返回项目详情未找到的错误
//		}
//	}
func DelItem(context *gin.Context) {
	projectID := context.Query("id")
	id, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		// 处理无效的项目ID
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的项目ID",
		})
		return
	}
	// 从数据库中获取要删除的项目
	project := model.GetItemId(id)
	if project.Id <= 0 {
		// 处理项目未找到的情况
		context.JSON(http.StatusNotFound, gin.H{
			"error": "找不到项目",
		})
		return
	}

	// 执行删除操作，将项目从数据库中删除
	if err := model.DeleteProject(id); err != true {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "删除失败",
		})
		return
	}
	// 返回删除成功的信息
	context.JSON(http.StatusOK, gin.H{
		"message": "项目删除成功",
	})
	return
}
func AddItem(context *gin.Context) {
	var project model.Item
	if err := context.ShouldBind(&project); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	if project.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "采购项目名称不能为空"})
		return
	}

	// 处理上传的文件
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "获取文件失败",
		})
		return
	}
	targetDir := "D:/毕业设计上传文档" // 目标文件夹路径

	// 创建目标文件
	dst, err := os.Create(filepath.Join(targetDir, file.Filename))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}
	defer dst.Close()

	// 将上传的文件拷贝到目标文件
	if err := context.SaveUploadedFile(file, filepath.Join(targetDir, file.Filename)); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	newProject := model.Item{
		Uid:         strconv.FormatInt(tools.GetUid(), 10),
		Name:        project.Name,                            //项目名称
		Description: project.Description,                     //项目描述
		Budget:      project.Budget,                          //预算
		Publisher:   project.Publisher,                       //发布人
		FileURL:     filepath.Join(targetDir, file.Filename), // 保存文件地址
		Status:      0,                                       //状态
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	// 保存采购项目到数据库
	if err := model.AddItem(newProject); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "采购项目创建失败",
		})
		return
	}

	// 返回创建成功
	context.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "采购项目创建成功",
	})
}
func UpdateItem(context *gin.Context) {
	idStr := context.Param("id") // 假设ID以路径参数的形式传递，比如 "/items/:id"
	// 将字符串ID转换为int64类型
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	// 解析请求中的项目数据
	var updatedProject model.Item
	if err := context.ShouldBind(&updatedProject); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的项目数据"})
		return
	}

	// 处理上传的新文件
	newFile, err := context.FormFile("file")
	if err != nil {
		// 如果没有新文件上传，则继续更新项目信息，不进行文件处理
	} else {
		// 有新文件上传，处理新文件

		// 确定目标文件夹路径
		targetDir := "D:/毕业设计上传文档" // 替换为你的目标文件夹路径

		// 创建目标文件
		dst, err := os.Create(filepath.Join(targetDir, newFile.Filename))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
			return
		}
		defer dst.Close()

		// 将上传的新文件拷贝到目标文件
		if err := context.SaveUploadedFile(newFile, filepath.Join(targetDir, newFile.Filename)); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
			return
		}

		// 更新项目信息中的文件路径
		updatedProject.FileURL = filepath.Join(targetDir, newFile.Filename)
	}
	// 更新项目信息
	updatedProject.UpdatedTime = time.Now()
	err = model.UpdateItem(id, updatedProject)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新项目"})
		return
	}

	// 返回更新成功的信息
	context.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "更新成功",
	})
}

func UpdateAdminItem(context *gin.Context) {

	userNamestr, ok := context.Get("user")
	if !ok {
		context.JSON(401, tools.ECode{
			Code:    0,
			Message: "失败",
		})
	}
	userToken, ok := userNamestr.(*model.UserToken)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效的用户令牌",
		})
		return
	}

	// 从JWT中获取角色ID
	roleID := userToken.Roleid

	// 查询角色权限表获取权限ID
	permissionName := model.GetPermissionByRoleID(roleID)

	// 查询权限表获取权限名称

	// 判断权限名称是否为"update_item1"
	if permissionName != "update_item1" {
		context.JSON(http.StatusForbidden, gin.H{
			"error": "无权限",
		})
		return
	}

	idStr := context.Param("id") // 假设ID以路径参数的形式传递，比如 "/items/:id"
	// 将字符串ID转换为int64类型
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	// 解析请求中的项目数据
	var updatedProject model.Item
	if err := context.ShouldBind(&updatedProject); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的项目数据"})
		return
	}
	// 调用更新项目的函数
	updatedProject.UpdatedTime = time.Now()
	err = model.UpdateAdminItem(id, updatedProject)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新项目"})
		return
	}
	// 返回更新成功的信息
	context.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "审核成功",
	})
}

func UpdateSecondAdminItem(context *gin.Context) {
	userNamestr, ok := context.Get("user")
	if !ok {
		context.JSON(401, tools.ECode{
			Code:    0,
			Message: "失败",
		})
	}
	userToken, ok := userNamestr.(*model.UserToken)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "无效的用户令牌",
		})
		return
	}

	// 从JWT中获取角色ID
	roleID := userToken.Roleid

	// 查询角色权限表获取权限ID
	permissionName := model.GetPermissionByRoleID(roleID)

	// 查询权限表获取权限名称

	// 判断权限名称是否为"update_item1"
	if permissionName != "update_item2" {
		context.JSON(http.StatusForbidden, gin.H{
			"error": "无权限",
		})
		return
	}

	idStr := context.Param("id") // 假设ID以路径参数的形式传递，比如 "/items/:id"
	// 将字符串ID转换为int64类型
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	// 解析请求中的项目数据
	var updatedProject model.Item
	if err := context.ShouldBind(&updatedProject); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "无效的项目数据"})
		return
	}

	// 调用更新项目的函数
	updatedProject.UpdatedTime = time.Now()
	err = model.SecondUpdateAdminItem(id, updatedProject)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新项目"})
		return
	}
	// 返回更新成功的信息
	context.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "采购项目创建成功",
	})
}

type ResultVoteOpt struct {
	Name  string
	Count int64
}
type ResultData struct {
	Title string
	Count int64
	Opt   []*ResultVoteOpt
}

func StatusNum(context *gin.Context) {
	ret := model.StatusNum(context)
	context.JSON(0, tools.ECode{
		Data: ret,
	})
}
