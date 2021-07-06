package sevenhand

import (
	"fmt"
	. "poker_demo/src"
	"sort"
	"strings"
)

type Counter struct{}

func (c *Counter) Start(hand1, hand2 string) int {
	//统计每张牌的出现的次数
	faceCount1, faceCount2 := GetFaceCount(hand1), GetFaceCount(hand2)

	rank1, newHand1 := c.getSevenHandRank(hand1, faceCount1)
	rank2, newHand2 := c.getSevenHandRank(hand2, faceCount2)

	//七张牌
	if rank1 > rank2 {
		return 1
	} else if rank1 < rank2 {
		return 2
	} else {
		switch rank1 {

		//  单张大牌和同花需要依次判断face点数
		case HandRank["单张大牌"]:
			fallthrough
		case HandRank["同花"]: //同花已经抽离了花色 直接比较
			for i := 0; i < 5; i++ {
				fmt.Printf("newHand手牌：%v\t%v\n",newHand1,newHand2)
				//v1, v2 := FaceRank[newHand1[i:i+1]], FaceRank[newHand2[i:i+1]]
				v1, v2 := FaceRank[string(newHand1[i])],FaceRank[string(newHand2[i])]
				fmt.Printf("newHand手牌：%v\t%v\n",string(newHand1[i]),string(newHand2[i]))
				if v1 > v2 {
					return 1
				} else if v1 < v2 {
					return 2
				}
			}
			return 0

		//  顺子和同花顺只需要判断头牌
		case HandRank["顺子"]:
			fallthrough
		case HandRank["同花顺"]:
			rst, _ := Max(FaceRank[newHand1[0:1]], FaceRank[newHand2[0:1]])
			return rst

		case HandRank["皇家同花顺"]:
			return 0

		default:
			return c.equalJudgePair(faceCount1, faceCount2, rank1)
		}
	}
}

//getSevenHandRank 得到牌型
func (c *Counter) getSevenHandRank(hand string, faceCount FaceCount) (int, string) {
	code := GetFaceCountCode(faceCount)
	rank := SevenHandCountCode[code]
	fmt.Printf("getSevenHandRank的牌型排名:%v\t%v\n",code,rank)
	newHand := ""

	//7000 即没有对子，都是散牌，可能是顺子、同花、同花顺、皇家同花顺
	if code == "7000" {
		var str string
		//  同花的点数大于顺子，所以优先判断是不是同花
		if flush, is := c.isFlush(hand); is {
			//  如果是 则判断是不是顺子
			if str, is = c.isStr(flush); is {
				if IsRoyalFlush(str) {
					rank = HandRank["皇家同花顺"]
				} else {
					rank = HandRank["同花顺"]
				}
			} else {
				rank = HandRank["同花"]
			}
		} else if str, is = c.isStr(flush); is {
			rank = HandRank["顺子"]
		}
		newHand = str
	}
	fmt.Printf("getSevenHandRank的牌型排名rank2:%v\n",rank)
	return rank, newHand
}

//isFlush 判断是否可能是同花，顺便把花色去了
func (c *Counter) isFlush(hand string) (string, bool) {
	Sa.Reset()
	var num = 0
	for i := 0; i < len(hand); i += 2 {
		Sa.WriteString(hand[i : i+1]) //去花色,只保留 0 2 4 6 8 10 12
	}
	flag := hand[1]
	for i := 1; i <= len(hand); i += 2{
		if flag == hand[i] {
			num++
		}
	}
	if num == 5 {
		for i := 1; i <= len(hand); i += 2{
			if flag == hand[i] {
				Sa.WriteString(hand[i-1 : i]) //返回都为同花的牌
			}
		}
		return Sa.String(), true
	}else{
		return Sa.String(), false //返回七张牌
	}
}

