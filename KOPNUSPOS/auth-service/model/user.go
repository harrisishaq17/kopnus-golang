package model

import "time"

// request
type (
	LoginUserRequest struct {
		UserID   string `json:"user_id" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	GetDataUserRequest struct {
		UserID string `json:"user_id" validate:"required"`
	}

	ChangePasswordRequest struct {
		NewPassword        string `json:"new_password" validate:"required"`
		ConfirmNewPassword string `json:"confirm_new_password" validate:"required,eqfield=NewPassword"`
		OTP                string `json:"otp" validate:"required,len=5"`
		U                  string `json:"u" validate:"required"`
		UIH                string `json:"uih" validate:"required"`
	}

	UserPassReset struct {
		UserID      string
		KodeOTP     string
		ExpiredDate time.Time
		Status      string
		CreatedDate time.Time
	}

	SendOTPRequest struct {
		UserID      string `json:"user_id" validate:"required"`
		CallbackURL string `json:"callback_url" validate:"required"`
	}
)

// response
type (
	DataUserResponse struct {
		ID            string     `json:"id"`
		Name          string     `json:"name"`
		Email         string     `json:"email"`
		LastLoginDate *time.Time `json:"lastLoginDate"`
		Session       string     `json:"session"`
	}

	LoginUserResponse struct {
		Token               *Token `json:"token"`
		User                *User  `json:"user"`
		Cabang              string `json:"cabang"`
		KdWiliayah          string `json:"kd_wilayah"`
		NamaWilayah         string `json:"nama_wilayah"`
		JenisKantor         string `json:"jenis_kantor"`
		GroupName           string `json:"group_name"`
		FlagCabangPrioritas string `json:"flag_cabang_prioritas"`
		Autorisasi          string `json:"autorisasi"`
	}
)

// Struct Data
type (
	Token struct {
		AccessToken string `json:"access_token"`
		Type        string `json:"token_type"`
		ExpiresIn   string `json:"expires_in"`
	}

	User struct {
		UserID       string `json:"user_id"`
		KdCabang     string `json:"kd_cabang"`
		Nama         string `json:"nama"`
		StatusKantor string `json:"status_kantor"`
		FlagBudget   string `json:"flag_budget"`
	}
)
