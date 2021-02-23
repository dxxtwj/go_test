package orderRefund

import (
	"github.com/jinzhu/gorm"
	"refund/model"
)

type MOrderRefund struct {
	Id            int64   `gorm:"primary_key" json:"id""`
	TransferNo    string  `json:"transfer_no"`
	AdminId       int64   `json:"admin_id"`
	AdminName     string  `json:"admin_name"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	Status        int8    `json:"status"`
	RefundRemarks string  `json:"refund_remarks"`
	RefundPrice   float32 `json:"refund_price"`
	RefundReason  int8    `json:"refund_reason"`
	RefundType    int8    `json:"refund_type"`
}


func (m *MOrderRefund) GetDb() *gorm.DB {
	return model.InitDb()
}

