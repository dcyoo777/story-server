package mysql

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

func removeDoubleQuote(str string) string {
	return strings.ReplaceAll(str, "\"", "")
}

func whereUuid(id uuid.UUID) string {
	return fmt.Sprintf("WHERE id=UUID_TO_BIN('%s')", id)
}

type Mysql struct {
	Name           string
	PrimaryKey     string
	DatasourceName string
}

func (mysql Mysql) Tx(query string) (bool, error) {
	db, err := sqlx.Open("mysql", mysql.DatasourceName)

	if err != nil {
		fmt.Println(":: Fail to connect DB ::\n", err)
		return false, err
	}

	tx, err2 := db.Begin()

	if err2 != nil {
		fmt.Println(":: Fail to begin DB transaction ::\n", err2)
		return false, err2
	}

	_, err3 := tx.Exec(query)

	if err3 != nil {
		fmt.Println(":: Fail to execute DB transaction ::\n", err3)
		return false, err3
	}

	err4 := tx.Commit()

	if err4 != nil {
		fmt.Println(":: Fail to commit DB transaction ::\n", err4)
		return false, err4
	}

	fmt.Println(":: DB transaction succeed ::")

	return true, nil
}

func (mysql Mysql) Select(options ...exp.Expression) (*sqlx.Rows, error) {

	db, err := sqlx.Open("mysql", mysql.DatasourceName)

	if err != nil {
		log.Fatalln(err)
	}

	sql, _, _ := goqu.From(mysql.Name).Where(
		options...,
	).ToSQL()

	println(strings.ReplaceAll(sql, "\"", ""))

	rows, err2 := db.Queryx(strings.ReplaceAll(sql, "\"", ""))

	println("%+v", rows)

	if err2 != nil {
		return nil, err2
	}

	return rows, nil
}

func (mysql Mysql) GetAll() (*sqlx.Rows, error) {

	db, err := sqlx.Open("mysql", mysql.DatasourceName)

	if err != nil {
		log.Fatalln(err)
	}

	rows, err2 := db.Queryx(fmt.Sprintf("SELECT * FROM %s", mysql.Name))

	if err2 != nil {
		return nil, err2
	}

	return rows, nil
}

func (mysql Mysql) GetOne(id uuid.UUID) (*sqlx.Rows, error) {

	db, err := sqlx.Open("mysql", mysql.DatasourceName)

	if err != nil {
		log.Fatalln(err)
	}

	rows, err2 := db.Queryx(fmt.Sprintf("SELECT * FROM %s %s", mysql.Name, whereUuid(id)))

	if err2 != nil {
		return nil, err2
	}

	return rows, nil
}

func (mysql Mysql) Insert(item interface{}) (bool, error) {

	query, _, _ := goqu.Insert(mysql.Name).
		Rows(item).
		ToSQL()

	return mysql.Tx(removeDoubleQuote(query))

}

func (mysql Mysql) Update(item interface{}, id uuid.UUID) (bool, error) {

	query, _, _ := goqu.Update(mysql.Name).
		Set(item).ToSQL()

	return mysql.Tx(removeDoubleQuote(fmt.Sprintf("%s %s", query, whereUuid(id))))

}

func (mysql Mysql) Delete(id uuid.UUID) (bool, error) {

	query, _, _ := goqu.Delete(mysql.Name).ToSQL()

	return mysql.Tx(removeDoubleQuote(fmt.Sprintf("%s %s", query, whereUuid(id))))

}
