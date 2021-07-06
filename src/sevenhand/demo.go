package sevenhand

import (
	"fmt"
	. "poker_demo/src"

)

type Counter1 struct{}

func (c *Counter1) Start(hand1, hand2 string) int {
	//统计每张牌的出现的次数
	faceCount1, faceCount2 := GetFaceCount(hand1), GetFaceCount(hand2)
     arr1 := [...]int{0,0,0,0}
	 arr2 := [...]int{0,0,0,0}
	for i := 0;i < len(faceCount1);i++ {
		if faceCount1[i] == 4 {
			arr1[3]++
		} else if faceCount1[i] == 3 {
			arr1[2]++
		}else if faceCount1[i] == 2 {
			arr1[1]++
		}else if faceCount1[i] == 1 {
			arr1[0]++
		}
	}
	for i := 0;i < len(faceCount2);i++ {
		if faceCount1[i] == 4 {
			arr2[3]++
		} else if faceCount1[i] == 3 {
			arr2[2]++
		}else if faceCount1[i] == 2 {
			arr2[1]++
		}else if faceCount1[i] == 1 {
			arr2[0]++
		}
	}
	fmt.Println("数组：",arr1,arr2)
	return 0
}
//getSevenHandRank 得到牌型
func (c *Counter1) getSevenHandRank(hand string, faceCount FaceCount) (int, string) {
	code := GetFaceCountCode(faceCount)
	rank := SevenHandCountCode[code]
	fmt.Printf("getSevenHandRank的牌型排名:%v\t%v\n",code,rank)
	newHand := ""

	//7000 即没有对子，都是散牌，可能是顺子、同花、同花顺、皇家同花顺
	if code == "7000" {
		//var str string
		////  同花的点数大于顺子，所以优先判断是不是同花
		//if flush, is := c.isFlush(hand); is {
		//	//  如果是 则判断是不是顺子
		//	if str, is = c.isStr(flush); is {
		//		if IsRoyalFlush(str) {
		//			rank = HandRank["皇家同花顺"]
		//		} else {
		//			rank = HandRank["同花顺"]
		//		}
		//	} else {
		//		rank = HandRank["同花"]
		//	}
		//} else if str, is = c.isStr(flush); is {
		//	rank = HandRank["顺子"]
		//}
		//newHand = str
	}
	fmt.Printf("getSevenHandRank的牌型排名rank2:%v\n",rank)
	return rank, newHand
}
