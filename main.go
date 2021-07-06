package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	. "poker_demo/src"
	"poker_demo/src/fivehand"
	//"poker_demo/src/sevenhand"
	"time"
)

// 结果比较在main_test.go里
func main() {
	fiveHandSpendTime := testFiveHandsSpend("./input/match_result.json")
	//sevenHandSpendTime := testSevenHandsSpend("./input/seven_cards_with_ghost.json")
	//ghostHandSpendTime := testGhostHandsSpend("./input/seven_cards_with_ghost.result.json")

	/*fmt.Printf("五手牌耗时：%v\n"+
		"七手牌耗时：%v\n"+
		"癞子牌耗时：%v\n",
		fiveHandSpend, sevenHandSpendTime, ghostHandSpendTime)*/
	fmt.Printf("五张牌的耗时：%v\n",fiveHandSpendTime)

}

func testFiveHandsSpend(filePath string) time.Duration {
	var matches Matches
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &matches); err != nil {
		panic(err)
	}
	counter := fivehand.Counter{}
	var rst int

	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		if rst = counter.Start(v.Hand1, v.Hand2); rst != v.Result {
			fmt.Printf("Resuel not equal.Hand1:%s Hand2:%s Want:%d Output:%d",
				v.Hand1, v.Hand2, v.Result, rst,
			)
		}
	}
	return time.Since(startTime)
}

//func testSevenHandsSpend(filePath string) time.Duration {
//	var matches Matches
//	file, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		panic(err)
//	}
//	if err = json.Unmarshal(file, &matches); err != nil {
//		panic(err)
//	}
//	counter := sevenhand.Counter{}
//	var rst int
//
//	startTime := time.Now()
//	for _, v := range matches.MatchSlice {
//		if rst = counter.Start(v.Hand1, v.Hand2); rst != v.Result {
//			fmt.Printf("Resuel not equal.Hand1:%s Hand2:%s Want:%d Output:%d",
//				v.Hand1, v.Hand2, v.Result, rst,
//			)
//		}
//	}
//	return time.Since(startTime)
//}
//
//func testGhostHandsSpend(filePath string) time.Duration {
//	var matches Matches
//	file, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		panic(err)
//	}
//	if err = json.Unmarshal(file, &matches); err != nil {
//		panic(err)
//	}
//	counter := sevenhand.Counter{}
//	var rst int
//
//	startTime := time.Now()
//	for _, v := range matches.MatchSlice {
//		if rst = counter.Start(v.Hand1, v.Hand2); rst != v.Result {
//			fmt.Printf("Resuel not equal.Hand1:%s Hand2:%s Want:%d Output:%d",
//				v.Hand1, v.Hand2, v.Result, rst,
//			)
//		}
//	}
//	return time.Since(startTime)
//}
