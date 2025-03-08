package models

type JSONReponseMany struct {
	Success bool                `json:"success"`
	Code    int                 `json:"code"`
	Limit   int                 `json:"limit,omitempty"`
	Offset  int                 `json:"offset,omitempty"`
	Total   int                 `json:"total,omitempty"`
	Errors  map[string][]string `json:"errors,omitempty"`
	Data    any                 `json:"data"`
}

type JSONReponseOne struct {
	Success bool `json:"success"`
	Code    int  `json:"code"`
	Data    any  `json:"data"`
}

type UserLoginOutput struct {
	User         User   `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
