package infra

import "gorm.io/gorm"

type Friend struct {
	FriendID int64 `gorm:"primaryKey" json:"friend_id"`
	UserID   int64 `gorm:"primaryKey" json:"user_id"`
}

type FriendID struct {
	ID int64 `json:"id"`
}

func CreateFriends(db *gorm.DB, friends []Friend) (*[]Friend, error) {
	for _, i := range friends {
		err := db.Create(&i).Error
		if err != nil {
			return nil, err
		}
	}
	return &friends, nil
}

func CreateFriend(db *gorm.DB, friend Friend) (*Friend, error) {
	err := db.Create(friend).Error
	if err != nil {
		return nil, err
	}
	return &friend, nil
}

func GetFriends(db *gorm.DB, userID int64) (*[]Friend, error) {
	var friends []Friend
	err := db.Where(&Friend{UserID: userID}).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	return &friends, nil
}

func DeleteFriends(db *gorm.DB, uid string) error {
	var friend Friend
	err := db.Where("user_id = ?", uid).Delete(&friend).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteFriend(db *gorm.DB, userID string, friendID string) error {
	var friend Friend
	err := db.Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&friend).Error
	if err != nil {
		return err
	}
	return nil
}
