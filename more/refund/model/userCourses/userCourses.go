package userCourses

import (
	"github.com/jinzhu/gorm"
	"refund/model"
)
type MUserCourses struct {
	Id           int64  `gorm:"primary_key" json:"id""`
	UserId       int64  `json:"user_id"`
	Sku          string `json:"sku"`
	OutTradeNo   string `json:"out_trade_no"`
	StartTime    string `json:"start_time"`
	EndTimes     string `json:"end_time"`
	Status       int8   `json:"status"`
	CreatedAt    string `json:"created_at"`
	DeletedAt    string `json:"deleted_at"`
	UpdatedAt    string `json:"updated_at"`
}
func (m *MUserCourses) GetDb() *gorm.DB {
	return model.InitDb()
}
