package lrcdb_test

import (
	"testing"
	"github.com/Loopring/ringminer/lrcdb"
	"os"
	"path/filepath"
	"strconv"
)

const dbname = "leveldb"

var sep = func() string {return string(filepath.Separator)}

func file() string {
	gopath := os.Getenv("GOPATH")
	proj := "github.com/Loopring/ringminer"
	return gopath + sep() + "src" + sep() + proj + sep() + dbname
}

func getdb() lrcdb.Database {
	return lrcdb.NewDB(file(), 12,12)
}

func TestLDBDatabase_Path(t *testing.T) {
	path := lrcdb.NewDB(file(),12,12).Path()
	t.Log("db path is:",path)
}

func TestLDBDatabase_Put(t *testing.T) {
	ldb := getdb()
	ldb.Put([]byte("key_1"), []byte("value_2"))
}

func TestLDBDatabase_Get(t *testing.T) {
	ldb := getdb()
	if value, err := ldb.Get([]byte("key_1")); err != nil {
		t.Log(err.Error())
	} else {
		t.Log("value:", string(value))
	}
}

func TestLDBDatabase_Delete(t *testing.T) {
	ldb := getdb()
	if err := ldb.Delete([]byte("k1")); err != nil {
		t.Log(err.Error())
	}
}

func TestLDBDatabase_Close(t *testing.T) {
	ldb := getdb()
	ldb.Put([]byte("k3"), []byte("v3"))
	ldb.Close()
	if value, err := ldb.Get([]byte("k3")); err != nil {
		t.Log(err.Error())
	} else {
		t.Log(string(value))
	}
}

/////////////////////////////////////////////////////////////////////////////////////
// batch 相关操作
// 这里要注意，batch是一次性的，put&write在一起操作
// batch不能再次寻址
/////////////////////////////////////////////////////////////////////////////////////
func getbatch() lrcdb.Batch {
	ldb := getdb()
	return ldb.NewBatch()
}

func TestLdbBatch_Put(t *testing.T) {
	batch := getbatch()
	for i := 1; i < 100; i++ {
		sn := strconv.Itoa(i)
		batch.Put([]byte("key_"+sn), []byte("value_"+sn))
	}
	if err := batch.Write(); err != nil {
		t.Log(err.Error())
	}
}

/////////////////////////////////////////////////////////////////////////////////////
// table 相关操作
// 这里要注意，table是持久化的，可寻址
// 即便是table和ldb的key相同，也会存储到不同的地方
/////////////////////////////////////////////////////////////////////////////////////

func gettable() lrcdb.Database {
	ldb := getdb()
	return lrcdb.NewTable(ldb, "lrc_test")
}

func TestTable_Put(t *testing.T) {
	table := gettable()
	table.Put([]byte("key_1"), []byte("value_1"))
}

func TestTable_Get(t *testing.T) {
	table := gettable()
	if value, err := table.Get([]byte("key_1")); err != nil {
		t.Log(err.Error())
	} else {
		t.Log("value:", string(value))
	}
}

/////////////////////////////////////////////////////////////////////////////////////
// tablebatch 相关操作
// tablebatch跟batch类似
/////////////////////////////////////////////////////////////////////////////////////

func gettablebatch() lrcdb.Batch {
	table := gettable()
	return table.NewBatch()
}

func TestTableBatch_Put(t *testing.T) {
	tablebatch := gettablebatch()
	for i := 1; i < 100; i++ {
		sn := strconv.Itoa(i)
		tablebatch.Put([]byte("key_"+sn), []byte("valuevalue_"+sn))
	}
	tablebatch.Write()
}

