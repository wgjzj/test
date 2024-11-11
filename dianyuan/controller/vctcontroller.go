package controller

import (
	"dianyuan/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateVct(c *gin.Context) {
	//从请求之中绑定数据
	var vct_001 models.Vct_001
	c.BindJSON(&vct_001)
	//将绑定的数据存入数据库and返回响应
	fmt.Printf("%v", vct_001)
	err := models.CreateAVct(&vct_001)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, vct_001)
	}
}

func GetVctlist(c *gin.Context) {
	vctlist, err := models.GetAVctlist()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, vctlist)
	}
}

func GetVct(c *gin.Context) {
	time_id, ok := c.Params.Get("time_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	vctlist, err := models.GetAVct(time_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, vctlist)
	}
}

func UpdateVct(c *gin.Context) {
	//
	time_id, ok := c.Params.Get("time_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	vct_001, err := models.GetAVct(time_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&vct_001)
	err = models.UpdateAVct(vct_001)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, vct_001)
	}
}

func DeleteVct(c *gin.Context) {
	time_id, ok := c.Params.Get("time_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	err := models.DeleteAVct(time_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"time_id": "delete"})
	}
}
