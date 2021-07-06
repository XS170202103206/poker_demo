package sevenghost

import (
	"encoding/json"
	"fmt"
	. "poker_demo/src"

	"io/ioutil"
	"testing"
	"time"
)

func TestSevenHand_Spend_Time(t *testing.T) {
	var matches Matches
	file, err := ioutil.ReadFile("../../input/seven_cards_with_ghost.result.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(file, &matches)
	if err != nil {
		fmt.Println(err)
	}
	c := Counter{}
	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		c.Start(v.Hand1, v.Hand2)
	}
	endTime := time.Since(startTime)
	fmt.Println("Spend:", endTime)
}

func TestSevenHand_Result(t *testing.T) {
	var matches Matches
	file, err := ioutil.ReadFile("../../input/seven_cards_with_ghost.result.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(file, &matches)
	if err != nil {
		fmt.Println(err)
	}
	c := Counter{}
	errCount := 0
	for _, v := range matches.MatchSlice {
		rst := c.Start(v.Hand1, v.Hand2)
		if rst != v.Result {
			errCount++
			fmt.Printf("测试json输入比较之后的是：%d\t%d\t\n",rst,v.Result)
			//t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",i, v.Hand1, v.Hand2, v.Result, rst)
		}
		fmt.Printf("测试json输入比较之后的是：%d\t%d\t\n",rst,v.Result)
	}
	fmt.Println("all err :", errCount)
}