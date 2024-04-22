package request

import (
	"example/mysql"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"
)

type CommonRequestInterface interface {
	GetName() string
	GetPrimaryKey() string
	Select(...exp.Expression) (any, error)
	GetOne(any) (any, error)
	Create(any) (any, error)
	Update(any, any) (any, error)
	Delete(any) (any, error)
}

type CommonRequests struct {
	Name           string
	PrimaryKey     string
	DatasourceName string
}

func (c CommonRequests) connectDB() mysql.Mysql {
	return mysql.Mysql{
		Name:           c.Name,
		PrimaryKey:     c.PrimaryKey,
		DatasourceName: c.DatasourceName,
	}
}

func (c CommonRequests) GetName() string {
	return c.Name
}

func (c CommonRequests) GetPrimaryKey() string {
	return c.PrimaryKey
}

func (c CommonRequests) Select(options ...exp.Expression) (*sqlx.Rows, error) {
	db := c.connectDB()

	result, err := db.Select(options...)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (c CommonRequests) GetOne(id any) (*sqlx.Rows, error) {

	db := c.connectDB()

	result, err := db.GetOne(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c CommonRequests) Create(item interface{}) (any, error) {

	db := c.connectDB()

	result, err := db.Insert(item)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c CommonRequests) Update(id any, item interface{}) (any, error) {

	db := c.connectDB()

	result, err := db.Update(item, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c CommonRequests) Delete(id any) (any, error) {

	db := c.connectDB()

	result, err := db.Delete(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}
