package table

import (
	"testing"

	"github.com/dragonfly-tech/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
