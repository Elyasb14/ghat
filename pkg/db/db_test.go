package db_test

import (
	"testing"

	"github.com/Elyasb14/ghat/pkg/db"
)

func TestDb(t *testing.T) {
	Db := db.InitDB("file:/tmp/db.go")

	Db.WriteMSG("hello", "192.168.33.1")
	Db.WriteMSG("fuck", "192.168.33.1")
	Db.WriteMSG("nuts", "10.10.10.1")

	t.Log(Db.ReadDB())
}
