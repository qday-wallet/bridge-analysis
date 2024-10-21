package db

import (
	"testing"
	"time"

	"github.com/qday-wallet/bridge-analysis/common/driver"
	"github.com/sunjiangjun/xlog"
)

func init2() *DB {
	log := xlog.NewXLogger()
	conn, err := driver.Open("postgres", "123456789", "190.92.213.101", "postgres", 5432, log)
	if err != nil {
		panic(err)
	}
	return NewDB(conn, log)
}

func TestDB_AssetIncome(t *testing.T) {
	db := init2()
	ps, err := db.AssetIncome(time.Now().Add(-5*24*time.Hour), time.Now())
	if err != nil {
		t.Error(err)
		return
	}

	for _, v := range ps {
		t.Log(v)
	}
}

func TestDB_AssetPay(t *testing.T) {
	db := init2()
	ps, err := db.AssetPay(time.Now().Add(-5*24*time.Hour), time.Now())
	if err != nil {
		t.Error(err)
		return
	}

	for _, v := range ps {
		t.Log(v)
	}
}

func TestDB_QueryTx(t *testing.T) {
	db := init2()
	tx, err := db.QueryTx(93)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tx)
}

func TestDB_QueryTxByFrom(t *testing.T) {
	db := init2()
	txs, total, err := db.QueryTxByFrom("0x30ef9dF39C10C57a478f4c6733c3f210CE17C662", 5, 2)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("total: %d,len:%v", total, len(txs))
	for _, tx := range txs {
		t.Log(tx)
	}
}

func TestDB_QueryTxByHash(t *testing.T) {
	db := init2()
	tx, err := db.QueryTxByHash("0x1b49007858291af0770e7fad0cb30ddcfc38e6f3e23113d5be8819c2f755609b")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tx)
}

func TestDB_QueryTxs(t *testing.T) {
	db := init2()
	tx, err := db.QueryTxs(time.Now().Add(-5*24*time.Hour), time.Now(), []int64{9, 100})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tx)
}
