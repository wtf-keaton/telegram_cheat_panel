package config

import "gorm.io/gorm"

type Keys struct {
	ID           int    `gorm:"column:ID"`
	Key          string `gorm:"column:Key"`
	HWID         string `gorm:"column:HWID"`
	Status       string `gorm:"column:Status"`
	Subscribe    int64  `gorm:"column:Date"`
	SubscribeEnd int64  `gorm:"column:DateEnd"`
	Cheat        int    `gorm:"column:Cheat"`
}

type Cheats struct {
	ID          int    `gorm:"column:ID"`
	Name        string `gorm:"column:Name"`
	Status      string `gorm:"column:Status"`
	DllName     string `gorm:"column:DllName"`
	ProcessName string `gorm:"column:ProcessName"`
	WindowClass string `gorm:"column:WindowClass"`
}

type Admin struct {
	ID   int    `gorm:"column:ID"`
	Name string `gorm:"column:Name"`
}

type Hwids struct {
	ID   int    `gorm:"column:ID"`
	Name string `gorm:"column:Name"`
}

type ConfigurationData struct {
	AccessRights []int64 `json:"Access Rights Telegram IDs"`
}

var ApplicationConfig = ConfigurationData{
	AccessRights: []int64{1337, 1488},
}

type Models struct {
	DB    *gorm.DB
	Key   Keys
	Cheat Cheats
	Adm   Admin
	Hwid  Hwids
}
