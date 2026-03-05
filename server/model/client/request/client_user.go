package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ClientUserSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Username string `json:"username" form:"username" `
	Nickname string `json:"nickname" form:"nickname" `
	StartAge *int   `json:"startAge" form:"startAge"`
	EndAge   *int   `json:"endAge" form:"endAge"`
	request.PageInfo
}

type ClientUser struct {
	global.GVA_MODEL
	Username   string `json:"username" form:"username" gorm:"column:username;comment:用户名;" binding:"required"` //用户名
	Password   string `json:"password" form:"password" gorm:"column:password;comment:密码;" binding:"required"`  //密码
	RePassword string `json:"rePassword" form:"rePassword" gorm:"-"`                                           // 确认密码
}

type SetClientUser struct {
	global.GVA_MODEL
	Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;"`    //密码
	Nickname  string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:用户昵称;"`  //用户昵称
	Avatar    string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;"`          //头像
	Age       *int   `json:"age" form:"age" gorm:"column:age;comment:年龄;"`                   //年龄
	About     string `json:"about" form:"about" gorm:"column:about;comment:关于;"`             //关于
	FirstName string `json:"firstName" form:"firstName" gorm:"column:first_name;comment:名;"` //名
	LastName  string `json:"lastName" form:"lastName" gorm:"column:last_name;comment:姓;"`    //姓
	Email     string `json:"email" form:"email" gorm:"column:email;comment:邮箱;"`             //邮箱
	Gender    string `json:"gender" form:"gender" gorm:"column:gender;comment:性别;"`          //性别
}
