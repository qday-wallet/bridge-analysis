package util

import (
	"fmt"
	"math/big"
)

func ToBigInt(s string) (*big.Int, error) {

	// 创建一个新的 big.Int 对象
	num := new(big.Int)

	// 将字符串转换为 big.Int
	_, ok := num.SetString(s, 10) // 第二个参数 10 表示字符串是十进制
	if !ok {
		return nil, fmt.Errorf("Data conversion failed :%v ", s)
	}
	return num, nil
}
