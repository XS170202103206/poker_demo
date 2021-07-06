package sevenghost

import (
	. "poker_demo/src"
	"strconv"
	"strings"
)

//Counter 计算7手牌和鬼牌
type Counter struct{}

func (c *Counter) Start(hand1, hand2 string) int {
	faceCount1, faceCount2 := c.getFaceCount(hand1), c.getFaceCount(hand2)

	rank1, newHand1 := c.getHandRank(hand1, faceCount1)
	rank2, newHand2 := c.getHandRank(hand2, faceCount2)

	if rank1 > rank2 {
		return 1
	} else if rank1 < rank2 {
		return 2
	} else {
		switch rank1 {

		case HandRank["单张大牌"]:
			return c.equalJudgeHighCard(faceCount1, faceCount2)

		case HandRank["同花"]:
			//  同花总是挑出来的
			return c.equalJudgeFlush(newHand1, newHand2, faceCount1[15] != 0, faceCount2[15] != 0)

		//  顺子和同花顺只需要判断头牌
		case HandRank["顺子"]:
			fallthrough
		case HandRank["同花顺"]:
			// 顺子总是挑出来的。顺子比较头牌就可以判断大小
			rst, _ := Max(FaceRank[newHand1[0:1]], FaceRank[newHand2[0:1]])
			return rst

		default:
			if len(newHand1) != 0 {
				faceCount1 = c.getFaceCount(newHand1)
			}
			if len(newHand2) != 0 {
				faceCount2 = c.getFaceCount(newHand2)
			}
			return c.equalJudgePair(faceCount1, faceCount2, rank1)
		}
	}
}

//isRoyalFlush 在保证同花的前提下判断是否可能为皇家同花顺
func (*Counter) isRoyalFlush(hand string) bool {
	return strings.Compare(hand, "AKQJT") == 0
}

//getFaceCount 计算手牌中每种牌出现的次数
func (*Counter) getFaceCount(hand string) FaceCount {
	var count FaceCount
	for i := 0; i < len(hand); i += 2 {
		count[FaceRank[hand[i:i+1]]]++
	}
	return count
}

//getFaceCountCode 计算出现1次至4次的牌的code
func (*Counter) getFaceCountCode(count FaceCount) string {
	Sa.Reset()
	var card1, card2, card3, card4 int
	for i, v := range count {
		//  不记录鬼牌
		if i == 15 {
			continue
		}
		switch v {
		case 0:
			continue
		case 1:
			card1++
		case 2:
			card2++
		case 3:
			card3++
		case 4:
			card4++
		}
	}

	Sa.WriteString(strconv.Itoa(card1))
	Sa.WriteString(strconv.Itoa(card2))
	Sa.WriteString(strconv.Itoa(card3))
	Sa.WriteString(strconv.Itoa(card4))
	return Sa.String()
}

//getFaceCountMap 返回一个map，该map的key为出现的次数，v是一个[]int，存储了牌面值
func (*Counter) getFaceCountMap(count FaceCount) map[int][]int {
	countMap := make(map[int][]int, 5)
	countMap[1] = make([]int, 0, 7)
	countMap[2] = make([]int, 0, 3)
	countMap[3] = make([]int, 0, 2)
	countMap[4] = make([]int, 0, 1)
	//  此处i是牌面的值，v是牌面出现的次数，i从14开始，不记录鬼牌
	var i, v int
	for i = 14; i >= 2; i-- {
		v = count[i]
		if v == 0 {
			continue
		} else {
			countMap[v] = append(countMap[v], i)
		}
	}
	return countMap
}

func (c *Counter) getHandRank(hand string, count FaceCount) (int, string) {
	//  鬼牌
	if count[15] == 1 {
		return c.getGhostHandRank(hand, count)
	} else {
		return c.getSevenHandRank(hand, count)
	}
}

