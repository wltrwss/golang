package calculator

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	test := []struct {
		name string
		a, b int
		ex   int
	}{
		{"test-1", 2, 2, 4},
		{"test-2", 2, 0, 0},
		{"test-3", 0, 2, 0},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			res := Multiply(tt.a, tt.b)
			if tt.ex != res {
				t.Errorf("Ошибка ожидания: %d вместо: %d", res, tt.ex)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	test := []struct {
		name string
		a, b int
		ex   int
	}{
		{"test-1", 2, 2, 1},
		{"test-2", 2, 0, 0},
		{"test-3", 0, 2, 0},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			res := Division(tt.a, tt.b)
			if tt.ex != res {
				t.Errorf("Ошибка ожидания: %d вместо: %d", res, tt.ex)
			}
		})
	}
}
