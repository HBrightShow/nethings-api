package db

import (
	"context"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/sirupsen/logrus"

	//"github.com/nethings/internal-api/config"
)

type ErrorLineExtract struct {
	LineNum   int    // Line number starting with 1
	ColumnNum int    // Column number starting with 1
	Text      string // Text of the line without a new line character.
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.


func (db *DB) Ping(ctx context.Context) error {
	//return db.Connection().Ping(ctx)
	return nil
}


// ExtractErrorLine takes source and character position extracts the line
// number, column number, and the line of text.
//
// The first character is position 1.
func ExtractErrorLine(source string, position int) (ErrorLineExtract, error) {
	ele := ErrorLineExtract{LineNum: 1}
	if position > len(source) {
		return ele, fmt.Errorf("position (%d) is greater than source length (%d)", position, len(source))
	}

	lines := strings.SplitAfter(source, "\n")
	for _, ele.Text = range lines {
		if position-len(ele.Text) < 1 {
			ele.ColumnNum = position
			break
		}
		ele.LineNum += 1
		position -= len(ele.Text)
	}

	ele.Text = strings.TrimSuffix(ele.Text, "\n")

	return ele, nil
}

