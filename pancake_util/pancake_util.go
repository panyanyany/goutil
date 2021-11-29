package pancake_util

import "math/big"

type Token struct {
	Addr                  string
	Name                  string
	Symbol                string
	Decimals              *big.Int
	TotalSupply           *big.Int `json:"totalSupply"`
	BalanceOfPair         *big.Int `json:"balanceOfPair"`
	BusdPrice             *big.Int `json:"busdPrice"`
	TotalBusdAmountOfPair *big.Int `json:"totalBusdAmountOfPair"`
}

type Pair struct {
	PairAddress     string `json:"pairAddress"`
	Name            string
	Symbol          string
	Symbol1         string
	Symbol0         string
	Decimals        int64
	TotalSupply     *big.Int `json:"totalSupply"`
	Token0          *Token
	Token1          *Token
	TotalBusdAmount *big.Int `json:"totalBusdAmount"`
	FarmTvl         *big.Int `json:"farmTVL"`
}
