package model

type Morders struct {
	Model
	UserId      string `json:"user_id"`
	TransferNo      string `json:"transfer_no"`
	IsRefund      string `json:"is_refund"`
	UpdatedAt string `json:"updated_at"`
}

// 修改数据
func EditMorders () {

}