func (c *Counter) getGhostHandRank(hand string, faceCount FaceCount) (int, string) {
	code := c.getFaceCountCode(faceCount)
	rank := GhostHandCountCode[code]
	newHand := ""

	//  以下牌型不确定，需要继续判断
	if code == "3010" || code == "2200" || code == "4100" || code == "6000" {
		var straight string
		//  判断是否是同花
		if flush, maybe := c.maybeIsFlush(hand, true); maybe {
			if straight, maybe = c.maybeIsStraightByNoDuplicate(flush, true); maybe {
				if c.isRoyalFlush(straight) {
					rank = HandRank["皇家同花顺"]
				} else {
					rank = HandRank["同花顺"]
				}
				newHand = straight
			} else if code == "4100" || code == "6000" { //  3010,2200最低是四条，所以不需要是同花
				rank = HandRank["同花"]
				//  就算不是顺子，maybeISStraightByNoDuplicate函数也会返回排序后的flush
				newHand = straight
			}
		} else if code == "4100" || code == "6000" {
			//  这两种牌型还要再判断是不是顺子
			if straight, maybe = c.maybeIsStraightByFaceCount(faceCount); maybe {
				rank = HandRank["顺子"]
				newHand = straight
			}
		}
	}
	return rank, newHand
}

//getSevenHandRank 得到牌型，如果有顺子、同花，会返回新的子串
func (c *Counter) getSevenHandRank(hand string, faceCount FaceCount) (int, string) {
	code := c.getFaceCountCode(faceCount)
	rank := SevenHandCountCode[code]
	newHand := ""

	//  以下牌型可能是同花或顺子
	if code == "4010" || code == "3200" || code == "5100" || code == "7000" {
		var straight string
		//  同花的点数大于顺子，所以优先判断是不是同花
		if flush, maybe := c.maybeIsFlush(hand, false); maybe {
			//  如果是则判断是不是顺子
			if straight, maybe = c.isStraightByNoDuplicate(flush, true); maybe {
				if c.isRoyalFlush(straight) {
					rank = HandRank["皇家同花顺"]
				} else {
					rank = HandRank["同花顺"]
				}
			} else {
				rank = HandRank["同花"]
			}
			//  isStraightByNoDuplicate函数会返回排序后的flush
			newHand = straight
		} else if straight, maybe = c.isStraightByCount(faceCount); maybe {
			rank = HandRank["顺子"]
			newHand = straight
		}
	}
	return rank, newHand
}

//maybeIsFlush 判断是否可能是同花，可能时返回最大的同花牌
//传入的字符串要求包含花色信息。返回的字符串无花色信息
func (c *Counter) maybeIsFlush(hand string, isGhost bool) (string, bool) {
	S1.Reset()
	S2.Reset()
	S3.Reset()
	S4.Reset()
	for i := 0; i < len(hand); i += 2 {
		face := hand[i : i+1]
		suit := hand[i+1 : i+2]
		switch suit {
		case "s":
			S1.WriteString(face)
		case "h":
			S2.WriteString(face)
		case "d":
			S3.WriteString(face)
		case "c":
			S4.WriteString(face)
		}
	}

	flush := ""
	length := 5
	//  有鬼牌时4个同花色牌即可成同花
	if isGhost {
		length = 4
	}
	if S1.Len() >= length {
		flush = S1.String()
	} else if S2.Len() >= length {
		flush = S2.String()
	} else if S3.Len() >= length {
		flush = S3.String()
	} else if S4.Len() >= length {
		flush = S4.String()
	}
	return flush, len(flush) > 0
}

//  判断无重复字符串是否是顺子
func (c *Counter) isStraightByNoDuplicate(hand string, needSort bool) (string, bool) {
	if needSort {
		hand = Sort(hand)
	}
	straight := ""
	//  由于是无重复子串，通过头部减去尾部可以直接判断是不是顺子
	var head, tail string
	for i := 0; i < len(hand)-4; i++ {
		head, tail = hand[i:i+1], hand[i+4:i+5]
		if FaceRank[head]-FaceRank[tail] == 4 {
			straight = hand[i : i+5]
			break
		}
	}

	//  A5432的特殊情况
	if len(straight) == 0 && hand[0] == 'A' {
		if strings.Contains(hand, "5432") {
			straight = "5432A"
		}
	}

	isStraight := false
	if len(straight) > 0 {
		hand = straight
		isStraight = true
	}
	return hand, isStraight
}

func (c *Counter) isStraightByCount(count FaceCount) (string, bool) {
	Sa.Reset()
	for i := 14; i >= 2; i-- {
		v := count[i]
		if v == 0 {
			continue
		}
		Sa.WriteString(FaceName[i])
	}

	hand := Sa.String()
	if len(hand) < 5 {
		return "", false
	}

	return c.isStraightByNoDuplicate(hand, false)
}

