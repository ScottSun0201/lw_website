package client

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	clientReq "github.com/flipped-aurora/gin-vue-admin/server/model/client/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
)

type ClientUserService struct {
}

// CreateClientUser 创建客户端用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientUserService *ClientUserService) CreateClientUser(clientUser *client.ClientUser) (err error) {
	ferr := global.GVA_DB.First(&clientUser, "username = ?", clientUser.Username).Error
	if ferr == nil {
		return errors.New("Username already exists")
	}
	clientUser.UUID, _ = uuid.NewV4()
	clientUser.Password = utils.BcryptHash(clientUser.Password)
	err = global.GVA_DB.Create(&clientUser).Error
	return err
}

// DeleteClientUser 删除客户端用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientUserService *ClientUserService) DeleteClientUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&client.ClientUser{}, "id = ?", ID).Error
	return err
}

// DeleteClientUserByIds 批量删除客户端用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientUserService *ClientUserService) DeleteClientUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]client.ClientUser{}, "id in ?", IDs).Error
	return err
}

// UpdateClientUser 更新客户端用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientUserService *ClientUserService) UpdateClientUser(clientUser clientReq.SetClientUser) (err error) {
	clientUser.Password = utils.BcryptHash(clientUser.Password)
	err = global.GVA_DB.First(&client.ClientUser{}, "id = ?", clientUser.ID).Updates(&clientUser).Error
	return err
}

func (clientUserService *ClientUserService) SetClientUser(clientUser clientReq.SetClientUser) (err error) {
	db := global.GVA_DB.First(&client.ClientUser{}, "id = ?", clientUser.ID)

	if clientUser.Nickname != "" {
		db = db.Update("nickname", clientUser.Nickname)
	}

	if clientUser.Avatar != "" {
		db = db.Update("avatar", clientUser.Avatar)
	}

	if clientUser.Age != nil {
		db = db.Update("age", clientUser.Age)
	}

	if clientUser.FirstName != "" {
		db = db.Update("first_name", clientUser.FirstName)
	}

	if clientUser.LastName != "" {
		db = db.Update("last_name", clientUser.LastName)
	}
	if clientUser.Email != "" {
		db = db.Update("email", clientUser.Email)
	}

	if clientUser.About != "" {
		db = db.Update("about", clientUser.About)
	}

	if clientUser.Password != "" {
		db = db.Update("password", utils.BcryptHash(clientUser.Password))
	}

	err = db.Error
	return err
}

// GetClientUser 根据ID获取客户端用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientUserService *ClientUserService) GetClientUser(ID string) (clientUser client.ClientUser, err error) {
	err = global.GVA_DB.Omit("password").Where("id = ?", ID).First(&clientUser).Error
	return
}

// GetClientUserInfoList 分页获取客户端用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientUserService *ClientUserService) GetClientUserInfoList(info clientReq.ClientUserSearch) (list []client.ClientUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&client.ClientUser{})
	var clientUsers []client.ClientUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.StartAge != nil && info.EndAge != nil {
		db = db.Where("age BETWEEN ? AND ? ", info.StartAge, info.EndAge)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Omit("password").Find(&clientUsers).Error
	return clientUsers, total, err
}

// 用户注册
func (clientUserService *ClientUserService) Register(req clientReq.ClientUser) (err error, clientUser *client.ClientUser) {
	ferr := global.GVA_DB.First(&clientUser, "username = ?", req.Username).Error
	if ferr == nil {
		return errors.New("Username already exists"), nil
	}
	clientUser.Username = req.Username
	clientUser.UUID, _ = uuid.NewV4()
	clientUser.Password = utils.BcryptHash(req.Password)
	clientUser.Avatar = "https://qmplusimg.henrongyi.top/gva_header.jpg"
	clientUser.Nickname = "NewUser—" + clientUser.UUID.String()[:8]
	err = global.GVA_DB.Create(&clientUser).Error
	return err, clientUser
}

// 用户登录
func (clientUserService *ClientUserService) Login(req clientReq.ClientUser) (err error, clientUser *client.ClientUser) {
	err = global.GVA_DB.Where("username = ?", req.Username).First(&clientUser).Error
	if err != nil {
		return errors.New("User not found"), nil
	}
	if !utils.BcryptCheck(req.Password, clientUser.Password) {
		return errors.New("密码错误"), nil
	}
	return nil, clientUser
}

// 获取自身信息
func (clientUserService *ClientUserService) GetUserInfo(id uint) (err error, clientUser *client.ClientUser) {
	err = global.GVA_DB.Omit("password").Where("id = ?", id).First(&clientUser).Error
	return
}
