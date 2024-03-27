package router

import (
	"bysj1/app/logic"
	"bysj1/app/model"
	"bysj1/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	index := r.Group("")
	r.Use(Cors())
	index.Use(checkJwt)
	index.Use(Cors())
	r.GET("/index", logic.Index)
	r.GET("/", logic.Index)
	{
		//r.GET("/login", logic.GetLogin)          //访问登录页面
		r.POST("/login", logic.DoLogin) //登录
		r.GET("/logout", logic.Logout)  //用户登出
		//r.GET("/create", logic.GetCreate)        //访问注册页面
		r.POST("/user/create", logic.CreateUser) //注册

	}
	{
		index.POST("/item/add", logic.AddItem)                                  //增加项目
		index.DELETE("/item/delete", logic.DelItem)                             //删除项目
		index.PUT("/item/update/:id", logic.UpdateItem)                         //修改项目
		index.PUT("/item/admin/update/:id", logic.UpdateAdminItem)              //一级审批人审批项目
		index.PUT("/item/admin/second/update/:id", logic.UpdateSecondAdminItem) //二级审批人审批项目

	}
	{
		index.PUT("/user/admin/update", logic.UpdateAdminUser) //最高级管理员修改用户权限
		index.PUT("/user/update", logic.UpdateUser)            //修改用户
	}
	{
		index.GET("/item/all", logic.GetItem)         //查所有项目信息
		index.GET("/item/first", logic.GetItems)      //查所有未审批的项目信息
		index.GET("/item/second", logic.GetItemss)    //查所有初次审核通过的项目信息
		index.GET("/item/one", logic.GetItemsss)      //查登录用户的项目信息
		index.GET("/item/details/:id", logic.Details) //查项目的详细信息
		index.GET("/item/chart", logic.StatusNum)     //阶段统计
	}
	go func() {
		if err := r.Run(":8080"); err != nil {
			panic("gin 启动失败!")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// 模拟程序的工作
		for {
			select {
			case <-quit:
				// 收到退出信号，执行清理操作
				fmt.Println("接收到信号。正在清理。。。")
				// TODO: 执行额外的清理操作
				time.Sleep(2 * time.Second) // 模拟清理操作的耗时

				// 退出程序
				fmt.Println("优雅地退出")
				os.Exit(0)
			}
		}
	}()

	// 阻塞主 goroutine，使程序保持运行状态
	select {}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()

	}

}
func checkJwt(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		fmt.Println("错误")
		c.JSON(http.StatusUnauthorized, tools.ECode{Message: "未提供身份验证令牌"})
		c.Abort()
		return
	}
	parts := strings.Split(tokenStr, "Bearer")[1]
	claims, err := model.CheckJwt(parts)
	if err != nil {
		c.JSON(http.StatusUnauthorized, tools.ECode{Message: "身份验证失败！"})
		c.Abort()
		return
	}
	c.Set("user", claims)
	c.Next()
}
