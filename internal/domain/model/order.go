package model

type Order struct {
	Id       int       `gorm:"primaryKey" json:"id"`
	UserId   int       `gorm:"not null" json:"userId,omitempty"`
	Quantity int       `gorm:"not null" json:"quantity,omitempty"`
	User     User      `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Products []Product `gorm:"many2many:order_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"products,omitempty"`
}
