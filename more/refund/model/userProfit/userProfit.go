package userProfit

import (
	"github.com/jinzhu/gorm"
	"refund/model"
)

type CUserProfit struct {
	Id               int64   `gorm:"primary_key" json:"id""`
	OrderAmount      float32 `json:"order_amount"`
	TransferNo       string  `json:"transfer_no"`
	TransferWay      int8    `json:"transfer_way"`
	PayTime          string  `json:"pay_time"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	UserId           int64   `json:"user_id"`
	ProfitProportion int8    `json:"profit_proportion"`
	CostProportion   int8    `json:"cost_proportion"`
	Sku              string  `json:"sku"`
	ResultProfit     float32 `json:"result_profit"`
	Status           int8    `json:"status"`
	IsResult         int8    `json:"is_result"`
	CreatorUsersId   int64   `json:"creator_users_id"`
	DeletedAt        string  `json:"deleted_at"`
}
func (m *CUserProfit) GetDb() *gorm.DB {
	return model.InitDb()
}

