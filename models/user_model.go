package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `gorm:"column:id;primaryKey;autoIncrement"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `grom:"column:email"`
	Password  []byte `gorm:"column:-"`
	Phone     string `gorm:"column:phone"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
