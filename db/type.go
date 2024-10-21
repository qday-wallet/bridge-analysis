package db

import (
	"time"
)

const (
	TxFail            = 101
	TxSuccess         = 100
	DepositSuccess    = 2
	Voting            = 3
	VoteFail          = 4
	VoteSuccess       = 5
	MintSuccess       = 6
	MintFail          = 7
	TxTransferSuccess = 8
	Locking           = 9
	WithdrawFail      = 10
)

type Tx struct {
	ID               int64     `json:"id" gorm:"column:id"`
	TxID             string    `json:"tx_id" gorm:"column:tx_id"`
	MintID           string    `json:"mint_id" gorm:"column:mint_id"`
	CreatedAt        time.Time `json:"create_time" gorm:"column:create_time"`
	UpdatedAt        time.Time `json:"update_time" gorm:"column:update_time"`
	FromChainId      string    `json:"from_chain_id" gorm:"column:from_chain_id"`
	ToChainId        string    `json:"to_chain_id" gorm:"column:to_chain_id"`
	FromAddress      string    `json:"from_address" gorm:"column:from_address"`
	ToAddress        string    `json:"to_address" gorm:"column:to_address"`
	ReceiverAddress  string    `json:"receiver_address" gorm:"column:receiver_address"`
	FromTokenAddress string    `json:"from_token_address" gorm:"column:from_token_address"`
	ToTokenAddress   string    `json:"to_token_address" gorm:"column:to_token_address"`
	Status           string    `json:"status" gorm:"column:status"`
	RetryVote        int64     `json:"retry_vote" gorm:"column:retry_vote"`
	RetryExec        int64     `json:"retry_exec" gorm:"column:retry_exec"`

	BridgeId  string `json:"bridge_id" gorm:"column:bridge_id"`
	ChannelId string `json:"channel_id" gorm:"column:channel_id"`
	Timestamp string `json:"timestamp" gorm:"column:timestamp"`
	Value     string `json:"value" gorm:"column:value"`
	Fee       string `json:"fee" gorm:"column:fee"`
}

func (receiver *Tx) TableName() string {
	return "relay.tx"
}
