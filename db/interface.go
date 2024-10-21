package db

import "time"

type Point struct {
	Hour   string
	Amount string
}
type TxManager interface {
	QueryTx(id int64) (*Tx, error)
	QueryTxs(start, end time.Time, status []int64) ([]*Tx, error)
	QueryTxByFrom(from string, page, index int64) ([]*Tx, int64, error)
	QueryTxByHash(bridgeId string) (*Tx, error)

	AssetIncome(start, end time.Time) ([]*Point, error)
	AssetPay(start, end time.Time) ([]*Point, error)
}
