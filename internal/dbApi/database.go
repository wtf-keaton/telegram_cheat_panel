package dbApi

import (
	"fmt"
	"telegram_webpanel/internal/generator"
	"telegram_webpanel/pgk/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Model config.Models

func HandleDB() {

	Model.DB, _ = gorm.Open(mysql.Open("username:password@tcp(127.0.0.1:3306)/database_name"), &gorm.Config{})

	Model.DB.AutoMigrate(&Model.Key)   // База с ключами
	Model.DB.AutoMigrate(&Model.Cheat) // База с читами
	Model.DB.AutoMigrate(&Model.Adm)   // База с админами
	Model.DB.AutoMigrate(&Model.Hwid)  // База с забанеными хвидами
}

func CreateKey(date int64, cheatid int) string {

	Model.DB.Create(&config.Keys{Key: generator.Key(32), HWID: "", Subscribe: date * 86400, Status: "waiting", Cheat: cheatid})

	var keys []config.Keys
	Model.DB.Find(&keys)

	var res string
	for i := 0; i < len(keys); i++ {
		res = fmt.Sprintf("ID: %d, Key: %s", i, keys[i].Key)
	}
	return res
}

func CreateCheat(name string) {
	Model.DB.Create(&config.Cheats{Name: name, Status: "Undetected", DllName: name + ".dll"})
}

func SetHWID(key, hwid string) {
	Model.DB.Model(&Model.Key).Where("`Key` = ?", key).Update("HWID", hwid)
}

func ResetHWID(key string) {
	Model.DB.Model(&Model.Key).Where("`Key` = ?", key).Update("HWID", "")
}

func BanKey(key string) {
	Model.DB.Model(&Model.Key).Where("`Key` = ?", key).Update("Status", "banned")
}

func BanHwid(hwid string) {
	Model.DB.Create(&config.Hwids{Name: hwid})
}

func SetActivate(key string) {
	Model.DB.Debug().Model(&Model.Key).Where("`Key` = ?", key).Update("Status", "activated")
}

func SetEndTime(key string, date int64) {
	Model.DB.Debug().Model(&Model.Key).Where("`Key` = ?", key).Update("DateEnd", date)
}

func GetCheatInfo(name string) string {
	var cheatsInfo config.Cheats

	Model.DB.Select("ID").Where("`Name` = ?", name).Find(&cheatsInfo)
	Model.DB.Select("Name").Where("`Name` = ?", name).Find(&cheatsInfo)
	Model.DB.Select("Status").Where("`Name` = ?", name).Find(&cheatsInfo)

	return fmt.Sprintf("ID: %d Name: %s Status: %s", cheatsInfo.ID, cheatsInfo.Name, cheatsInfo.Status)
}

func GetCheatDllById(id int) string {
	var cheatsInfo config.Cheats
	Model.DB.Select("DllName").Where("`ID` = ?", id).Find(&cheatsInfo)

	return cheatsInfo.DllName
}

func GetProcessNameById(id int) string {
	var cheatsInfo config.Cheats
	Model.DB.Select("ProcessName").Where("`ID` = ?", id).Find(&cheatsInfo)

	return cheatsInfo.ProcessName
}
