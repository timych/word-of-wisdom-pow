package wisdom

import "testing"

func TestGetWordOfWisdom(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Run("", func(t *testing.T) {
			if got := GetWordOfWisdom(); len(got) < 10 {
				t.Errorf("GetWordOfWisdom() = %v, len() = %v, want > 10", got, len(got))
			}
		})
	}
}
