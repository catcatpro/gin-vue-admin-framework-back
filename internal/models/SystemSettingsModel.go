package models

import (
	"gin_vue_admin_framework/common"
	"gorm.io/gorm"
)

type SystemSettings struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	SetKey   string `json:"set_key"  gorm:"NOT NULL"`
	SetValue string `json:"set_value"  gorm:"NOT NULL"`
}

//func (SystemSettingsModel) TableName() string {
//	return "system_settings"
//}

// 保存设置
func (sys_settings *SystemSettings) UpdateSysSettings(data []SystemSettings) error {
	var res *gorm.DB
	for i := range data {
		res = common.COM_DB.Where("set_key = ?", data[i].SetKey).First(sys_settings).Update("set_value", data[i].SetValue)
		if res.Error != nil {
			break
		}
	}
	return res.Error
}

// 获取系统设置
func (sys_settings *SystemSettings) GetSysSettings() ([]SystemSettings, error) {
	var systemSettings []SystemSettings
	res := common.COM_DB.Find(&systemSettings)
	return systemSettings, res.Error
}
