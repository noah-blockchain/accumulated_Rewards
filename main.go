package main

import (
	`fmt`
	`math/big`
	`os`
	
	"golang-book/rewards/core"
	"golang-book/rewards/types"
)

const lastBlock = 44150000 // 7 лет
// 2000 =
// 100 =

var startHeight uint64 = 0

func GetRewardForBlock(blockHeight uint64, firstReward uint64, lastReward uint64) *big.Int {
	blockHeight += startHeight
	
	if blockHeight > lastBlock {
		return big.NewInt(0)
	}
	
	if blockHeight == lastBlock {
		return helpers.NoahToQNoah(big.NewInt(int64(lastReward)))
	}
	
	reward := big.NewInt(int64(firstReward))
	reward.Sub(reward, big.NewInt(int64(blockHeight/200000)))
	
	if reward.Cmp(types.Big0) < 1 {
		return helpers.NoahToQNoah(big.NewInt(1))
	}
	
	return helpers.NoahToQNoah(reward)
}

func main() {
	var finishBlock, firstReward, lastReward uint64
	
	fmt.Println("duration of emission")
	fmt.Fscan(os.Stdin, &finishBlock)
	fmt.Println("first reward")
	fmt.Fscan(os.Stdin, &firstReward)
	fmt.Println("last reward")
	fmt.Fscan(os.Stdin, &lastReward)
	
	accumulatedRewards := big.NewInt(int64(firstReward))
	
	for i := 0; i < int(finishBlock); i++ {
		blockReward := GetRewardForBlock(uint64(i), firstReward, lastReward)
		accumulatedRewards.Add(accumulatedRewards, blockReward)
	}
	fmt.Println(accumulatedRewards)
}