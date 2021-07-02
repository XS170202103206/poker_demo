package src

import (
	"fmt"
	"strings"
)

type Counter struct{
}

func (c *Counter) QuickCount(hands []string) *CountRst {
	var newHands []string
	//rank := 0
	maxType := 0
	//字符串数组 转换成 字符串！
	var result string
	for _,v := range hands {
		result += v
	}

	l := len(hands)
	l2 := len(result)
	fmt.Println("hhhhhh:",l,hands)
	fmt.Println("xxxxxx:",l2,result)
	//[]string{hands1},[]string{hands2}

	hands3 := []string{result}

	l3 := len(hands3)
	fmt.Println("wwwwww:",l3,hands3)
	for _, v := range hands3 {
		rank := c.GetHandType(v)
		//fmt.Println("牌型排名：",rank)
		if rank > maxType {
			maxType = rank
			//newHands = newHands[0:0]
			newHands = append(newHands, v)
			fmt.Println("QuickCount的newHands:",newHands)
		} else if rank == maxType {
			newHands = append(newHands, v)
		}
	}
	return &CountRst{Hand: newHands,  HandType: maxType}
}

func (c *Counter) GetHandType(hand string) int {
	code := c.GetHandFaceCountInfo(c.GetHandFaceCount(hand))
	tp := FiveHandCount[code]

	if tp == 1 {
		hasFlush := c.HasFlush(hand) //判断是否是顺子
		isTongHua := c.IsTongHua(hand)//判断是否是同花
		if hasFlush {
			if isTongHua {
				if c.IsRoyalFlush(hand) {
					tp = HandRank["皇家同花顺"]
				} else {
					tp = HandRank["同花顺"]
				}
			} else {
				tp = HandRank["顺子"]
			}
		} else if isTongHua {
			tp = HandRank["同花"]
		} else {
			tp = HandRank["高牌"]
		}
	}
	fmt.Println("GetHandType得出牌型：",tp)
    return tp
}

//计算手牌中每种牌出现的次数
func (c *Counter) GetHandFaceCount(hand string) [15]int {
	count := [15]int{}
	fmt.Println("GetHandFaceCount的hand：",hand)////******
	for i := 0; i<len(hand); i += 2 {
		fmt.Println(string(hand[i]))
		count[FaceRank[string(hand[i])]] += 1
	}
	return count
}

//根据统计的数据,匹配对应的牌型(四条/葫芦/三条/两对/一对/单张大牌)
func (c *Counter) GetHandFaceCountInfo(count [15]int) string {
	var card1, card2, card3, card4 int
	for _, v := range count {
		switch v {
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
	return fmt.Sprintf("%d%d%d%d",card1, card2, card3, card4)
	// card1 card2 card3 card4 : 牌型
	// 0001 : 四条 权重8
	// 1010 : 三条 权重8
	// 0200 : 两对 权重7
	// 2100 : 一对 权重4
	// 4000 : 单牌 权重1
}

func (*Counter) IsTongHua(hand string) bool {
	rst := true
	key := hand[1]
	for i := 3; i < len(hand); i +=2 {
		if hand[i] != key {
			rst = false
			break
		}
	}
	return rst
}

func (c *Counter) HasFlush(hand string) bool {
	var rst = true
	last := FaceRank[string(hand[0])]
	fmt.Println("打印last：",last)
	for i := 2; i < len(hand); i += 2 {
		val := FaceRank[string(hand[i])]
		if last-1 != val {
			rst = false
			break
		}
		last = val
	}

	//  A2345在德州扑克里是最小顺子，由于事先对 手牌进行了排序，A总是出现在第一位，所以特殊判断
	if strings.Contains(hand, "A") &&
		strings.Contains(hand, "4") &&
		strings.Contains(hand, "3") &&
		strings.Contains(hand, "2") &&
		strings.Contains(hand, "5") {
			rst = true
		}
	return rst
}

//皇家同花顺
func (*Counter) IsRoyalFlush(hand string) bool {
	var rst bool
	if hand[0] == 'A' &&
		hand[2] == 'K' &&
		hand[4] == 'Q' &&
		hand[6] == 'J' &&
		hand[8] == 'T' {
		rst = true
	}

	return rst
}