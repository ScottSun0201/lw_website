package client

import (
	"errors"
	"html"
	"regexp"
	"strings"
	"unicode/utf8"

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
	if clientUser.Password != "" {
		clientUser.Password = utils.BcryptHash(clientUser.Password)
	}
	err = global.GVA_DB.First(&client.ClientUser{}, "id = ?", clientUser.ID).Updates(&clientUser).Error
	return err
}

func (clientUserService *ClientUserService) SetClientUser(clientUser clientReq.SetClientUser) (err error) {
	db := global.GVA_DB.First(&client.ClientUser{}, "id = ?", clientUser.ID)

	if clientUser.Nickname != "" {
		db = db.Update("nickname", sanitizeText(clientUser.Nickname))
	}

	if clientUser.Avatar != "" {
		db = db.Update("avatar", clientUser.Avatar)
	}

	if clientUser.Age != nil {
		db = db.Update("age", clientUser.Age)
	}

	if clientUser.FirstName != "" {
		db = db.Update("first_name", sanitizeText(clientUser.FirstName))
	}

	if clientUser.LastName != "" {
		db = db.Update("last_name", sanitizeText(clientUser.LastName))
	}
	if clientUser.Email != "" {
		db = db.Update("email", strings.TrimSpace(clientUser.Email))
	}

	if clientUser.About != "" {
		db = db.Update("about", sanitizeText(clientUser.About))
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
	if info.PageSize <= 0 {
		info.PageSize = 10
	}
	if info.PageSize > 100 {
		info.PageSize = 100
	}
	if info.Page <= 0 {
		info.Page = 1
	}
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

// sanitizeText 转义HTML特殊字符，防止XSS
func sanitizeText(s string) string {
	return html.EscapeString(strings.TrimSpace(s))
}

// validateUsername 校验用户名格式
var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,30}$`)

// 用户注册
func (clientUserService *ClientUserService) Register(req clientReq.ClientUser) (err error, clientUser *client.ClientUser) {
	// 校验用户名格式
	if !usernameRegex.MatchString(req.Username) {
		return errors.New("用户名只能包含字母、数字和下划线，长度3-30"), nil
	}
	// 校验密码强度
	if utf8.RuneCountInString(req.Password) < 6 || utf8.RuneCountInString(req.Password) > 30 {
		return errors.New("密码长度必须在6-30个字符之间"), nil
	}

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
	if err == nil {
		clientUser.Password = "" // 不返回密码哈希
	}
	return err, clientUser
}

// 用户登录
func (clientUserService *ClientUserService) Login(req clientReq.ClientUser) (err error, clientUser *client.ClientUser) {
	err = global.GVA_DB.Where("username = ?", req.Username).First(&clientUser).Error
	if err != nil {
		return errors.New("用户名或密码错误"), nil
	}
	if !utils.BcryptCheck(req.Password, clientUser.Password) {
		return errors.New("用户名或密码错误"), nil
	}
	clientUser.Password = "" // 不返回密码哈希
	return nil, clientUser
}

// 获取自身信息
func (clientUserService *ClientUserService) GetUserInfo(id uint) (err error, clientUser *client.ClientUser) {
	err = global.GVA_DB.Omit("password").Where("id = ?", id).First(&clientUser).Error
	return
}
