package system

import (
	"gin_vue_admin_framework/common"
	"gorm.io/gorm"
)

type SystemSettingsModel struct {
	SetKey   string `json:"set_key"  gorm:"NOT NULL"`
	SetValue int    `json:"set_value"  gorm:"NOT NULL"`
}

// 保存设置
func (sys_settings *SystemSettingsModel) SaveSysSettings(data []SystemSettingsModel) error {
	var res *gorm.DB
	for i := range data {
		res = common.COM_DB.Where("set_key = ?", data[i].SetKey).Save(&data[i])
		if res.Error != nil {
			break
		}
	}
	return res.Error
}
