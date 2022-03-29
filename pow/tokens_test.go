package pow

import (
	"testing"
)

func TestActiveTokens_Append(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		at   *ActiveTokens
		args args
	}{
		{
			&ActiveTokens{},
			args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077"},
		},
		{
			&ActiveTokens{tokens: []string{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077"}},
			args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077"},
		},
		{
			&ActiveTokens{},
			args{""},
		},
	}
	for _, tt := range tests {
		tt.at.Append(tt.args.t)
	}
}

func TestActiveTokens_Delete(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		at   *ActiveTokens
		args args
		want bool
	}{
		{
			"",
			&ActiveTokens{tokens: []string{""}},
			args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077"},
			false,
		},
		{
			"",
			&ActiveTokens{tokens: []string{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077"}},
			args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077"},
			true,
		},
		{
			"",
			&ActiveTokens{},
			args{""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.at.Delete(tt.args.t); got != tt.want {
				t.Errorf("ActiveTokens.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
