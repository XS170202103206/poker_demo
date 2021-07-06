package src

type Matches struct {
	MatchSlice []Match `json:"matches"`
}

type Match struct {
	Hand1  string `json:"alice"`
	Hand2  string `json:"bob"`
	Result int    `json:"result"`
}
