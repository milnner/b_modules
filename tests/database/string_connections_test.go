package database

import (
	"errors"
	"testing"

	"github.com/milnner/b_modules/database"
)

func TestSetRoot(t *testing.T) {
	d := database.NewDatabaseConn()
	str := "OIOIOIOI"

	database.SetRoot(d, str)

	if d.Class.GetDelete() != str {
		t.Fatal(errors.New("String errada"))
	}

}
