package main

//导入postpre的数据库驱动
import (
	"dianyuan/dao"
	"dianyuan/models"
	"dianyuan/routers"
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	//声明一个全局变量来调用数据库
	err := dao.InitPostgre()
	if err != nil {
		panic(err)
	}
	//关闭数据库连接
	defer dao.DB.Close()
	//模型绑定
	models.Automodel()
	//路由组
	r := routers.Router()
	//r.Use(routers.CorsMiddleware())
	//r.Use(cors.Default())
	r.Run(":9090")
}
