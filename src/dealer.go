package src

import "fmt"


type Dealer struct{
}

func (d *Dealer) Deal(hand1 string, hand2 string)([]string, []string) {
	fmt.Printf("输入的是：%s\t%s\t\n",hand1,hand2)
	hands1 := d.Sort(hand1)
	hands2 := d.Sort(hand2)
	fmt.Printf("排序完的是：%s\t%s\t\n",hands1,hands2)
	l1 := len(hands1)
	l2 := len(hands2)
	fmt.Println(l1,l2)
	val1 :=make([]string,10)
	val2 :=make([]string,10)
    for i, v := range hands1 {
    	val1[i] = string(v)
	}
	for i, v := range hands2 {
		val2[i] = string(v)
	}
	//return []string{hands1},[]string{hands2}
	fmt.Println(val1, val2)
	return val1, val2
}

//将牌从大到小排序
func (*Dealer) Sort(hand string) string {
	l := len(hand)
	fmt.Println("手牌长度：",l)////
	val := []byte(hand)
	fmt.Printf("%#v\n",val)//////

	for i := 2; i < l; i += 2 {
		for v := 0; v < i; v +=2 {
			if FaceRank[string(val[v])] < FaceRank[string(val[i])] {
				val[v], val[i] = val[i], val[v] //交换
				val[v+1], val[i+1] =val[i+1], val[v+1]
			}
		}
	}
	fmt.Printf("%#v\n",string(val))//////
	return string(val)
}
func (d *Dealer) Deal2(hand1 string, hand2 string)([]string, []string) {
	hands1 := d.Sort(hand1)
	hands2 := d.Sort(hand2)

	return []string{hands1},[]string{hands2}

}