//  maybeIsStraight 判断有鬼牌的手牌是不是顺子，是顺子时返回最大顺子手牌，没顺子时返回排序后的hand
func (c *Counter) maybeIsStraightByNoDuplicate(hand string, needSort bool) (string, bool) {
	if needSort {
		hand = Sort(hand)
	}
	Sa.Reset()
	var insertFlag bool
	var val, lastVal, curVal int
	var str, curKey string
	//  有鬼牌时手牌长度是4、5、6，所以至多判断3次
	//  有鬼牌时通过前后数字相减差值判断连续，允许出现一次数字不连续且可补位
	for i := 0; i < len(hand)-3; i++ {
		insertFlag = false
		str = hand[i : i+4]

		lastVal = FaceRank[str[0:1]]
		Sa.WriteString(str[0:1])
		for j := 1; j < len(str); j++ {
			curKey = str[j : j+1]
			curVal = FaceRank[curKey]
			val = lastVal - curVal
			if val == 1 {
				Sa.WriteString(curKey)
			} else if val == 2 && insertFlag == false {
				insertFlag = true
				//  将缺的键插入
				Sa.WriteString(FaceName[lastVal-1])
				Sa.WriteString(curKey)
			} else {
				Sa.Reset()
				break
			}
			lastVal = curVal
		}
		//  如果stringBuilder的长度不为0，说明已将找到了顺子
		if Sa.Len() > 0 {
			break
		}
	}

	//  如果有顺子且flag为false，说明要在头部或者尾部插入组成顺子
	if Sa.Len() > 0 && insertFlag == false {
		//  头为A说明是顺子是AKQJ
		if Sa.String()[0] == 'A' {
			Sa.WriteString("T")
		} else {
			hand = Sa.String()
			Sa.Reset()
			Sa.WriteString(FaceName[FaceRank[hand[0:1]]+1])
			Sa.WriteString(hand)
		}
	} else if Sa.Len() == 0 && hand[0] == 'A' {
		//  A5432的特殊情况
		face := hand[1:]
		if strings.Contains(face, "543") ||
			strings.Contains(face, "542") ||
			strings.Contains(face, "532") ||
			strings.Contains(face, "432") {
			Sa.WriteString("5432A")
		}
	}

	isStraight := false
	if Sa.Len() > 0 {
		hand = Sa.String()
		isStraight = true
	}

	return hand, isStraight
}

func (c *Counter) maybeIsStraightByFaceCount(count FaceCount) (string, bool) {
	Sa.Reset()
	//  获取无重复子串
	for i := 14; i >= 2; i-- {
		v := count[i]
		if v == 0 {
			continue
		}
		Sa.WriteString(FaceName[i])
	}

	if Sa.Len() < 4 {
		return "", false
	}
	return c.maybeIsStraightByNoDuplicate(Sa.String(), false)
}

func (c *Counter) equalJudgeHighCard(count1, count2 FaceCount) int {
	m1, m2 := c.getFaceCountMap(count1)[1], c.getFaceCountMap(count2)[1]

	rst := 0
	var v1, v2 int
	for i := 0; i < 5; i++ {
		v1, v2 = m1[i], m2[i]
		if v1 > v2 {
			rst = 1
			break
		} else if v1 < v2 {
			rst = 2
			break
		}
	}
	return rst
}

func (c *Counter) equalJudgeFlush(hand1, hand2 string, isGhost1, isGhost2 bool) int {
	if isGhost1 {
		hand1 = c.fillFlush(hand1)
	}

	if isGhost2 {
		hand2 = c.fillFlush(hand2)
	}

	rst := 0
	var v1, v2 int
	for i := 0; i < 5; i++ {
		v1, v2 = FaceRank[hand1[i:i+1]], FaceRank[hand2[i:i+1]]
		if v1 > v2 {
			rst = 1
			break
		} else if v1 < v2 {
			rst = 2
			break
		}
	}
	return rst
}

//fillFlush 4手牌时填充为5手牌同花，5手牌时换掉一个牌成为最大牌(有鬼牌)
//要求传入的手牌是有序的
func (c *Counter) fillFlush(hand string) string {
	for i := 14; i >= 2; i-- {
		insertFace := FaceName[i]
		if !strings.Contains(hand, insertFace) {
			Sa.Reset()
			for j := 0; j < len(hand); j++ {
				face := hand[j : j+1]
				if FaceRank[insertFace] > FaceRank[face] {
					Sa.WriteString(insertFace)
					Sa.WriteString(hand[j:])
					break
				} else {
					Sa.WriteString(face)
				}
			}
			break
		}
	}
	newHand := Sa.String()
	if len(newHand) >= 5 {
		newHand = newHand[0:5]
	}
	return newHand
}

