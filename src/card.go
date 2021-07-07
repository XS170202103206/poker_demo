package src

import (
	"strconv"
	"strings"
)

//FaceCount 用于记录手牌每个牌面出现的次数。下标为牌面，值为次数
//一共有13种牌(含鬼牌)，最小牌在map中值为2，最大为15，为了方便计算，数组长度为16
type FaceCount [16]int

var (
	Sa strings.Builder
	S1 strings.Builder
	S2 strings.Builder
	S3 strings.Builder
	S4 strings.Builder
)

func init() {
	Sa.Grow(5)
	S1.Grow(5)
	S2.Grow(5)
	S3.Grow(5)
	S4.Grow(5)
}

var FaceRank = map[string]int {
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
	"X": 15,
}

var FaceName = map[int]string {
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "T",
	11: "J",
	12: "Q",
	13: "K",
	14: "A",
	15: "X",
}

var HandRank = map[string]int {
	"皇家同花顺": 10,
	"同花顺":   9,
	"四条":    8,
	"葫芦":    7,
	"同花":    6,
	"顺子":    5,
	"三条":    4,
	"两对":    3,
	"一对":    2,
	"单张大牌":    1,
}

var FiveHandCount = map[string]int {
	"1001": HandRank["四条"], // 8
	"0110": HandRank["葫芦"], // 7
	"2010": HandRank["三条"], // 4
	"1200": HandRank["两对"], // 3
	"3100": HandRank["一对"], // 2
	//  非确定，可能为顺子、同花、同花顺、皇家同花顺
	"5000": HandRank["单张大牌"], // 1
}

var SevenHandCountCode = map[string]int {
	"0011": HandRank["四条"],
	"1101": HandRank["四条"],
	"3001": HandRank["四条"],

	"1020": HandRank["葫芦"],
	"2110": HandRank["葫芦"],
	"0210": HandRank["葫芦"],
	"1300": HandRank["两对"],
	//  非确定
	"4010": HandRank["三条"], //  AAABCDE 可能是同花6 或 顺子5
	"3200": HandRank["两对"], //  AABBCDE 可能是同花或顺子
	"5100": HandRank["一对"], //  AABCDEF 可能是同花或顺子
	"7000": HandRank["单张大牌"], //  ABCDEFG 可能是同花或顺子
}

var GhostHandCountCode = map[string]int { // X+6;无两对牌型
	"0101": HandRank["四条"],
	"2001": HandRank["四条"],
	"0020": HandRank["四条"],
	"1110": HandRank["四条"],

	"0300": HandRank["葫芦"],
	//  非确定
	"3010": HandRank["四条"], //可能为同花顺9、皇家同花顺10 同花5的等级太低了，不需要考虑
	"2200": HandRank["葫芦"], //可能为同花顺、皇家同花顺
	"4100": HandRank["三条"], //可能为顺子、同花、同花顺、皇家同花顺
	"6000": HandRank["一对"], //可能为顺子、同花、同花顺、皇家同花顺
}


// Sort 对手牌进行排序，由于手牌最少长4，最多长7，插入排序性能更好
func Sort(hand string) string {
	runes := []rune(hand)
	l := len(hand)
	for i := 1; i < l; i++ {
		for v := 0; v < i; v++ {
			if FaceRank[string(runes[v])] < FaceRank[string(runes[i])] {
				runes[v], runes[i] = runes[i], runes[v]
			}
		}
	}
	return string(runes)
}

func Max(x, y int) (int, bool) {
	if x > y {
		return 1, false
	} else if x < y {
		return 2, false
	} else {
		return 0, true
	}
}

//GetFaceCount 计算手牌中每种牌出现的次数
func GetFaceCount(hand string) FaceCount {
	var count FaceCount
	for i := 0; i < len(hand); i += 2 {
		count[FaceRank[hand[i:i+1]]]++
	}
	return count
}

//计算出现1次至4次的牌的code
func GetFaceCountCode(count FaceCount) string {
	Sa.Reset()
	var card1, card2, card3, card4 int
	for i, v := range count {
		//  不记录赖子
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
    //绑定切片
	Sa.WriteString(strconv.Itoa(card1))
	Sa.WriteString(strconv.Itoa(card2))
	Sa.WriteString(strconv.Itoa(card3))
	Sa.WriteString(strconv.Itoa(card4))
	return Sa.String()
}

//GetFaceCountMap 返回一个map，该map的key为出现的次数，v是一个[]int，存储了牌面值
func GetFaceCountMap(count FaceCount) map[int][]int {
	countMap := make(map[int][]int, 5) //countMap 等于 键值int 对应 []int{1,2,3} int数组
	countMap[1] = make([]int, 0, 7)
	countMap[2] = make([]int, 0, 3)
	countMap[3] = make([]int, 0, 2)
	countMap[4] = make([]int, 0, 1)
	//  此处i为牌点数，v是牌面出现的次数，i从14开始，不记录鬼牌
	var i, v int
	for i = 14; i >= 2; i-- {
		v = count[i]  //count[i]是记录牌的次数 v的值有 0 1 2 3 4
		if v == 0 {
			continue
		} else {
			countMap[v] = append(countMap[v], i)
		}
	}
	return countMap
}

// IsRoyalFlush isRoyalFlush 在保证同花的前提下判断是否可能为皇家同花顺
func IsRoyalFlush(hand string) bool {
	return strings.Contains(hand, "AKQJT")
	//return strings.Compare(hand, "AKQJT") == 0
}


