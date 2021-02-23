package userSourses

import (
	"github.com/jinzhu/gorm"
	"refund/model"
)

type MUserSourses struct {
	Id           int64  `gorm:"primary_key" json:"id""`
	UserId       int64  `json:"user_id"`
	ResourceId   int64  `json:"resource_id"`
	Sku          string `json:"sku"`
	OutTradeNo   string `json:"out_trade_no"`
	StartTime    string `json:"start_time"`
	EndTimes     string `json:"end_time"`
	TicketQrCode string `json:"ticket_qr_code"`
	ResourcType  int8   `json:"resource_type"`
	LiveSourses  int64  `json:"live_sourses"`
	Status       int8   `json:"status"`
	CreatedAt    string `json:"created_at"`
	DeletedAt    string `json:"deleted_at"`
	UpdatedAt    string `json:"updated_at"`
}
func (m *MUserSourses) GetDb() *gorm.DB {
	return model.InitDb()
}

