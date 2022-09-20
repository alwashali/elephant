package cache

import (
	"bytes"
	"fmt"
	"log"
	"time"

	opts "github.com/alwashali/elephant/options"
	"github.com/dgraph-io/badger"
)

func GetCachedKeys(db *badger.DB) []string {

	var dbKeys []string

	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			dbKeys = append(dbKeys, string(k))
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return dbKeys
}

func IsChached(db *badger.DB, key []byte) bool {
	//fmt.Println("aaaaa", key)
	found := false
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()

			if bytes.Compare(k, key) == 0 {
				found = true
			}

		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return found
}

func Cache(db *badger.DB, key []byte, value []byte) bool {
	cacheDuraiton, err := time.ParseDuration(opts.TTL)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("New Entry...", cacheDuraiton.String())
	err = db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, value).WithTTL(cacheDuraiton)
		err := txn.SetEntry(e)
		return err
	})
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func GetItem(db *badger.DB, key []byte) []byte {
	var valcopy []byte

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			valcopy = append(valcopy, val...)
			return err
		})

		return err
	})

	if err != nil {
		fmt.Print(err)
	}

	return valcopy
}
