package shop

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"gorm.io/gorm"
)

type ShopAddressService struct{}

// CreateAddress 创建收货地址
func (s *ShopAddressService) CreateAddress(address *shop.ShopUserAddress) error {
	// 检查用户地址数量是否超过20
	var count int64
	if err := global.GVA_DB.Model(&shop.ShopUserAddress{}).Where("user_id = ?", address.UserID).Count(&count).Error; err != nil {
		return err
	}
	if count >= 20 {
		return errors.New("收货地址数量已达上限(20个)")
	}

	// 如果设置为默认地址，需要在事务中清除其他默认地址
	if address.IsDefault != nil && *address.IsDefault == 1 {
		return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			// 清除该用户其他默认地址
			if err := tx.Model(&shop.ShopUserAddress{}).
				Where("user_id = ? AND is_default = ?", address.UserID, 1).
				Update("is_default", 0).Error; err != nil {
				return err
			}
			return tx.Create(address).Error
		})
	}

	return global.GVA_DB.Create(address).Error
}

// DeleteAddress 删除收货地址
func (s *ShopAddressService) DeleteAddress(ID uint, userID uint) error {
	// 验证地址归属
	var address shop.ShopUserAddress
	if err := global.GVA_DB.Where("id = ? AND user_id = ?", ID, userID).First(&address).Error; err != nil {
		return errors.New("地址不存在或无权操作")
	}
	return global.GVA_DB.Delete(&shop.ShopUserAddress{}, "id = ? AND user_id = ?", ID, userID).Error
}

// UpdateAddress 更新收货地址
func (s *ShopAddressService) UpdateAddress(address shop.ShopUserAddress) error {
	// 验证地址归属
	var existing shop.ShopUserAddress
	if err := global.GVA_DB.Where("id = ? AND user_id = ?", address.ID, address.UserID).First(&existing).Error; err != nil {
		return errors.New("地址不存在或无权操作")
	}

	// 如果设置为默认地址，需要在事务中清除其他默认地址
	if address.IsDefault != nil && *address.IsDefault == 1 {
		return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			// 清除该用户其他默认地址
			if err := tx.Model(&shop.ShopUserAddress{}).
				Where("user_id = ? AND is_default = ? AND id != ?", address.UserID, 1, address.ID).
				Update("is_default", 0).Error; err != nil {
				return err
			}
			return tx.Save(&address).Error
		})
	}

	return global.GVA_DB.Save(&address).Error
}

// GetAddress 根据ID获取收货地址
func (s *ShopAddressService) GetAddress(ID uint) (address shop.ShopUserAddress, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&address).Error
	return
}

// GetAddressList 获取用户所有收货地址列表
func (s *ShopAddressService) GetAddressList(userID uint) (list []shop.ShopUserAddress, err error) {
	err = global.GVA_DB.Where("user_id = ?", userID).
		Order("is_default DESC, created_at DESC").
		Find(&list).Error
	return
}

// SetDefaultAddress 设置默认地址
func (s *ShopAddressService) SetDefaultAddress(ID uint, userID uint) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 清除该用户所有默认地址
		if err := tx.Model(&shop.ShopUserAddress{}).
			Where("user_id = ? AND is_default = ?", userID, 1).
			Update("is_default", 0).Error; err != nil {
			return err
		}
		// 设置指定地址为默认
		result := tx.Model(&shop.ShopUserAddress{}).
			Where("id = ? AND user_id = ?", ID, userID).
			Update("is_default", 1)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("地址不存在或无权操作")
		}
		return nil
	})
}

// GetDefaultAddress 获取用户默认地址
func (s *ShopAddressService) GetDefaultAddress(userID uint) (address shop.ShopUserAddress, err error) {
	err = global.GVA_DB.Where("user_id = ? AND is_default = ?", userID, 1).First(&address).Error
	return
}
