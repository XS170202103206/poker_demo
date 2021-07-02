package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"poker_demo/src"
)

func TestFiveCard(t *testing.T) {
	var matches Matches
	file, err := ioutil.ReadFile("./input/match_result.json")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(file))
	err = json.Unmarshal(file,&matches)
	if err != nil {
		fmt.Println(err)
	}
	c := src.Compare{}
	for _, v := range matches.MatchSlice {
		rst := c.Start(v.Hand1, v.Hand2)
		if rst != v.Result {
			//fmt.Printf("%#v",matches.MatchSlice)
            fmt.Printf("测试json输入的是：%s\t%s\t\n",v.Hand1,v.Hand2)
			fmt.Printf("测试json输入比较之后的是：%d\t%d\t",rst,v.Result)
			fmt.Println("default")
		}
		fmt.Println("------------分割线-----------------")
		/*t.Run(strconv.Itoa(i), func(t *testing.T) {
			rst := c.Start(v.Hand1, v.Hand2)
			if rst != v.Result {
				fmt.Printf("%#v",matches.MatchSlice)
				//t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d", i, v.Hand1, v.Hand2, v.Result, rst)
			}
		})*/

	}
}

