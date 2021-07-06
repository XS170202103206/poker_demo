package fivehand

import (
	. "poker_demo/src"
	"strings"
)

type Counter struct{}

func (c *Counter) Start(hand1, hand2 string) int {
	//统计每张牌的出现的次数
	faceCount1, faceCount2 := GetFaceCount(hand1), GetFaceCount(hand2)

	rank1, newHand1 := c.getSevenHandRank(hand1, faceCount1)
	rank2, newHand2 := c.getSevenHandRank(hand2, faceCount2)

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
				//v1, v2 := FaceRank[newHand1[i:i+1]], FaceRank[newHand2[i:i+1]]
				v1, v2 := FaceRank[string(newHand1[i])],FaceRank[string(newHand2[i])]
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
	rank := FiveHandCount[code]
	newHand := ""

	//5000 即没有对子，都是散牌，可能是顺子、同花、同花顺、皇家同花顺
	if code == "5000" {
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
	return rank, newHand
}

//isFlush 判断是否可能是同花，顺便把花色去了
func (c *Counter) isFlush(hand string) (string, bool) {
	Sa.Reset()
	for i := 0; i < len(hand); i += 2 {
		Sa.WriteString(hand[i : i+1]) //去花色 0 2 4 6 8
	}
	if  hand[1] == hand[3] &&
		hand[3] == hand[5] &&
		hand[5] == hand[7] &&
		hand[7] == hand[9] {
		return Sa.String(), true
	} else {
		return Sa.String(), false
	}
}

//  判断无重复字符的是否是顺子，返回排序后的手牌
func (c *Counter) isStr(hand string) (string, bool) {
	hand = Sort(hand)
	isStr := false
	//  由于是无重复子串，通过头部减去尾部可以直接判断是不是顺子
	head, tail := hand[0:1], hand[4:5]
	if FaceRank[head]-FaceRank[tail] == 4 {
		isStr = true
	}

	//  A5432的特殊情况
	if strings.Compare(hand, "A5432") == 0 {
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
	case HandRank["四条"]:
		v1, v2 = map1[4][0], map2[4][0]  //将4条的牌存入map，就无需判断是AAAAX型 还是 XAAAA型，下面同理
		if rst, isEqual = Max(v1, v2); isEqual {
			v1, v2 = map1[1][0], map2[1][0]
			rst, _ = Max(v1, v2)
		}
	case HandRank["葫芦"]:
		v1, v2 = map1[3][0], map2[3][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1 = map1[2][0]
			v2 = map2[2][0]
			rst, _ = Max(v1, v2)
		}
	case HandRank["三条"]://如果三张牌都相同，比较第四张牌，必要时比较第五张
		v1, v2 = map1[3][0], map2[3][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1s, v2s := map1[1], map2[1]
			for i := 0; i < 2; i++ {
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
	case HandRank["两对"]:
		v1, v2 = map1[2][0], map2[2][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1, v2 = map1[2][1], map2[2][1]
			if rst, isEqual = Max(v1, v2); isEqual {
				v1, v2 = map1[1][0], map1[2][0]
				rst, _ = Max(v1, v2)
			}
		}
	case HandRank["一对"]:
		v1, v2 = map1[2][0], map2[2][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			vs1, vs2 := map1[1], map2[1]
			for i := 0; i < 3; i++ {
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
