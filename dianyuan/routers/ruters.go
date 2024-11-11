package routers

import (
	"dianyuan/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	//渲染模板
	r := gin.Default()
	// 使用中间件处理跨域请求
	//
	r.Use(cors.Default())
	V1Group := r.Group("eis")
	{
		//添加操作API
		V1Group.POST("/index", controller.CreateEis)

		//查询API(all)
		V1Group.GET("/index", controller.GetEislist)

		//查询API（only）
		V1Group.GET("/index/:eis_id", controller.GetEis)

		//修改操作API
		V1Group.PUT("/index/:eis_id", controller.UpdateEis)

		//删除操作API
		V1Group.DELETE("/index/:eis_id", controller.DeleteEis)
	}
	V1Group = r.Group("vct")
	{
		//添加操作API
		V1Group.POST("/index", controller.CreateVct)

		//查询API(all)
		V1Group.GET("/index", controller.GetVctlist)

		//查询API（only）
		V1Group.GET("/index/:time_id", controller.GetVct)

		//修改操作API
		V1Group.PUT("/index/:time_id", controller.UpdateVct)

		//删除操作API
		V1Group.DELETE("/index/:time_id", controller.DeleteVct)
	}
	return r
}
