package models

type Post struct {
	Id     uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Title  string `gorm:"column:title"`
	Desc   string `gorm:"column:desc"`
	Image  string `gorm:"column:image"`
	UserID string `gorm:"column:user_id"`
	User   User   `gorm:"foreignKey:UserID"`
}
