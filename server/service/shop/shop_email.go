package shop

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"runtime"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/client"
	emailUtils "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
	"go.uber.org/zap"
)

type ShopEmailService struct{}

// getTemplateDir returns absolute path to email_templates directory
func getTemplateDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "email_templates")
}

// getUserEmail 根据用户ID获取邮箱
func getUserEmail(userID uint) string {
	var user client.ClientUser
	if err := global.GVA_DB.Select("email").Where("id = ?", userID).First(&user).Error; err != nil {
		return ""
	}
	return user.Email
}

// renderTemplate 渲染邮件模板
func renderTemplate(tmplName string, data interface{}) (string, error) {
	tmplPath := filepath.Join(getTemplateDir(), tmplName)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return "", fmt.Errorf("解析邮件模板失败: %v", err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("渲染邮件模板失败: %v", err)
	}
	return buf.String(), nil
}

// sendEmail 发送邮件（内部方法，失败只记日志）
func sendEmail(to, subject, body string) {
	if to == "" {
		return
	}
	if err := emailUtils.Email(to, subject, body); err != nil {
		global.GVA_LOG.Error("发送邮件失败", zap.String("to", to), zap.String("subject", subject), zap.Error(err))
	}
}

// SendOrderCreatedEmail 发送订单创建邮件
func (s *ShopEmailService) SendOrderCreatedEmail(userID uint, orderNo string, payAmount float64) {
	email := getUserEmail(userID)
	if email == "" {
		return
	}
	body, err := renderTemplate("order_created.html", map[string]interface{}{
		"OrderNo":   orderNo,
		"PayAmount": fmt.Sprintf("%.2f", payAmount),
	})
	if err != nil {
		global.GVA_LOG.Error("渲染订单创建邮件失败", zap.Error(err))
		return
	}
	sendEmail(email, "您的订单已创建 - "+orderNo, body)
}

// SendOrderPaidEmail 发送订单支付成功邮件
func (s *ShopEmailService) SendOrderPaidEmail(userID uint, orderNo string, payAmount float64) {
	email := getUserEmail(userID)
	if email == "" {
		return
	}
	body, err := renderTemplate("order_paid.html", map[string]interface{}{
		"OrderNo":   orderNo,
		"PayAmount": fmt.Sprintf("%.2f", payAmount),
	})
	if err != nil {
		global.GVA_LOG.Error("渲染支付成功邮件失败", zap.Error(err))
		return
	}
	sendEmail(email, "支付成功 - "+orderNo, body)
}

// SendOrderShippedEmail 发送订单发货邮件
func (s *ShopEmailService) SendOrderShippedEmail(userID uint, orderNo, shipCompany, shipNo string) {
	email := getUserEmail(userID)
	if email == "" {
		return
	}
	body, err := renderTemplate("order_shipped.html", map[string]interface{}{
		"OrderNo":     orderNo,
		"ShipCompany": shipCompany,
		"ShipNo":      shipNo,
	})
	if err != nil {
		global.GVA_LOG.Error("渲染发货邮件失败", zap.Error(err))
		return
	}
	sendEmail(email, "您的订单已发货 - "+orderNo, body)
}

// SendOrderCompletedEmail 发送订单完成邮件
func (s *ShopEmailService) SendOrderCompletedEmail(userID uint, orderNo string) {
	email := getUserEmail(userID)
	if email == "" {
		return
	}
	body, err := renderTemplate("order_completed.html", map[string]interface{}{
		"OrderNo": orderNo,
	})
	if err != nil {
		global.GVA_LOG.Error("渲染订单完成邮件失败", zap.Error(err))
		return
	}
	sendEmail(email, "订单已完成 - "+orderNo, body)
}

// SendRefundApprovedEmail 发送退款通过邮件
func (s *ShopEmailService) SendRefundApprovedEmail(userID uint, orderNo string, amount float64) {
	email := getUserEmail(userID)
	if email == "" {
		return
	}
	body, err := renderTemplate("refund_approved.html", map[string]interface{}{
		"OrderNo": orderNo,
		"Amount":  fmt.Sprintf("%.2f", amount),
	})
	if err != nil {
		global.GVA_LOG.Error("渲染退款通过邮件失败", zap.Error(err))
		return
	}
	sendEmail(email, "退款已通过 - "+orderNo, body)
}

// SendRefundRejectedEmail 发送退款拒绝邮件
func (s *ShopEmailService) SendRefundRejectedEmail(userID uint, orderNo, reason string) {
	email := getUserEmail(userID)
	if email == "" {
		return
	}
	body, err := renderTemplate("refund_rejected.html", map[string]interface{}{
		"OrderNo": orderNo,
		"Reason":  reason,
	})
	if err != nil {
		global.GVA_LOG.Error("渲染退款拒绝邮件失败", zap.Error(err))
		return
	}
	sendEmail(email, "退款申请被拒绝 - "+orderNo, body)
}
