package models

import "time"

//Model
/**
@Author :  liushuiyuan
@Date 2026/4/22 21:34
*/
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
