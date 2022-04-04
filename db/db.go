package db

import (
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
	db := &DB{
		logger: logrus.WithField("module", "db"),
	}

	pool, err := sql.Open("mysql", "root:123456@tcp(192.168.2.65:3306)/shop?charset=utf8")
	if err != nil {
		return nil, errors.Wrap(err, "could not build mysql config")
	} else {
		log.Error("connect mysql successful. ", err.Error())
	}

	db.pool = pool

	return db, nil
}

func (db *DB) Connection() *sql.DB {
	return db.pool
}
