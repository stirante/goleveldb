package leveldb

import (
	"testing"

	"github.com/dragonfly-tech/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
