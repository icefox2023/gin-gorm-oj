package test

import (
	"fmt"
	"testing"

	"getcharzp.cn/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGormTest(t *testing.T) {
	dsn := "root:1234@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("User ==> %v \n", v)
	}

	var cnt int64
	err = models.DB.Where("mail = ?", "2826365325@qq.com").Model(new(models.UserBasic)).Count(&cnt).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Count ==> %d \n", cnt)
}