//  判断无重复字符的是否是顺子，返回排序后的手牌
func (c *Counter) isStr(hand string) (string, bool) {
	hand = Sort(hand)//从大到小排
	isStr := false
	//  由于是无重复子串，通过头部减去尾部可以直接判断是不是顺子
	head1, tail1 := hand[0:1], hand[4:5]
	head2, tail2 := hand[2:3], hand[6:7]
	if FaceRank[head1]-FaceRank[tail1] == 4 ||
	   FaceRank[head2]-FaceRank[tail2] == 4 {
		isStr = true
	}
	//  A5432的特殊情况
/*	if strings.Compare(hand, "A5432") == 0 {
		hand = "5432A"
		isStr = true
	}*/
	if strings.Contains(hand, "A5432") {
		hand = "5432A"
		isStr = true
	}

	return hand, isStr
}

//判断相等的对子牌型、三条、四条
func (c *Counter) equalJudgePair(count1, count2 FaceCount, rank int) int {
	map1 := GetFaceCountMap(count1)
	map2 := GetFaceCountMap(count2)
	var rst, v1, v2 int
	var isEqual bool
	switch rank {
	case HandRank["四条"]:// 0011、1101、3001
		v1, v2 = map1[4][0], map2[4][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			arr1 := []int{map1[1][8],map1[2][3],map1[3][1]}
			sort.Ints(arr1)
			arr2 := []int{map2[1][8],map2[2][3],map2[3][1]}
			sort.Ints(arr2)
			//v1, v2 = map1[1][0], map2[1][0]
			fmt.Printf("sort4应用arr1和arr2：%v\t%v\n",arr1,arr2)
			fmt.Println(map1[1][8],map1[2][3],map1[3][1])
			rst, _ = Max(arr1[2], arr2[2])
		}
	case HandRank["葫芦"]://1020、2110、0210
		v1, v2 = map1[3][1], map2[3][1]
		//v3, v4 = map1[3][1], map2[3][1]
		if rst, isEqual = Max(v1, v2); isEqual {
			arr1 := []int{map1[1][8],map1[1][7],map1[2][3],map1[2][2],map1[3][1]}
			sort.Ints(arr1)
			arr2 := []int{map2[1][8],map2[1][7],map2[2][3],map2[2][2],map2[3][1]}
			sort.Ints(arr2)
			fmt.Printf("sort3应用arr1和arr2：%v\t%v\n",arr1,arr2)
			fmt.Println(map1[1][8],map1[1][7],map1[2][3],map1[2][2],map1[3][1],map2[1][8],map2[1][7],map2[2][3],map2[2][2],map2[3][1])
			//v1 = map1[3][1]
			//v2 = map2[3][1]
			rst, _ = Max(arr1[4], arr2[4])
		}
	case HandRank["三条"]:// 4010
		v1, v2 = map1[3][0], map2[3][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1s, v2s := map1[1], map2[1]
			for i := 0; i < 4; i++ {
				v1, v2 = v1s[i], v2s[i]
				if v1 > v2 {
					rst = 1
					break
				} else if v1 < v2 {
					rst = 2
					break
				}
			}
		}
	case HandRank["两对"]:// 1300 和 3200
		v1, v2 = map1[2][4], map2[2][4]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1, v2 = map1[2][3], map2[2][3]
			if rst, isEqual = Max(v1, v2); isEqual {
				arr1 := []int{map1[1][8],map1[1][7],map1[1][6],map1[2][2]}
				sort.Ints(arr1)
				arr2 := []int{map2[1][8],map2[1][7],map2[1][6],map2[2][2]}
				sort.Ints(arr2)
				fmt.Printf("sort2应用arr1和arr2：%v\t%v\t%v\n",arr1,arr2,map1[1][0])
				fmt.Println(map1[1][8],map1[1][7],map1[1][7],map1[2][2],map2[1][8],map2[1][7],map2[1][6],map2[2][2])
				//v1, v2 = map1[2][2], map1[2][2]
				rst, _ = Max(arr1[3], arr2[3])
			}
		}


	case HandRank["一对"]:// 5100
		v1, v2 = map1[2][0], map2[2][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			vs1, vs2 := map1[1], map2[1]
			for i := 0; i < 5; i++ {
				v1, v2 = vs1[i], vs2[i]
				if v1 > v2 {
					rst = 1
					break
				} else if v1 < v2 {
					rst = 2
					break
				}
			}
		}
	}
	return rst
}
