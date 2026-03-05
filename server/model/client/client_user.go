// 自动生成模板ClientUser
package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
)

// 客户端用户 结构体  ClientUser
type ClientUser struct {
	global.GVA_MODEL
	Username  string    `json:"username" form:"username" gorm:"column:username;comment:用户名;" binding:"required"` //用户名
	Password  string    `json:"password" form:"password" gorm:"column:password;comment:密码;" binding:"required"`  //密码
	Nickname  string    `json:"nickname" form:"nickname" gorm:"column:nickname;comment:用户昵称;"`                   //用户昵称
	Avatar    string    `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;" binding:"required"`        //头像
	Age       *int      `json:"age" form:"age" gorm:"column:age;comment:年龄;"`                                    //年龄
	Gender    string    `json:"gender" form:"gender" gorm:"column:gender;comment:性别;"`                           //性别
	About     string    `json:"about" form:"about" gorm:"column:about;comment:关于;"`                              //关于
	FirstName string    `json:"firstName" form:"firstName" gorm:"column:first_name;comment:名;"`                  //名
	LastName  string    `json:"lastName" form:"lastName" gorm:"column:last_name;comment:姓;"`                     //姓
	Email     string    `json:"email" form:"email" gorm:"column:email;comment:邮箱;"`                              //邮箱
	UUID      uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:UUID;"`
}

// TableName 客户端用户 ClientUser自定义表名 client_user
func (ClientUser) TableName() string {
	return "client_user"
}
