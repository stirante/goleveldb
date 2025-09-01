package main

import (
	"fmt"
	"os"

	"github.com/df-mc/goleveldb/leveldb"
	"github.com/df-mc/goleveldb/leveldb/opt"
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
	//os.Mkdir("out", 0755)
	fmt.Println(ToHex(iter.Key()))
	//os.WriteFile("out/"+ToHex(iter.Key()), iter.Value(), 0644)
	if iter.Error() != nil {
		panic(iter.Error())
	}
	for iter.Next() {
		if iter.Error() != nil {
			panic(iter.Error())
		}
		fmt.Println(ToHex(iter.Key()))
		//os.WriteFile("out/"+ToHex(iter.Key()), iter.Value(), 0644)
	}
	iter.Release()
	err = file.Close()
}

func ToHex(bytes []byte) string {
	hex := ""
	for _, b := range bytes {
		hex += fmt.Sprintf("%02x", b)
	}
	return hex
}
