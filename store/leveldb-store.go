package store

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/Loopring/ringminer/orderbook"
)

type LevelDbStore struct {
	Version orderbook.StoreVersion
}

func (s *LevelDbStore) Store() {
	//o := &opt.Options{
	//	Filter: filter.NewBloomFilter(2),
	//}
	db, err := leveldb.OpenFile("./data/", nil)
	defer db.Close()
	if (err != nil) {
		println(err.Error())
	}
	//
	//err = db.Put([]byte("bkey2"), []byte("value2"), nil)
	//err = db.Put([]byte("key3"), []byte("value3"), nil)
	//err = db.Put([]byte("key4"), []byte("value4"), nil)
	//err = db.Put([]byte("key5"), []byte("value5"), nil)
	//err = db.Put([]byte("key6"), []byte("value6"), nil)

	data, err := db.Get([]byte("key1"), nil)

	println(string(data))
	snapShot,err := db.GetSnapshot()
	println(snapShot.String())

	//iter := snapShot.NewIterator(nil, nil)
	iter := db.NewIterator(nil, nil)

	//iter := db.NewIterator(&util.Range{Start: []byte("foo"), Limit: []byte("xoo")}, nil)
	//f := true
	for iter.Next() {
		//if f {
		//	f = false
		//	err = db.Put([]byte("gkey2"), []byte("value3"), nil)
		//}
		key := iter.Key()
		value := iter.Value()
		println("key:" + string(key) + " value:" + string(value))
		key = []byte("key9")
	}
	iter.Release()
	err = iter.Error()

}

func (s *LevelDbStore) GetOrder(hash string) {

}

func (s *LevelDbStore) GetAllOrder(version orderbook.StoreVersion) {

}

func (s *LevelDbStore) SetStoreVersion(version orderbook.StoreVersion) {

}



