package main

import (
	"github.com/hectorjiang/workbench/model"
	"github.com/hectorjiang/workbench/route"
)

func main() {
	model.InitMysqlDB()
	route.InitRouter()
}
