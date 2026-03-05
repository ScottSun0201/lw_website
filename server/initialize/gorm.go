package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cms"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Mysql.Dbname
		return GormMysql()
	case "pgsql":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Pgsql.Dbname
		return GormPgSql()
	case "oracle":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Oracle.Dbname
		return GormOracle()
	case "mssql":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Mssql.Dbname
		return GormMssql()
	case "sqlite":
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Sqlite.Dbname
		return GormSqlite()
	default:
		global.GVA_ACTIVE_DBNAME = &global.GVA_CONFIG.Mysql.Dbname
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysIgnoreApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCodePackage{},
		system.SysExportTemplate{},
		system.Condition{},
		system.JoinTemplate{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
		client.ClientUser{},
		cms.CmsPlate{},
		cms.CmsBanner{},
		cms.CmsFriend{},
		cms.CmsTag{},
		cms.CmsArticle{},
		cms.CmsFriendInfo{},
		client.ClientSEO{},
		client.ClientAbout{},
		client.ClientAboutInfo{}, cms.CmsCommit{},

		shop.ShopCategory{},
		shop.ShopBrand{},
		shop.ShopSpu{},
		shop.ShopSku{},
		shop.ShopProductImage{},
		shop.ShopUserAddress{},
		shop.ShopCart{},
		shop.ShopInventory{},
		shop.ShopInventoryLog{},
		shop.ShopOrder{},
		shop.ShopOrderItem{},
		shop.ShopOrderLog{},
		shop.ShopPayment{},
		shop.ShopUserWallet{},
		shop.ShopWalletLog{},

		// Phase 2
		shop.ShopNotification{},
		shop.ShopAdminNotification{},
		shop.ShopReview{},
		shop.ShopReviewStat{},
		shop.ShopCoupon{},
		shop.ShopUserCoupon{},
		shop.ShopRefund{},
		shop.ShopRefundLog{},
		shop.ShopFavorite{},
		shop.ShopLogisticsTrace{},
		shop.ShopInventoryAlert{},
		shop.ShopInventorySetting{},
		shop.ShopRecommendation{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel()

	if err != nil {
		global.GVA_LOG.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
