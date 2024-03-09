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

func (mysql Mysql) GetAll(table string, items *any) ([]any, error) {

	//db, err := sqlx.Open("mysql", mysql.DatasourceName)
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//rows, err2 := db.Queryx(fmt.Sprintf("SELECT story_id, title, content, created_at, updated_at FROM %s", table))
	////var result []any
	//for rows.Next() {
	//	var p struct{}
	//	scanErr := rows.StructScan(&p)
	//	if scanErr != nil {
	//		log.Fatalln(scanErr)
	//	}
	//	items = append(result, p)
	//}
	//
	//if err2 != nil {
	//	return nil, err2
	//}
	//
	//return result, nil

	return nil, nil
}

func (mysql Mysql) GetOne(id any) (any, error) {

	db, err := sqlx.Open("mysql", mysql.DatasourceName)

	if err != nil {
		log.Fatalln(err)
	}

	var item any

	err2 := db.Get(item, fmt.Sprintf("SELECT * FROM %s WHERE %s = %d", mysql.Table, mysql.PrimaryKey, id))

	if err != nil {
		fmt.Println("Failed to get user:", err)
		return false, err2
	}

	return item, nil
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
