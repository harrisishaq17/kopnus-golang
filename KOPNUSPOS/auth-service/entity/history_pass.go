package entity

type HistoryPass struct {
	UUID        string `gorm:"column:uuid"`
	UserID      string `gorm:"column:user_id"`
	PasswordOld string `gorm:"column:password_old"`
	PasswordNew string `gorm:"column:password_new"`
	SubmitDate  string `gorm:"column:submit_date"`
	ExpiredDate string `gorm:"column:expired_date"`
	KodeOTP     string `gorm:"column:kode_otp"`
	Status      string `gorm:"column:status"`
}

func (tbl *HistoryPass) TableName() string {
	return "tbl_history_change_pass"
}
