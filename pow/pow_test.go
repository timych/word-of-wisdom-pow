package pow

import (
	"context"
	"testing"
)

func TestGenerateChallengeToken(t *testing.T) {
	if got := GenerateChallengeToken(); len(got) != 36 {
		t.Errorf("GenerateChallengeToken() = %v, len() = %v, want %v", got, len(got), 36)
	}
}

func TestCompute(t *testing.T) {
	type args struct {
		ctx        context.Context
		token      string
		complexity uint32
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"", args{context.Background(), "a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 5}, 645555, false},
		{"", args{context.Background(), "ba441627-db1a-483c-9d14-6b754ce03d2b", 3}, 644, false},
		{"", args{context.Background(), "cc1a6d60-e9bf-482a-8b9c-c57e04a1debd", 4}, 11042, false},
		{"", args{context.Background(), "cc1a6d60-e9bf-482a-8b9c-c57e04a1debd", 5}, 0, true},
		{"", args{context.Background(), "e35477ce-049c-4b1b-b94a-235d013c3ced", 0}, 0, false},
		{"", args{context.Background(), "", 5}, 1862378, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Compute(tt.args.ctx, tt.args.token, tt.args.complexity)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		token      string
		complexity uint32
		solution   int64
	}
	tests := []struct {
		name         string
		activeTokens []string
		args         args
		want         bool
	}{
		{
			"",
			[]string{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077"},
			args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 5, 645555},
			true,
		},
		{
			"",
			[]string{},
			args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 5, 645555},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			activeTokens = ActiveTokens{tokens: tt.activeTokens}
			if got := Verify(tt.args.token, tt.args.complexity, tt.args.solution); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkSolution(t *testing.T) {
	type args struct {
		token      string
		complexity uint32
		solution   int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 5, 645555}, true},
		{"", args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 5, 645554}, false},
		{"", args{"ba441627-db1a-483c-9d14-6b754ce03d2b", 3, 644}, true},
		{"", args{"", 5, 1862378}, true},
		{"", args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 0, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSolution(tt.args.token, tt.args.complexity, tt.args.solution); got != tt.want {
				t.Errorf("checkSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shaHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"123"}, "40bd001563085fc35165329ea1ff5c5ecbdbbeef"},
		{"", args{"hello"}, "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{"", args{"!!!"}, "9a7b006d203b362c8cef6da001685678fc1d463a"},
		{"", args{"40bd001563085fc35165329ea1ff5c5ecbdbbeef"}, "9adcb29710e807607b683f62e555c22dc5659713"},
		{"", args{""}, "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shaHash(tt.args.s); got != tt.want {
				t.Errorf("shaHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
