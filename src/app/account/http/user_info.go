package http

import (
	//"github.com/go-chassis/go-chassis/core/registry"
	rf "github.com/go-chassis/go-chassis/server/restful"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserInfo struct {
	gorm.Model

	UserAccount	 	string
	Password		string
	UserName 		string
}

func (user *UserInfo) URLPatterns() []rf.Route{

	return []rf.Route{
		{http.MethodGet,"/login","UserLogin","",nil,nil,nil,nil,nil},
		{http.MethodGet,"/registry","UserRegistry","",nil,nil,nil,nil,nil},
	}
}
/*
func (user *UserInfo) UserLogin(ctx *rf.Context) {
	res := struct {
		res		string
	}{}

	userAccount := ctx.ReadQueryParameter("userAccount")
	Password	:= ctx.ReadQueryParameter("password")

	//check database   db
}

func (user *UserInfo) UserRegistry(ctx *rf.Context){
	res := struct{
		res 	string
	}{}

	userAccount := ctx.ReadQueryParameter("userAccount")
	password	:= ctx.ReadQueryParameter("password")
	userName 	:=ctx.ReadQueryParameter("userName")

	//database
	db,err := gorm.Open("mysql","mingfeng:6963038@/mingfeng?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	defer db.Close()



}

*/
