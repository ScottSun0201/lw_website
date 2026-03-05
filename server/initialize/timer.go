package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/task"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 订单自动取消（每分钟执行）
		_, err = global.GVA_Timer.AddTaskByFunc("AutoCancelExpiredOrders", "0 */1 * * * *", func() {
			orderService := service.ServiceGroupApp.ShopServiceGroup.ShopOrderService
			orderService.AutoCancelExpiredOrders()
		}, "定时取消超时未支付订单", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 订单自动确认收货（每小时执行）
		_, err = global.GVA_Timer.AddTaskByFunc("AutoConfirmReceived", "0 0 */1 * * *", func() {
			orderService := service.ServiceGroupApp.ShopServiceGroup.ShopOrderService
			orderService.AutoConfirmReceived()
		}, "定时自动确认收货", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 订单自动完成（每小时执行）
		_, err = global.GVA_Timer.AddTaskByFunc("AutoComplete", "0 0 */1 * * *", func() {
			orderService := service.ServiceGroupApp.ShopServiceGroup.ShopOrderService
			orderService.AutoComplete()
		}, "定时自动完成订单", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 优惠券自动过期（每小时执行）
		_, err = global.GVA_Timer.AddTaskByFunc("AutoExpireUserCoupons", "0 0 */1 * * *", func() {
			couponService := service.ServiceGroupApp.ShopServiceGroup.ShopCouponService
			couponService.AutoExpireUserCoupons()
		}, "定时标记过期优惠券", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 库存预警检查（每30分钟执行）
		_, err = global.GVA_Timer.AddTaskByFunc("CheckInventoryAlerts", "0 */30 * * * *", func() {
			alertService := service.ServiceGroupApp.ShopServiceGroup.ShopInventoryAlertService
			alertService.CheckInventoryAlerts()
		}, "定时检查库存预警", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 推荐商品刷新（每日凌晨3点）
		_, err = global.GVA_Timer.AddTaskByFunc("RefreshRecommendations", "0 0 3 * * *", func() {
			recService := service.ServiceGroupApp.ShopServiceGroup.ShopRecommendationService
			recService.RefreshHotProducts()
			recService.RefreshNewProducts()
		}, "定时刷新推荐商品", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}
	}()
}
