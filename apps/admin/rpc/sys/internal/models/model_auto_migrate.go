package models

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

//InitModels
/**
@desc 数据库初始化
@param db *gorm.DB
@Author :  liushuiyuan
@Date 2026/4/22 21:34
*/
func InitModels(db *gorm.DB) {
	err := db.AutoMigrate(
		&UserModel{},
		&UserConfModel{},
	)
	if err != nil {
		logx.Errorf("数据库迁移失败：%s", err)
		return
	}
	logx.Infof("数据库迁移成功")
}