func (c *Counter) equalJudgePair(count1, count2 FaceCount, rank int) int {
	//  FaceRank[15]=X，如果15号位有值说明是鬼牌
	isGhost1 := count1[15] != 0
	isGhost2 := count2[15] != 0
	map1 := c.getFaceCountMap(count1)
	map2 := c.getFaceCountMap(count2)
	var rst, v1, v2 int
	var isEqual bool
	switch rank {
	case HandRank["一对"]:
		//  如果鬼牌是一对，说明现在还没有对子，挑出一张最大的单牌成对子
		if isGhost1 {
			face := map1[1][0]
			map1[1] = map1[1][1:] //去除
			map1[2] = append(map1[2], face)
		}
		if isGhost2 {
			face := map2[1][0]
			map2[1] = map2[1][1:] //去除
			map2[2] = append(map2[2], face)
		}

		v1, v2 = map1[2][0], map2[2][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			vs1, vs2 := map1[1], map2[1]
			for i := 0; i < 3; i++ {
				v1 = vs1[i]
				v2 = vs2[i]
				if v1 > v2 {
					rst = 1
					break
				} else if v1 < v2 {
					rst = 2
					break
				}
			}
		}

	//  鬼牌永远不可能是两对，因为可以组成三条
	case HandRank["两对"]:
		v1s, v2s := map1[2], map2[2]
		v1, v2 = v1s[0], v2s[0]

		if rst, isEqual = Max(v1, v2); isEqual {
			v1, v2 = v1s[1], v2s[1]
			if rst, isEqual = Max(v1, v2); isEqual {
				//  可能有3个两对，即AABBCCD牌型
				var v3, v4 int
				if len(v1s) == 3 {
					v3 = v1s[2]
				}
				v4 = map1[1][0]
				if v3 > v4 {
					v1 = v3
				} else {
					v1 = v4
				}

				if len(v2s) == 3 {
					v3 = v2s[2]
				}
				v4 = map2[1][0]
				if v3 > v4 {
					v2 = v3
				} else {
					v2 = v4
				}

				rst, _ = Max(v1, v2)
			}
		}

	case HandRank["三条"]:
		//  鬼牌为三条时将出现两次牌填充至三次即可
		if isGhost1 {
			v1 = map1[2][0]
			map1[2] = map1[2][1:] //去除
			map1[3] = append(map1[3], v1)
		}
		if isGhost2 {
			v1 = map2[2][0]
			map2[2] = map2[2][1:] //去除
			map2[3] = append(map2[3], v1)
		}

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

	case HandRank["葫芦"]:
		//  鬼牌为葫芦时缺少一个出现三次牌，将一个两次牌填充至三次牌即可
		if isGhost1 {
			v1 = map1[2][0]
			map1[2] = map1[2][1:]
			map1[3] = append(map1[3], v1)
		}
		if isGhost2 {
			v1 = map2[2][0]
			map2[2] = map2[2][1:]
			map2[3] = append(map2[3], v1)
		}

		v1, v2 = map1[3][0], map2[3][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			//  7手牌时候葫芦可能没有两次牌，如AAABBBC
			if len(map1[2]) == 0 {
				v1 = map1[3][1]
			} else {
				v1 = map1[2][0]
			}
			if len(map2[2]) == 0 {
				v2 = map2[3][1]
			} else {
				v2 = map2[2][0]
			}
			rst, _ = Max(v1, v2)
		}

	//  鬼牌为四条时可能缺牌，也可能不缺牌
	case HandRank["四条"]:
		if isGhost1 && len(map1[4]) == 0 {
			v1 = map1[3][0]
			map1[3] = map1[3][1:] //去除
			map1[4] = append(map1[4], v1)
		}
		if isGhost2 && len(map2[4]) == 0 {
			v1 = map2[3][0]
			map2[3] = map2[3][1:]
			map2[4] = append(map2[4], v1)
		}

		v1, v2 = map1[4][0], map2[4][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1, v2 = map1[1][0], map2[1][0]
			rst, _ = Max(v1, v2)
		}
	}
	return rst
}