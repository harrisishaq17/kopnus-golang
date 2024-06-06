package entity

type ApiSession struct {
	ID          string `gorm:"column:id"`
	UserID      string `gorm:"column:user_id"`
	Source      string `gorm:"column:source"`
	CreatedDate string `gorm:"column:created_date"`
	ExpiredDate string `gorm:"column:expired_date"`
	Token       string `gorm:"column:token"`
	LastUpdate  string `gorm:"column:last_update"`
}

func (u *ApiSession) TableName() string {
	return "tbl_api_session"
}
