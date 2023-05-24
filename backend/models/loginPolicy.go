package models

type LoginPolicy struct {
	Id                   uint `json:"id" gorm:"primaryKey"`
	MaxLoginAttemptCount uint `json:"maxLoginAttemptCount" gorm:"default:5"`
	LoginTimeInterval    uint `json:"loginTimeInterval" gorm:"default:5"`
	TwoFA                bool `json:"twofa" gorm:"default:false"`
}
