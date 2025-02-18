package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserID        uint32 `json:"user_id"`
	OrderID       string `json:"order_id"`
	TransactionID string `gorm:"type:varchar(36);unique_index"`
	Amount        float32 `json:"amount"`
	PayAt         time.Time  `json:"pay_at"`
}

func (p PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB, ctx context.Context, payment *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
}

