package logic

import (
	"bysj1/app/model"
	"bysj1/app/tools"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
)

func DoLogin(context *gin.Context) {
	var user model.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, tools.ECode{
			Code:    1,
			Message: "登录失败，请检查输入信息",
		})
		return
	}
	ret := model.GetUser(user.Name)
	if ret.Id < 1 || ret.Password != encrypt(user.Password) {
		context.JSON(http.StatusUnauthorized, tools.ECode{
			Code:    1,
			Message: "用户名或密码错误",
		})
		return
	}
	// 登录成功
	// 执行其他逻辑，例如设置会话信息等
	token, _ := model.GetJwt(ret.Id, ret.Name, ret.RoleId)
	context.Header("jwt", token)
	context.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "登录成功",
		Data:    token,
	})
}
func CreateUser(context *gin.Context) {
	var user model.CUser
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10001,
			Message: err.Error(),
		})
		return
	}
	//对数据进行校验
	if user.Name == "" {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10002,
			Message: "账号不能为空",
		})
		return
	}
	if user.Password == "" || user.Password2 == "" {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10002,
			Message: "密码或确认密码不能为空",
		})
		return
	}
	if user.Phone == "" {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10002,
			Message: "手机号不能为空",
		})
		return
	}

	//校验密码
	if user.Password != user.Password2 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10003,
			Message: "两次密码不同！", //这里有风险
		})
		return
	}

	//校验用户是否存在，这种写法非常不安全。有严重的并发风险
	if oldUser := model.GetUser(user.Name); oldUser.Id > 0 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10004,
			Message: "用户名已存在", //这里有风险
		})
		return
	}
	//判断位数
	lenName := len(user.Name)
	lenPwd := len(user.Password)
	if lenName < 8 || lenName > 16 || lenPwd < 8 || lenPwd > 16 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: "用户名或者密码要大于等于8，小于等于16！", //这里有风险
		})
		return
	}
	lenpn := len(user.Phone)
	if lenpn != 11 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10005,
			Message: "手机号错误！", //这里有风险
		})
		return
	}
	//密码不能是纯数字
	regex := regexp.MustCompile(`^[0-9]+$`)
	if regex.MatchString(user.Password) {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "密码不能为纯数字", //这里有风险
		})
		return
	}
	//开始添加用户
	newUser := model.User{
		Uid:         tools.GetUid(),
		Name:        user.Name,
		Password:    encrypt(user.Password),
		Phone:       user.Phone,
		RoleId:      1,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	if err := model.CreateUser(&newUser); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "用户创建失败", //这里有风险
		})
		return
	}
	userRole := model.UserRole{
		Userid: newUser.Id,
		Roleid: newUser.RoleId,
	}
	if err := model.CreateUserRole(&userRole); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10007,
			Message: "用户角色关联失败",
		})
		return
	}

	//返回添加成功
	context.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "用户注册成功",
	})
}
func Logout(context *gin.Context) {
	context.SetCookie("name", "", 3600, "/", "", true, false)
	context.SetCookie("Id", "", 3600, "/", "", true, false)
	context.Redirect(http.StatusFound, "/login")
}
func encrypt(password string) string {
	newPassword := password + "xxxy" //不能随便起，且不能暴露
	hash := md5.New()
	hash.Write([]byte(newPassword))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	fmt.Printf("加密后的密码：%s\n", hashString)
	return hashString
}
