package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"negative size", -5, 0},
		{"zero size", 0, 0},
		{"small size", 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			if len(result) != tt.want {
				t.Errorf("expected length %d, got %d", tt.want, len(result))
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"positive numbers", []int{1, 3, 7, 2, 5}, 7},
		{"negative numbers", []int{-10, -3, -50, -1}, -1},
		{"mixed numbers", []int{-5, 0, 10, -2, 3}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d; want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{"empty slice", []int{}},
		{"less than chunks", []int{1, 5, 3}},
		{"exact chunks", []int{1, 5, 3, 9, 2, 8, 4, 7}},
		{"more than chunks", []int{1, 9, 3, 4, 15, 6, 2, 8, 10, 7, 11}},
		{"negative numbers", []int{-10, -5, -2, -30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := maximum(tt.data)
			got := maxChunks(tt.data)

			if got != want {
				t.Errorf("maxChunks(%v) = %d; want %d", tt.data, got, want)
			}
		})
	}
}
