package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			name: "отрицательный размер",
			size: -5,
			want: 0,
		},
		{
			name: "нулевой размер",
			size: 0,
			want: 0,
		},
		{
			name: "один элемент",
			size: 1,
			want: 1,
		},
		{
			name: "несколько элементов",
			size: 100,
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if len(data) != tt.want {
				t.Fatalf(
					"длина слайса = %d, ожидалось %d",
					len(data),
					tt.want,
				)
			}

			for i, value := range data {
				if value < 1 || value > 1_000_000 {
					t.Errorf(
						"data[%d] = %d: число вне диапазона [1, 1000000]",
						i,
						value,
					)
				}
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
		{
			name: "nil",
			data: nil,
			want: 0,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "максимум в начале",
			data: []int{99, 4, 7, 2},
			want: 99,
		},
		{
			name: "максимум в середине",
			data: []int{4, 99, 7, 2},
			want: 99,
		},
		{
			name: "максимум в конце",
			data: []int{4, 7, 2, 99},
			want: 99,
		},
		{
			name: "одинаковые элементы",
			data: []int{5, 5, 5},
			want: 5,
		},
		{
			name: "повторяющийся максимум",
			data: []int{9, 1, 9, 3},
			want: 9,
		},
		{
			name: "нули",
			data: []int{0, 0, 0},
			want: 0,
		},
		{
			name: "отрицательные элементы",
			data: []int{-10, -3, -7},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d",
					tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil",
			data: nil,
			want: 0,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "меньше восьми элементов",
			data: []int{3, 9, 1},
			want: 9,
		},
		{
			name: "ровно восемь элементов",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 8,
		},
		{
			name: "длина кратна восьми",
			data: []int{
				1, 2, 3, 4, 5, 6, 7, 8,
				99, 10, 11, 12, 13, 14, 15, 16,
			},
			want: 99,
		},
		{
			name: "максимум в остатке",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100},
			want: 100,
		},
		{
			name: "одинаковые элементы",
			data: []int{5, 5, 5, 5, 5, 5, 5, 5, 5},
			want: 5,
		},
		{
			name: "отрицательные элементы в коротком слайсе",
			data: []int{-10, -3, -7},
			want: -3,
		},
		{
			name: "отрицательные элементы во всех частях",
			data: []int{-8, -7, -6, -5, -4, -3, -2, -1},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)

			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d",
					tt.data, got, tt.want)
			}
		})
	}
}