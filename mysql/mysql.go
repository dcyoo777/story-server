package mysql

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

func removeDoubleQuote(str string) string {
	return strings.ReplaceAll(str, "\"", "")
}

type Mysql struct {
	Table          string
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

func (mysql Mysql) Select(result any, options any) (*sqlx.Rows, error) {

	db, err := sqlx.Open("mysql", mysql.DatasourceName)

	if err != nil {
		log.Fatalln(err)
	}

	sql, _, _ := goqu.From(mysql.Table).Where().ToSQL()

	rows, err2 := db.Queryx(sql)

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

	rows, err2 := db.Queryx(fmt.Sprintf("SELECT * FROM %s", mysql.Table))

	if err2 != nil {
		return nil, err2
	}

	return rows, nil
}

func (mysql Mysql) GetOne(id any) (*sqlx.Rows, error) {

	db, err := sqlx.Open("mysql", mysql.DatasourceName)

	if err != nil {
		log.Fatalln(err)
	}

	rows, err2 := db.Queryx(fmt.Sprintf("SELECT * FROM %s WHERE %s = %s", mysql.Table, mysql.PrimaryKey, id))

	if err2 != nil {
		return nil, err2
	}

	return rows, nil
}

func (mysql Mysql) Insert(item interface{}) (bool, error) {

	query, _, _ := goqu.Insert(mysql.Table).
		Rows(item).
		ToSQL()

	return mysql.Tx(removeDoubleQuote(query))

}

func (mysql Mysql) Update(item interface{}, id any) (bool, error) {

	query, _, _ := goqu.Update(mysql.Table).
		Set(item).
		Where(goqu.Ex{
			mysql.PrimaryKey: id,
		}).ToSQL()

	return mysql.Tx(removeDoubleQuote(query))

}

func (mysql Mysql) Delete(id any) (bool, error) {

	query, _, _ := goqu.Delete(mysql.Table).
		Where(goqu.Ex{
			mysql.PrimaryKey: id,
		}).ToSQL()

	return mysql.Tx(removeDoubleQuote(query))

}
