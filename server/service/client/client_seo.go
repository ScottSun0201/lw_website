package client

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
)

type ClientSEOService struct {
}

// CreateClientSEO 创建客户端SEO记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientSEOService *ClientSEOService) CreateClientSEO(clientSEO *client.ClientSEO) (err error) {
	err = global.GVA_DB.Save(clientSEO).Error
	return err
}

// GetClientSEO 根据ID获取客户端SEO记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientSEOService *ClientSEOService) GetClientSEO() (clientSEO client.ClientSEO, err error) {
	err = global.GVA_DB.First(&clientSEO).Error
	return
}
