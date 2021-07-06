package src

//import (
//	"fmt"
//)
//
//type Judge struct{
//}
//
//func (j *Judge) ResultJudge(countRst1, countRst2 *CountRst) int {
//	rank1, rank2 := countRst1.HandType, countRst2.HandType
//	fmt.Println("ResultJudge的牌型大小：", rank1, rank2)
//	rst := j.QuickJudge(rank1, rank2)
//
//	//  rst == 0 表示牌型相同，从头到尾比较，得出最大的那一方
//	if rst == 0 && rank1 != HandRank["皇家同花顺"] {
//		fmt.Println("测试是否是rst等于0出了问题！")
//		rst = j.EqualJudge(countRst1.Hand, countRst2.Hand, rank1)
//		fmt.Printf("排查问题是否是EqualJudge的问题:%d\thand1:%s\thand2:%s\trank1:%d\n",rst,countRst1.Hand, countRst2.Hand, rank1)
//	}
//
//	return rst
//}
////
//func (j *Judge) QuickJudge(handType1, handType2 int) int {
//	fmt.Printf("judge内的handType:%d\t%d\n",handType1,handType2)
//	if handType1 > handType2 {
//		return 1
//	} else if handType1 < handType2 {
//		return 2
//	} else {
//		return 0
//	}
//}
//
////平牌，从大到小 逐个比较他们的大小
//func (j *Judge) EqualJudge(hands1,hands2 []string,handRank int) int {
//	fmt.Printf("平牌Judge检查手牌hands1是否已经排序%s：",hands1)
//    rst := 0
//	if handRank != HandRank["皇家同花顺"] {
//		hand1, _ := j.GetBestHand(hands1, handRank)
//		hand2, _ := j.GetBestHand(hands2, handRank)
//		if hand1 != hand2 {
//			_, rst = j.GetBestHand([]string{hand1, hand2}, handRank)
//		}
//	}
//    return rst
//}
//
//// GetBestHand 返回牌点数最大的牌型，返回在切片中的下标，相同时返回0,不比较皇家同花顺
//func (j *Judge) GetBestHand(hands []string, handRank int) (string, int) {
//	fmt.Println("GetBestHand的hands：",hands)
//	var max = hands[0]
//	//var cur string
//	fmt.Println(hands[0])
//	dealer := Dealer{}
//	for i := 1; i < len(hands); i++ {
//		//cur = hands[i]
//		cur := dealer.Sort(hands[i])
//		switch handRank {
//		//  按顺序比较牌点数
//		case HandRank["顺子"]:
//			fallthrough
//		case HandRank["同花顺"]:
//			fallthrough
//		case HandRank["同花"]:
//			fallthrough
//		case HandRank["高牌"]:
//			rst := whoIsMax(max, cur)
//			if rst == 2 {
//				max = cur
//			}
//		case HandRank["一对"]:
//		case HandRank["两对"]:
//		case HandRank["三条"]:
//		case HandRank["葫芦"]:
//		case HandRank["四条"]:
//		}
//	}
//
//	return max, 0
//}
//
//func whoIsMax(s1, s2 string) int {
//	if len(s1) == 0 {
//		return 0
//	}
//
//	v1, v2 := FaceRank[s1[0:1]], FaceRank[s2[0:1]]
//	if v1 > v2 {
//		return 1
//	} else if v1 < v2 {
//		return 2
//	} else {
//		return whoIsMax(s1[2:], s2[2:])
//	}
//}