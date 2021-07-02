package main

import "poker_demo/src"

type Matches struct {
	MatchSlice []*Matches `json:"matches"`
	Hand1 string `json:"alice"`
	Hand2 string `json:"bob"`
	Result int `json:"result"`
}
/*type Match struct {
	Hand1 string `json:"alice"`
	Hand2 string `json:"bob"`
	Result int `json:"result"`
}*/

func main(){
	src.NewCompare()
}
