package entity

type Favorite struct {
	Id        int64 `gorm:"primaryKey"`
	TitleId   int64
	AccountId int64
}
