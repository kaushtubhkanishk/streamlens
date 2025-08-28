package main

import (
	"github.com/kaushtubhkanishk/streamlens/internal/btree"
)

func main() {
	err := btree.SplitLine("SELECT * FROM t1")
	if err != nil {
		panic(err)
	}
}
