package main

import (
	"fmt"
	"github.com/df-mc/goleveldb/leveldb"
	"github.com/df-mc/goleveldb/leveldb/opt"
	"os"
)

func main() {
	env, b := os.LookupEnv("db_path")
	if !b {
		panic("db_path not set")
	}
	file, err := leveldb.OpenFile(env, &opt.Options{
		Compression:      opt.FlateCompression,
		CompressionLevel: 9,
		BlockSize:        4 * opt.GiB,
	})
	if err != nil {
		panic(err)
	}
	iter := file.NewIterator(nil, nil)
	iter.First()
	fmt.Println(string(iter.Key()))
	if iter.Error() != nil {
		panic(iter.Error())
	}
	for iter.Next() {
		if iter.Error() != nil {
			panic(iter.Error())
		}
		fmt.Println(string(iter.Key()))
	}
	iter.Release()
	err = file.Close()
}
