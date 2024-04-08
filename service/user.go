package service

import (
	"example/request"
	"github.com/jmoiron/sqlx"
)

type UserToDB struct {
	UserName  string
	UserBio   string
	UserImage string
}

type User struct {
	UserId    int `db:"user_id"`
	UserName  string
	UserBio   string
	UserImage string
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (user User) ToDB() UserToDB {
	return UserToDB{
		UserName:  user.UserName,
		UserBio:   user.UserBio,
		UserImage: user.UserImage,
	}
}

type UserCommonRequests struct {
	request.CommonRequests
}

func CastToUser(item *sqlx.Rows) User {
	var p User
	scanErr := item.StructScan(&p)
	if scanErr != nil {
		return User{}
	}
	return p
}

func (c UserCommonRequests) GetAll() (any, error) {

	var result []User

	items, err := c.CommonRequests.GetAll()

	if err != nil {
		return result, err
	}

	for items.Next() {
		result = append(result, CastToUser(items))
	}

	return result, nil
}

func (c UserCommonRequests) GetOne(userId any) (any, error) {

	var result User

	items, err := c.CommonRequests.GetOne(userId)

	if err != nil {
		return result, err
	}
	for items.Next() {
		result = CastToUser(items)
	}

	return result, nil
}

var UserCommonReq UserCommonRequests
