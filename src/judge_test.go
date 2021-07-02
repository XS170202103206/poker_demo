package src

import "testing"

func TestJudge_EqualJudge(t *testing.T) {
	type args struct {
		hands1   []string
		hands2   []string
		handRank int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Judge{}
			if got := j.EqualJudge(tt.args.hands1, tt.args.hands2, tt.args.handRank); got != tt.want {
				t.Errorf("EqualJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJudge_GetBestHand(t *testing.T) {
	type args struct {
		hands    []string
		handRank int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Judge{}
			got, got1 := j.GetBestHand(tt.args.hands, tt.args.handRank)
			if got != tt.want {
				t.Errorf("GetBestHand() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetBestHand() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestJudge_QuickJudge(t *testing.T) {
	type args struct {
		handType1 int
		handType2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Judge{}
			if got := j.QuickJudge(tt.args.handType1, tt.args.handType2); got != tt.want {
				t.Errorf("QuickJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJudge_ResultJudge(t *testing.T) {
	type args struct {
		countRst1 *CountRst
		countRst2 *CountRst
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Judge{}
			if got := j.ResultJudge(tt.args.countRst1, tt.args.countRst2); got != tt.want {
				t.Errorf("ResultJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whoIsMax(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := whoIsMax(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("whoIsMax() = %v, want %v", got, tt.want)
			}
		})
	}
}