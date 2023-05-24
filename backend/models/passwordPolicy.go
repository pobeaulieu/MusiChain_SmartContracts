package models

type PasswordPolicy struct {
	Id            uint  `gorm:"primaryKey"`
	MinLength     *int  `json:"minLength"`
	RequireNumber *bool `json:"requireNumber"`
	RequireLower  *bool `json:"requireLower"`
	RequireUpper  *bool `json:"requireUpper"`
	RequireSymbol *bool `json:"requireSymbol"`
	HistorySize   *int  `json:"historySize"`
}
