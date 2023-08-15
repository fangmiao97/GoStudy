package others

import "testing"

func Test_makefun(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test_makefun",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makefun()
		})
	}
}
