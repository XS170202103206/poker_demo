package fivehand

import (
	"encoding/json"
	"fmt"
	. "poker_demo/src"

	"io/ioutil"
	"testing"
	"time"
)

func TestFiveHand_Spend_Time(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("../../input/match.json")
	_ = json.Unmarshal(file, &matches)

	c := Counter{}
	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		c.Start(v.Hand1, v.Hand2)
	}
	endTime := time.Since(startTime)
	fmt.Println("Spend:", endTime)
}

func TestFiveHand_Result(t *testing.T) {
	var matches Matches
	file, err := ioutil.ReadFile("../../input/match_result.json")
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
			//fmt.Printf("测试json输入比较之后的是：%d\t%d\t",rst,v.Result)
			//t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",i, v.Hand1, v.Hand2, v.Result, rst)
		}
	}
	fmt.Println("all err :", errCount)
}

func TestFiveHand_ByHand(t *testing.T) {
	c := Counter{}
	hand1, hand2 := "6s5h4c3s2c", "As2h3s4c5s"
	result := 1
	rst := c.Start(hand1, hand2)
	if rst != result {
		t.Fatalf("h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
			hand1, hand2, result, rst)
	}
}