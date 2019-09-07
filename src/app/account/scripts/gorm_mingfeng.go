package scripts

import (

	 userinfo"FakeBilibili/src/app/account/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func  CreateUserInfosTable(){
	db,err := gorm.Open("mysql","mingfeng:6963038@/mingfeng?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&userinfo.UserInfo{})

}


