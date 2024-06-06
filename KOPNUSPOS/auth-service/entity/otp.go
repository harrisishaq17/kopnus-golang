package entity

type OTP struct {
	UUID        string `gorm:"column:uuid"`
	UserID      string `gorm:"column:user_id"`
	Email       string `gorm:"column:email"`
	CreatedDate string `gorm:"column:created_date"`
	SubmitDate  string `gorm:"column:submit_date"`
	ExpiredDate string `gorm:"column:expired_date"`
	KodeOTP     string `gorm:"column:kode_otp"`
	Status      string `gorm:"column:status"`
}

func (tbl *OTP) TableName() string {
	return "tbl_user_pass_reset"
}
