package db_test

import (
	"testing"

	"github.com/Elyasb14/ghat/pkg/db"
)

func TestDb(t *testing.T) {
	Db := db.InitDB("file:/tmp/db.go")

}
