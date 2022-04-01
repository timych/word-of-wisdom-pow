package pow

import (
	"context"
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewChallenge(t *testing.T) {
	if got := NewChallenge(); len(got) != 16 {
		t.Errorf("NewChallenge() = %x, len() = %v, want %v", got, len(got), 16)
	}
}

func TestCompute(t *testing.T) {
	type args struct {
		ctx        context.Context
		seed       []byte
		complexity uint8
	}
	tests := []struct {
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args{context.Background(), []byte{0xa5, 0x6f, 0xbf, 0xcb, 0xc6, 0xa8, 0x4a, 0x3b, 0x8d, 0x2e, 0x12, 0xe8, 0xd8, 0x2a, 0xf0, 0x77}, 20},
			[]byte{184, 246, 11, 0, 0, 0, 0, 0},
			false,
		},
		{
			args{context.Background(), []byte{}, 20},
			[]byte{26, 254, 33, 0, 0, 0, 0, 0},
			false,
		},
		{
			args{context.Background(), []byte{}, 0},
			[]byte{0, 0, 0, 0, 0, 0, 0, 0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := Compute(tt.args.ctx, tt.args.seed, tt.args.complexity)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompute_Indirect(t *testing.T) {
	type args struct {
		ctx        context.Context
		seed       string
		complexity uint8
	}
	tests := []struct {
		args    args
		want    uint64
		wantErr bool
	}{
		{args{context.Background(), "a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 20}, 784056, false},
		{args{context.Background(), "ba441627-db1a-483c-9d14-6b754ce03d2b", 20}, 1004822, false},
		{args{context.Background(), "cc1a6d60-e9bf-482a-8b9c-c57e04a1debd", 20}, 689067, false},
		{args{context.Background(), "cc1a6d60-e9bf-482a-8b9c-c57e04a1debd", 19}, 537309, false},
		{args{context.Background(), "cc1a6d60-e9bf-482a-8b9c-c57e04a1debd", 18}, 24872, false},
		{args{context.Background(), "cc1a6d60-e9bf-482a-8b9c-c57e04a1debd", 5}, 41, false},
		{args{context.Background(), "cc1a6d60-e9bf-482a-8b9c-c57e04a1debd", 1}, 0, false},
		{args{context.Background(), "cc1a6d60-e9bf-482a-8b9c-c57e04a1debd", 0}, 0, false},
		{args{context.Background(), "", 5}, 6, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			wb := make([]byte, 8)
			binary.LittleEndian.PutUint64(wb, tt.want)
			u, _ := uuid.Parse(tt.args.seed)
			ub, _ := u.MarshalBinary()
			got, err := Compute(tt.args.ctx, ub, tt.args.complexity)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, wb) {
				t.Errorf("Compute() = %v, want %v", binary.LittleEndian.Uint64(got), binary.LittleEndian.Uint64(wb))
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		seed       string
		complexity uint8
		solution   uint64
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 20, 784056}, true},
		{args{"ba441627-db1a-483c-9d14-6b754ce03d2b", 20, 1004822}, true},
		{args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 20, 1004822}, false},
		{args{"a56fbfcb-c6a8-4a3b-8d2e-12e8d82af077", 0, 0}, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			sb := make([]byte, 8)
			binary.LittleEndian.PutUint64(sb, tt.args.solution)
			u, _ := uuid.Parse(tt.args.seed)
			ub, _ := u.MarshalBinary()
			if got := Verify(ub, tt.args.complexity, sb); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shaHash(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		args args
		want string
	}{
		{args{[]byte("123")}, "40bd001563085fc35165329ea1ff5c5ecbdbbeef"},
		{args{[]byte("hello")}, "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{args{[]byte("!!!")}, "9a7b006d203b362c8cef6da001685678fc1d463a"},
		{args{[]byte("40bd001563085fc35165329ea1ff5c5ecbdbbeef")}, "9adcb29710e807607b683f62e555c22dc5659713"},
		{args{[]byte("")}, "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{args{[]byte{0xa5, 0x6f, 0xbf, 0xcb, 0xc6, 0xa8, 0x4a, 0x3b, 0x8d, 0x2e, 0x12, 0xe8, 0xd8, 0x2a, 0xf0, 0x77}}, "4b8ef01f9ed3cf1a4db2c09fcfabd2daf2eae10c"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := shaHash(tt.args.in); fmt.Sprintf("%x", got) != tt.want {
				t.Errorf("shaHash() = %x, want %v", got, tt.want)
			}
		})
	}
}
