package router

import (
	"github.com/astaxie/beego"
	"github.com/tengrommel/go_jobs/ch14/SecKill/SecProxy/controller"
)
func init()  {
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}