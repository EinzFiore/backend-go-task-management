package users

type UserFormatter struct {
	ID          int    `json:"id"`
	Fullname    string `json:"fullname"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Token       string `json:"token"`
}

func UserFormat(user User, token string) UserFormatter {
	resFormat := UserFormatter{
		ID:          user.Id,
		Fullname:    user.Fullname,
		Description: user.Description,
		Avatar:      user.Avatar,
		Email:       user.Email,
		Token:       token,
	}

	return resFormat
}
