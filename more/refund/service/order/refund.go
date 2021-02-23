package order

import (
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"reflect"
	"refund/model/order"
	"refund/model/orderRefund"
	"refund/model/userCourses"
	"refund/model/userProfit"
	"refund/model/userSourses"
	"refund/model/usersVip"
	"time"
)

type Queue struct {
	Id         int64
	UserId     int64
	Amount     float32
	Sku        string
	TransferNo string
}

func ExcelData() map[int]map[int]map[int]string {
	dir, _ := os.Getwd()
	var (
		excel_file_path string                         = dir + "/service/order/2.xlsx"
		file_result     map[int]map[int]map[int]string = make(map[int]map[int]map[int]string)
		sheet_result    map[int]map[int]string         = make(map[int]map[int]string)
	)
	//打开一个excel文件资源
	f, err := xlsx.OpenFile(excel_file_path)
	if err != nil {
		log.Println(err.Error())
	}
	//循环文件中所有工作表
	for sheet_key, sheet := range f.Sheets {
		//循环对应工作表中行数
		for key, row := range sheet.Rows {
			row_result := make(map[int]string)
			//循环工作表行数的每一列
			for k, cell := range row.Cells {
				row_result[k] = cell.Value
			}
			//如果为空不添加对应值到 数组
			if !Empty(row_result) {
				sheet_result[key] = row_result
			}
		}
		//如果为空不添加对应值到 数组
		if !Empty(sheet_result) {
			file_result[sheet_key] = sheet_result
		}
	}
	return file_result
}
func Empty(params interface{}) bool {
	//初始化变量
	var (
		flag          bool = true
		default_value reflect.Value
	)
	r := reflect.ValueOf(params)
	//获取对应类型默认值
	default_value = reflect.Zero(r.Type())
	//由于params 接口类型 所以default_value也要获取对应接口类型的值 如果获取不为接口类型 一直为返回false
	if !reflect.DeepEqual(r.Interface(), default_value.Interface()) {
		flag = false
	}
	return flag
}

// 执行退款操作
func RefundOrder() bool {
	//输出表格的结果
	for _, sheet := range ExcelData() { // 获取表格结果
		for k, _ := range sheet {
			m := order.MOrders{}
			m.TransferNo = sheet[k][0]
			if m.TransferNo == "" { // 筛选掉表格为空的数据
				continue
			}
			db := m.GetDb()
			db.Select([]string{"id, user_id", "amount", "sku"}).Where("transfer_no=?", m.TransferNo).First(&m)
			if m.Id == 0 {
				log.Printf("暂无该笔订单, %v", m.TransferNo)
				continue
			}
			db.Close()

			log.Printf("--------------------------------------------------------")
			log.Printf("开始对该订单进行操作%v\n", m.TransferNo)
			// 1 添加退款表
			_, err := AddOrderRefund(m.TransferNo, m.Amount)
			if err != nil {
				log.Printf("添加退款订单表失败, %v", m.TransferNo)
				continue
			}
			bool := in(m.Sku, []string{"vip5", "monthly_vip"}) // 判断是什么类型的退款
			if bool {                                          // vip退款
				VipRefundOrder(m.TransferNo, m.UserId, m.Sku)
			} else { // 线上课程退款
				OnlineRefund(m.TransferNo)
			}
		}
	}
	return true
}

//vip 退款
func VipRefundOrder(TransferNo string, UserId int64, Sku string) bool {

	log.Printf("开始退款vip的订单 %v", TransferNo)

	var month int64 = 2626560                      // 月份
	var year int64 = 31536000                      // 年份
	t1 := time.Now().Format("2006-01-02 15:04:05") // 当前格式化时间
	// 1 订单表
	m := order.MOrders{}
	db1 := m.GetDb()
	err1 := db1.Model(&m).Where("transfer_no=?", TransferNo).Update(map[string]interface{}{"deleted_at": t1, "is_refund": 1}).Error
	if err1 != nil {
		log.Printf("订单表退订失败 TransferNo=%v", TransferNo)
		return false
	}
	defer db1.Close()

	// 2 vip用户表
	u := usersVip.MUsersVip{}
	db2 := u.GetDb()
	err2 := db2.Select([]string{"end_time"}).Where("user_id=?", UserId).Where("status=?", 1).First(&u).Error

	if u.EndTime == "" || err2 != nil {
		log.Printf("VipUsers无数据 UserId=%v ; EndTime=%v", u.UserId, u.EndTime)
		return false
	}
	defer db2.Close()

	loc, _ := time.LoadLocation("Local") // 获取时区
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", u.EndTime, loc)
	endTimes := theTime.Unix() // 转化为时间戳
	var endTimeUpdate int64
	switch Sku {
	case "vip5":
		endTimeUpdate = endTimes - year // 用户当前vip减去一年
	case "monthly_vip":
		endTimeUpdate = endTimes - month // 用户当前vip减去一个月
	default:
		endTimeUpdate = 0
	}
	// 3 修改用户vip资源表
	db3 := u.GetDb()
	err3 := db3.Model(&u).Where("user_id=?", UserId).Update(map[string]interface{}{
		"end_time": time.Unix(endTimeUpdate, 0).Format("2006-01-02 15:04:05"),
	}).Error
	if err3 != nil {
		log.Printf("VipUsers表退订失败 TransferNo=%v", TransferNo)
	}
	defer db3.Close()

	// 4 修改退款表
	a := orderRefund.MOrderRefund{}
	db4 := a.GetDb()
	err4 := db4.Model(&a).Where("transfer_no=?", TransferNo).Update(map[string]interface{}{"status": 3}).Error
	if err4 != nil {
		log.Printf("退款表修改失败 TransferNo=%v", TransferNo)
		return false
	}
	defer db4.Close()

	// 5 用户资源表
	s := userSourses.MUserSourses{}
	db5 := s.GetDb()
	err5 := db5.Model(&s).Where("out_trade_no=?", TransferNo).Update(map[string]interface{}{"status": -1, "deleted_at": t1}).Error
	if err5 != nil {
		log.Printf("用户所拥有的资源数据表退订失败 TransferNo=%v", TransferNo)
		return false
	}
	defer db5.Close()

	log.Printf("vip退款成功 TransferNo=%v", TransferNo)
	return true
}

