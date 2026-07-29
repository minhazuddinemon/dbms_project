package rest

import "testing"

func Test_extractNameWithoutNickname(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fullName string
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractNameWithoutNickname(tt.fullName)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("extractNameWithoutNickname() = %v, want %v", got, tt.want)
			}
		})
	}
}
