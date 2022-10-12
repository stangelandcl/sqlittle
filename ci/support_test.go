//go:build ci
// +build ci

package ci

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
	"github.com/stangelandcl/sqlittle"
)

func tmpfile(t *testing.T) (string, func()) {
	t.Helper()
	f, err := ioutil.TempFile("", t.Name())
	if err != nil {
		t.Fatal(err)
	}
	file := f.Name()
	f.Close()
	return file, func() {
		os.Remove(file)
	}
}

// Compare is a helper to create a table, and compare sqlite and sqlittle
// queries.
func Compare(
	t *testing.T,
	sqlCreate string,
	sqlSelect string,
	little func(*testing.T, *sqlittle.DB) [][]string,
) {
	t.Helper()

	file, close := tmpfile(t)
	defer close()

	create(t, file, sqlCreate)

	CompareSelect(t, file, sqlSelect, little)
}

func CompareSelect(
	t *testing.T,
	file string,
	sqlSelect string,
	little func(*testing.T, *sqlittle.DB) [][]string,
) {
	t.Helper()

	lite := execute(t, file, sqlSelect)

	db, err := sqlittle.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	have, want := little(t, db), lite
	// t.Logf("have:\n%s\nwant:\n%s\n", spew.Sdump(have), spew.Sdump(want))
	if !reflect.DeepEqual(have, want) {
		t.Errorf("diff:\n%s", diff.LineDiff(spew.Sdump(want), spew.Sdump(have)))
	}

}

func sqlite(file, sql string) (string, error) {
	out, err := exec.Command("sqlite3", "-batch", file, sql).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s: %s", err, out)
	}
	return string(out), nil
}

func create(
	t *testing.T,
	file string,
	sql string,
) {
	t.Helper()
	if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
	if _, err := sqlite(file, sql); err != nil {
		t.Fatal(err)
	}
}

func execute(
	t *testing.T,
	file string,
	sql string,
) [][]string {
	t.Helper()
	r, err := sqlite(file, sql)
	if err != nil {
		t.Fatal(err)
	}
	cr := csv.NewReader(strings.NewReader(r))
	cr.Comma = '|'
	rec, err := cr.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	return rec
}
