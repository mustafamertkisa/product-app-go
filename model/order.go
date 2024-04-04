package model

type Order struct {
	Id        uint    `gorm:"primaryKey" json:"id"`
	UserId    uint    `gorm:"not null" json:"userId,omitempty"`
	ProductId uint    `gorm:"not null" json:"productId,omitempty"`
	Quantity  int     `gorm:"not null" json:"quantity,omitempty"`
	User      User    `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Product   Product `gorm:"foreignKey:ProductId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
