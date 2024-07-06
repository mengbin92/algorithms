package sort

import (
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "test1",
			arr:  []int{1, 3, 2, 4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.arr, tt.want)
			}
		})
	}
}

func TestString(t *testing.T){
	tests := []struct {
		name string
		arr  []string
		want []string
	}{
		{
			name: "test1",
			arr:  []string{"a", "b", "c", "d", "e", "f"},
			want: []string{"a", "b", "c", "d", "e", "f"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.arr, tt.want)
			}
		})
	}
}

func TestFloat(t *testing.T){
	tests := []struct {
		name string
		arr  []float64
		want []float64
	}{
		{
			name: "test1",
			arr:  []float64{1.1, 3.3, 2.2, 4.4, 5.5, 6.6},
			want: []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", tt.arr, tt.want)
			}
		})
	}
}