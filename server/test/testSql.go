package main

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type HospitalVo struct {
	ID        uint `gorm:"primarykey"`
	//企业名称
	HospitalName string 		`json:"name" gorm:"comment:企业名称;"`
	//社会信用代码
	Code string 				`json:"code" gorm:"comment:社会信用代码;"`
	//注册地址
	Address string 				`json:"address" gorm:"comment:注册地址;"`
	//负责人姓名
	Username string				`json:"user_name"`
	//身份证号
	IdentityCard string 		`json:"identity_card" gorm:"comment:身份证号;"`//想写一个，注册时默认登录密码为身份证号码后六位的接口
	//电话号码
	Phone 		 string 		`json:"phone" gorm:"comment:电话号码;"`
	//关联区县字段
	DistrictName string			`json:"district_name"` //县区名
}

func main() {
	dsn := "root:56994902@tcp(127.0.0.1:3306)/db_fever_information?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var voDate []HospitalVo
	/*err := db.Table("hospital").
		Select("hospital.id", "hospital.hospital_name","code","address","sys_users.name","sys_users.identity_card","sys_users.phone","district.district_name").
		Joins("left join sys_users on sys_users.hospital_id = hospital.id").
		Joins("left join district on district.id = hospital.district_id").Debug().Scan(&voDate).Error*/

	err := db.Model(&system.Hospital{}).Debug().Select("hospitals.*,districts.district_name,sys_users.*").
		Joins("left join districts on districts.id = hospitals.district_id").
		Joins("left join sys_users on sys_users.id = hospitals.boos_id").
		Scan(&voDate).Error


	if err != nil {
		fmt.Printf("失败，%v",err)
	}
	fmt.Println(voDate)
}
