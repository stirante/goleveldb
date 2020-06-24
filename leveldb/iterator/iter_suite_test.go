package iterator_test

import (
	"testing"

	"github.com/df-mc/golevelb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
