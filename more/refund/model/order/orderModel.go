package order

import (
	"github.com/jinzhu/gorm"
	"refund/model"
)

type MOrders struct {
	Id         int64  `gorm:"primary_key" json:"id""`
	UserId     int64 `json:"user_id"`
	TransferNo string `json:"transfer_no"`
	IsRefund   int8 `json:"is_refund"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
	Amount     float32 `json:"amount"`
	Sku        string `json:"sku"`
}

func (m *MOrders) GetDb() *gorm.DB {
	return model.InitDb()
}