// 判断字符串是否存在切片中
func in(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}

//退款综合方法
func OnlineRefund(TransferNo string) bool {
	bool1 := OnlineZiYuan(TransferNo) //资源表，退款表操作
	if bool1 == false {
		return false
	}
	UserProfitRefund(TransferNo) // 导师中心收益退款
	return true
}

// 导师中心退收益
func UserProfitRefund(TransferNo string) bool {

	log.Printf("退款导师中心开始;TransferNo:%v", TransferNo)
	t1 := time.Now().Format("2006-01-02 15:04:05")

	// 1 查询是否有收益
	u := userProfit.CUserProfit{}
	db := u.GetDb()
	err := db.Select([]string{"id"}).Where("transfer_no=?", TransferNo).First(&u).Error

	if u.Id == 0 || err != nil {
		log.Printf("没有收益，不需要处理导师中心; %v ;id=%v", u.TransferNo, u.Id)
		return false
	}

	// 2 有收益，退收益
	err2 := db.Model(&u).Where("transfer_no=?", TransferNo).Where("is_result=?", 1).Update(map[string]interface{}{"status": -1, "deleted_at": t1}).Error

	if err2 != nil {
		log.Printf("退收益失败; %v", u.TransferNo)
		return false
	}

	db.Close()
	log.Printf("退款导师中心成功;TransferNo:%v", u.TransferNo)
	return true
}

//墨尔线上课程退款
func OnlineZiYuan(TransferNo string) bool {

	log.Printf("退款线上课程订单开始;%v", TransferNo)
	t1 := time.Now().Format("2006-01-02 15:04:05")

	// 1 订单表
	m := order.MOrders{}
	db1 := m.GetDb()
	err1 := db1.Model(&m).Where("transfer_no=?", TransferNo).Update(map[string]interface{}{"deleted_at": t1, "is_refund": 1}).Error
	if err1 != nil {
		log.Printf("订单表修改失败;%v", TransferNo)
		return false
	}
	db1.Close()

	// 2 用户资源表
	c := userCourses.MUserCourses{}
	db2 := c.GetDb()
	err2 := db2.Model(&c).Where("out_trade_no=?", TransferNo).Update(map[string]interface{}{"status": -1, "deleted_at": t1}).Error
	if err2 != nil {
		log.Printf("用户资源表修改失败;%v", TransferNo)
		return false
	}
	db2.Close()

	// 3 用户资源章节表
	s := userSourses.MUserSourses{}
	db3 := s.GetDb()
	err3 := db3.Model(&s).Where("out_trade_no=?", TransferNo).Update(map[string]interface{}{"status": -1, "deleted_at": t1}).Error
	if err3 != nil {
		log.Printf("用户资源章节表修改失败;%v", TransferNo)
		return false
	}
	db3.Close()

	// 4 修改退款表
	a := orderRefund.MOrderRefund{}
	db4 := a.GetDb()
	err4 := db4.Model(&a).Where("transfer_no=?", TransferNo).Update(map[string]interface{}{"status": 3}).Error
	if err4 != nil {
		log.Printf("退款表修改失败;%v", TransferNo)
		return false
	}
	db4.Close()

	log.Printf("退款线上课程订单成功;%v", TransferNo)
	return true
}

// 添加退款表
func AddOrderRefund(TransferNo string, amount float32) (int64, error) {
	// 添加退款表
	t1 := time.Now().Format("2006-01-02 15:04:05")
	r := orderRefund.MOrderRefund{
		TransferNo:    TransferNo,
		CreatedAt:     t1,
		UpdatedAt:     t1,
		RefundRemarks: "后台脚本批量退款,操作人为技术人员",
		RefundPrice:   amount,
		RefundReason:  1,
		RefundType:    1,
	}
	db := r.GetDb()
	err := db.Create(&r).Error
	defer db.Close()
	return r.Id, err
}
