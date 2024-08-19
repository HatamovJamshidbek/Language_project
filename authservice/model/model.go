package model

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string  `json:"expires_at"`
}

type ResetPasswordRequest struct {
	ResetToken string `json:"reset_token"`
	NewPassword string `json:"new_password"`
}
