// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePaymentUser = "payment_users"

// PaymentUser mapped from table <payment_users>
type PaymentUser struct {
	ID           string  `gorm:"column:id;type:text;primaryKey" json:"id"`
	TwinteUserID string  `gorm:"column:twinte_user_id;type:uuid;not null" json:"twinte_user_id"`
	DisplayName  *string `gorm:"column:display_name;type:text" json:"display_name"`
	Link         *string `gorm:"column:link;type:text" json:"link"`
}

// TableName PaymentUser's table name
func (*PaymentUser) TableName() string {
	return TableNamePaymentUser
}
