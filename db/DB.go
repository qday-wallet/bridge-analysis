package db

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/qday-wallet/bridge-analysis/common/util"
	"github.com/sunjiangjun/xlog"
	"gorm.io/gorm"
)

type DB struct {
	core *gorm.DB
	log  *xlog.XLog
}

func (db *DB) AssetIncome(start, end time.Time) ([]*Point, error) {
	var txs []*Tx

	sql := "create_time >=? and create_time<=? and status like '%" + fmt.Sprintf("%v", TxSuccess) + "%' and to_chain_id=?"
	err := db.core.Model(Tx{}).Where(sql, start, end, "1001").Order("create_time desc").Scan(&txs).Error
	if err != nil {
		return nil, err
	}
	mp := make(map[string]*big.Int, 3)
	for end.After(start) {
		k := start.Format("2006-01-02 15")
		mp[k] = big.NewInt(0)
		start = start.Add(1 * time.Hour)
	}

	for _, tx := range txs {
		k := tx.CreatedAt.Format("2006-01-02 15")
		if v, ok := mp[k]; ok {
			amount, err := util.ToBigInt(tx.Value)
			if err != nil {
				continue
			}
			v = v.Add(v, amount)
		}
	}

	list := make([]*Point, 0, 10)
	for k, v := range mp {
		list = append(list, &Point{
			Hour:   k,
			Amount: v.String(),
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Hour < list[j].Hour
	})

	return list, nil

}

func (db *DB) AssetPay(start, end time.Time) ([]*Point, error) {
	var txs []*Tx
	sql := "create_time >=? and create_time<=? and status like '%" + fmt.Sprintf("%v", TxSuccess) + "%' and from_chain_id=?"
	err := db.core.Model(Tx{}).Where(sql, start, end, "1001").Order("create_time desc").Scan(&txs).Error
	if err != nil {
		return nil, err
	}
	mp := make(map[string]*big.Int, 3)
	for end.After(start) {
		k := start.Format("2006-01-02 15")
		mp[k] = big.NewInt(0)
		start = start.Add(1 * time.Hour)
	}

	for _, tx := range txs {
		k := tx.CreatedAt.Format("2006-01-02 15")
		if v, ok := mp[k]; ok {
			amount, err := util.ToBigInt(tx.Value)
			if err != nil {
				continue
			}
			v = v.Add(v, amount)
		}
	}

	list := make([]*Point, 0, 10)
	for k, v := range mp {
		list = append(list, &Point{
			Hour:   k,
			Amount: v.String(),
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Hour < list[j].Hour
	})

	return list, nil
}

func (db *DB) QueryTx(id int64) (*Tx, error) {
	var tx Tx
	err := db.core.Model(Tx{}).Where("id=?", id).First(&tx).Error
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func (db *DB) QueryTxs(start, end time.Time, status []int64) ([]*Tx, error) {
	var txs []*Tx

	sql := strings.Builder{}
	sql.WriteString(" create_time>=? and create_time<=?	and ")

	sql.WriteString("(")

	for index, v := range status {
		temp := "status like '%" + fmt.Sprintf("%v", v) + "%'"
		if index != len(status)-1 {
			temp = temp + " or "
		}
		sql.WriteString(temp)
	}

	sql.WriteString(")")

	err := db.core.Model(Tx{}).Where(sql.String(), start, end).Order("create_time desc").Scan(&txs).Error
	if err != nil {
		return nil, err
	}
	return txs, nil
}

func (db *DB) QueryTxByFrom(from string, pageSize, pageNumber int64) ([]*Tx, int64, error) {

	var total int64
	err := db.core.Model(Tx{}).Where("from_address =? or to_address=?", from, from).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var txs []*Tx
	offset := (pageNumber - 1) * pageSize
	err = db.core.Model(Tx{}).Where("from_address =? or to_address=?", from, from).Order("create_time desc").Limit(int(pageSize)).Offset(int(offset)).Scan(&txs).Error
	if err != nil {
		return nil, total, err
	}
	return txs, total, nil
}

func (db *DB) QueryTxByHash(bridgeId string) (*Tx, error) {
	var tx Tx
	err := db.core.Model(Tx{}).Where("bridge_id=?", bridgeId).First(&tx).Error
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func NewDB(db *gorm.DB, log *xlog.XLog) *DB {
	return &DB{
		core: db,
		log:  log,
	}
}
