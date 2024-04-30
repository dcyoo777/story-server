package appUser

import (
	"example/request"
	"fmt"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CommonRequests struct {
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

func (c CommonRequests) Select(options ...exp.Expression) (any, error) {

	var result []User

	items, err := c.CommonRequests.Select(options...)

	fmt.Printf("ITEMS : %v\n", items)
	fmt.Printf("ERROR : %v\n", err)

	if err != nil {
		return result, err
	}

	for items.Next() {
		result = append(result, CastToUser(items))
	}

	return result, nil
}

func (c CommonRequests) GetOne(userId uuid.UUID) (any, error) {

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
