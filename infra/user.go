package infra

import (
	"gorm.io/gorm"
)

type User struct {
	ID    int64  `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type UserModel struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Age     int     `json:"age"`
	Friends []int64 `json:"friends"`
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) GetUsers(db *gorm.DB) (*[]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *User) GetUser(db *gorm.DB, userID string) (*User, error) {
	//var user User
	err := db.Where("id = ?", userID).Take(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) UpdateUser(db *gorm.DB, uid string) (*User, error) {
	var user User
	userToUpdate, err := user.GetUser(db, uid)
	if err != nil {
		return nil, err
	}

	userToUpdate.Name = u.Name
	userToUpdate.Email = u.Email
	userToUpdate.Age = u.Age

	err = db.Where("id = ?", uid).Updates(userToUpdate).Error
	if err != nil {
		return nil, err
	}

	return userToUpdate, nil

}

func (u *User) DeleteUser(db *gorm.DB, uid string) (*User, error) {
	var user User
	err := db.Where("id = ?", uid).Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
