package entity

type User struct {
	UserID              string `gorm:"column:user_id;primaryKey"`
	KdCabang            string `gorm:"column:kd_cabang"`
	MacAddress          string `gorm:"column:mac_address"`
	Password            string `gorm:"column:password"`
	Nama                string `gorm:"column:nama"`
	NIP                 string `gorm:"column:nip"`
	KdAutorisasi        string `gorm:"column:kd_autorisasi"`
	KdBtsWewenang       string `gorm:"column:kd_bts_wewenang"`
	KdBagian            string `gorm:"column:kd_bagian"`
	KdDirektorat        string `gorm:"column:kd_direktorat"`
	KdDepartemen        string `gorm:"column:kd_departemen"`
	StartNota           string `gorm:"column:start_nota"`
	TrxAntarCabang      string `gorm:"column:trx_antar_cabang"`
	Status              string `gorm:"column:status"`
	Flag                string `gorm:"column:flag"`
	StatusKantor        string `gorm:"column:status_kantor"`
	FirstPassword       string `gorm:"column:first_password"`
	IpAddress           string `gorm:"column:ip_address"`
	KdJabatan           string `gorm:"column:kd_jabatan"`
	PenggantianPassword string `gorm:"column:penggantian_password"`
	Tgl                 string `gorm:"column:tgl"`
	MasaBerlaku         string `gorm:"column:masa_berlaku"`
	LastUpdate          string `gorm:"column:last_update"`
	UserUpdate          string `gorm:"column:user_update"`
	SubledgerKas        string `gorm:"column:subledger_kas"`
	FlagBudget          string `gorm:"column:flag_budget"`
	SingleJurnal        string `gorm:"column:single_jurnal"`
	GroupCode           string `gorm:"column:group_code"`
	Email               string `gorm:"column:email"`
	Photo               string `gorm:"column:photo"`
	GroupCodeCore       string `gorm:"column:group_code_core"`
	GroupCodeMIS        string `gorm:"column:group_code_mis"`
	GroupCodeSalesforce string `gorm:"column:group_code_salesforce"`
	GroupMitra          string `gorm:"column:group_mitra"`
	Counter             string `gorm:"column:counter"`
}

func (u *User) TableName() string {
	return "tbl_user"
}
