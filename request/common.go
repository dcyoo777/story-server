package request

import (
	"example/mysql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CommonRequestInterface interface {
	GetAll() (any, error)
	GetOne(any) (any, error)
	Create(any) (any, error)
	Update(any, any) (any, error)
	Delete(any) (any, error)
}

type CommonRequests struct {
	Table          string
	PrimaryKey     string
	DatasourceName string
}

func (c CommonRequests) connectDB() mysql.Mysql {
	return mysql.Mysql{
		Table:          c.Table,
		PrimaryKey:     c.PrimaryKey,
		DatasourceName: c.DatasourceName,
	}
}

func (c CommonRequests) GetAll() (*sqlx.Rows, error) {

	db := c.connectDB()

	result, err := db.GetAll()

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", result)

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

	fmt.Printf("%+v\n", result)

	return result, nil
}

func (c CommonRequests) Update(id any, item interface{}) (any, error) {

	db := c.connectDB()

	result, err := db.Update(item, id)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", result)

	return result, nil
}

func (c CommonRequests) Delete(id any) (any, error) {

	db := c.connectDB()

	result, err := db.Delete(id)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", result)

	return result, nil
}
