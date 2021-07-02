package src

type Compare struct {
	Dealer
	Counter
	Judge
}

type CountRst struct {
	Hand []string
	IsFlush bool //顺子
	IsTongHua bool  //同花
	//IsGhost bool //赖子
	HandType int //牌型
}

func NewCompare() *Compare {
	compare := &Compare{}
	return compare
}

func (c *Compare) Start(hand1, hand2 string) int {
	hands1, hands2 := c.Deal(hand1,hand2) //处理排序
    handRst1, handRst2 := c.QuickCount(hands1), c.QuickCount(hands2) //返回计数结果

	return c.ResultJudge(handRst1, handRst2) //得出结果
}