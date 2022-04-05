package db

import (
	"fmt"
	//"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "db")

type DB struct {
	pool   *sql.DB
	logger *logrus.Entry
}

func New() (*DB, error) {
	fmt.Println("3+++++++++++")
	db := &DB{
		logger: logrus.WithField("module", "db"),
	}

	pool, err := sql.Open("mysql", "root:123456@tcp(192.168.2.65:3306)/shop?charset=utf8")

	fmt.Println("4+++++++++++")
	if err != nil {
		fmt.Println("6+++++++++++")
		return nil, errors.Wrap(err, "could not build mysql config")
	} else {
		fmt.Println("7+++++++++++")
		//log.Info("connect mysql successful. ", err.Info())
		fmt.Println("mysql connect success !")
	}
	fmt.Println("5+++++++++++")

	db.pool = pool

	return db, nil
}

func (db *DB) Connection() *sql.DB {
	return db.pool
}
