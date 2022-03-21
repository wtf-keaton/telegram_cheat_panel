package client_api

import (
	"fmt"
	"os"
	"telegram_webpanel/internal/dbApi"
	"telegram_webpanel/pgk/config"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMultiplexer() *gin.Engine {
	var r *gin.Engine = gin.Default()

	r.POST("/auth", func(c *gin.Context) {

		if len(c.PostForm("key")) <= 5 || len(c.PostForm("hwid")) <= 5 {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Wrong data"})
			return
		}

		var findKeys config.Keys
		dbApi.Model.DB.Where("`Key` = ?", c.PostForm("key")).Find(&findKeys)
		dbApi.Model.DB.Select("HWID").Where("`Key` = ?", c.PostForm("key")).Find(&findKeys)
		dbApi.Model.DB.Select("DateEnd").Where("`Key` = ?", c.PostForm("key")).Find(&findKeys)
		dbApi.Model.DB.Select("Status").Where("`Key` = ?", c.PostForm("key")).Find(&findKeys)

		var findHwids config.Hwids
		dbApi.Model.DB.Select("Name").Find(&findHwids)

		if findHwids.Name == c.PostForm("hwid") {
			dbApi.BanKey(c.PostForm("key"))

			c.JSON(200, gin.H{"Status": "banned"})
			return
		}
		if len(findKeys.HWID) == 0 {
			dbApi.SetHWID(c.PostForm("key"), c.PostForm("hwid"))
		}
		dbApi.Model.DB.Select("HWID").Where("`Key` = ?", c.PostForm("key")).Find(&findKeys)

		if findKeys.Status == "waiting" {
			dbApi.SetEndTime(c.PostForm("key"), findKeys.Subscribe+time.Now().Unix())
			dbApi.SetActivate(c.PostForm("key"))

			c.JSON(200, gin.H{"Status": "Activated"})
			return
		}

		if findKeys.Status == "banned" {
			c.JSON(200, gin.H{"Status": "hwid banned"})
			return

		}
		if findKeys.Status == "ended" {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Subscribe Expired"})
			return
		}

		if findKeys.Key != c.PostForm("key") {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Key not found"})
			return
		}

		if findKeys.Key == c.PostForm("key") && findKeys.HWID == c.PostForm("hwid") && findKeys.Status != "banned" && findKeys.SubscribeEnd >= time.Now().Unix() {
			c.JSON(200, gin.H{"Status": "Authorized"})
			return
		}

		if findKeys.HWID != c.PostForm("hwid") {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Wrong HWID"})
			return
		}

		if findKeys.Status == "banned" {
			c.JSON(200, gin.H{"Status": "Error", "msg": "key banned"})
			return
		}

		if findKeys.SubscribeEnd <= time.Now().Unix() {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Subscribe Expired"})
			return
		}
	})

	r.POST("/dll", func(c *gin.Context) {
		key := c.PostForm("key")
		hwid := c.PostForm("hwid")

		if len(key) <= 5 || len(hwid) <= 5 {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Wrong data"})
			return
		}
		var findKeys config.Keys
		dbApi.Model.DB.Where("`Key` = ?", key).Find(&findKeys)
		dbApi.Model.DB.Select("HWID").Where("`Key` = ?", key).Find(&findKeys)
		dbApi.Model.DB.Select("DateEnd").Where("`Key` = ?", key).Find(&findKeys)
		dbApi.Model.DB.Select("Status").Where("`Key` = ?", key).Find(&findKeys)

		if findKeys.Key != key {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Key not found"})
			return
		}

		if findKeys.Key == key && findKeys.HWID == hwid && findKeys.Status != "banned" && findKeys.SubscribeEnd >= time.Now().Unix() {
			var keyInfo config.Keys
			dbApi.Model.DB.Select("Cheat").Where("`Key` = ?", key).Find(&keyInfo)

			cheatId := keyInfo.Cheat
			cheatName := dbApi.GetCheatDllById(cheatId)

			fmt.Println("Cheat name: ", cheatName)

			e, _ := os.Getwd()
			cheatPath := e + "/" + cheatName
			fmt.Println("Path: ", cheatPath)

			c.FileAttachment(cheatPath, "cheat")
			return
		}

	})
	r.POST("/driver", func(c *gin.Context) {
		key := c.PostForm("key")
		hwid := c.PostForm("hwid")

		if len(key) <= 5 || len(hwid) <= 5 {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Wrong data"})
			return
		}
		var findKeys config.Keys
		dbApi.Model.DB.Where("`Key` = ?", key).Find(&findKeys)
		dbApi.Model.DB.Select("HWID").Where("`Key` = ?", key).Find(&findKeys)
		dbApi.Model.DB.Select("DateEnd").Where("`Key` = ?", key).Find(&findKeys)
		dbApi.Model.DB.Select("Status").Where("`Key` = ?", key).Find(&findKeys)

		if findKeys.Key != key {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Key not found"})
			return
		}

		if findKeys.Key == key && findKeys.HWID == hwid && findKeys.Status != "banned" && findKeys.SubscribeEnd >= time.Now().Unix() {

			e, _ := os.Getwd()
			driverPath := e + "/" + "driver.sys"
			c.FileAttachment(driverPath, "driver")
			return
		}

	})
	r.POST("/process", func(c *gin.Context) {
		key := c.PostForm("key")
		hwid := c.PostForm("hwid")

		if len(key) <= 5 || len(hwid) <= 5 {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Wrong data"})
			return
		}

		var findKeys config.Keys
		dbApi.Model.DB.Where("`Key` = ?", key).Find(&findKeys)
		dbApi.Model.DB.Select("HWID").Where("`Key` = ?", key).Find(&findKeys)
		dbApi.Model.DB.Select("DateEnd").Where("`Key` = ?", key).Find(&findKeys)
		dbApi.Model.DB.Select("Status").Where("`Key` = ?", key).Find(&findKeys)

		if findKeys.Key != key {
			c.JSON(200, gin.H{"Status": "Error", "msg": "Key not found"})
			return
		}

		if findKeys.Key == key && findKeys.HWID == hwid && findKeys.Status != "banned" && findKeys.SubscribeEnd >= time.Now().Unix() {
			var keyInfo config.Keys
			dbApi.Model.DB.Select("Cheat").Where("`Key` = ?", key).Find(&keyInfo)

			cheatId := keyInfo.Cheat
			processName := dbApi.GetProcessNameById(cheatId)
			c.JSON(200, gin.H{"Process": processName})

			return
		}
	})
	r.POST("/ban", func(c *gin.Context) {
		hwid := c.PostForm("hwid")

		dbApi.BanHwid(hwid)

		c.JSON(200, gin.H{"Status": "Success"})

		return
	})
	return r
}
