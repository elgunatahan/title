package entity

type Account struct {
	Id           int64 `gorm:"primaryKey"`
	Username     string
	Password     string
	AccountRoles []AccountRole
}
