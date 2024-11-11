package controller

import (
	"dianyuan/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateEis(c *gin.Context) {
	//从请求之中绑定数据
	var eis_001 models.Eis_001
	c.BindJSON(&eis_001)
	//将绑定的数据存入数据库and返回响应
	err := models.CreateAEis(&eis_001)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, eis_001)
	}
}

func GetEislist(c *gin.Context) {
	eislist, err := models.GetAEislist()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, eislist)
	}
}

func GetEis(c *gin.Context) {
	eis_id, ok := c.Params.Get("eis_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	eis, err := models.GetAEis(eis_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, eis)
	}
}

func UpdateEis(c *gin.Context) {
	//eis_id
	eis_id, ok := c.Params.Get("eis_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	eis, err := models.GetAEis(eis_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&eis)
	err = models.UpdateAEis(eis)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, eis)
	}
}

func DeleteEis(c *gin.Context) {
	eis_id, ok := c.Params.Get("eis_id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	err := models.DeleteAEis(eis_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"eis_id": "delete"})
	}
}
