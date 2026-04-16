package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

type StateDB struct {
	path string
	file *os.File
}

func NewStateDB(path string) *StateDB {
	return &StateDB{path: path}
}

func (db *StateDB) Open() error {
	os.MkdirAll(filepath.Dir(db.path), 0755)
	f, err := os.OpenFile(db.path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	db.file = f
	return nil
}

func (db *StateDB) Put(key []byte, value []byte) error {
	binary.Write(db.file, binary.BigEndian, int64(len(key)))
	db.file.Write(key)
	binary.Write(db.file, binary.BigEndian, int64(len(value)))
	db.file.Write(value)
	return nil
}

func (db *StateDB) Close() {
	db.file.Close()
}

func main() {
	db := NewStateDB("./data/state.db")
	db.Open()
	defer db.Close()
	fmt.Println("状态数据库已连接")
}
