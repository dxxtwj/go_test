package usersVip

import (
	"github.com/jinzhu/gorm"
	"refund/model"
)

type MUsersVip struct {
	Id        int64  `gorm:"primary_key" json:"id""`
	VipId     int64  `json:"vip_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	UserId    int64  `json:"user_id"`
	Status    int8   `json:"status"`
	StartTime string `json:"start_time"`
	EndTime  string `json:"end_time"`
}

func (m *MUsersVip) GetDb() *gorm.DB {
	return model.InitDb()
}
