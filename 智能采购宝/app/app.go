package app

import (
	"bysj1/app/model"
	"bysj1/app/router"
	"bysj1/app/tools"
)

func Start() {
	tools.LoadConfig()

	model.NewMysql()
	model.NewRdb()
	defer func() {
		model.Close()
	}()
	tools.NewLogger()
	router.New()
}
