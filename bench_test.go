package main

import "testing"

func BenchmarkEncodeDecodeWithBinary(b *testing.B) {
	point := &DataPoint{Value: 0.1, Timestamp: 1}
	for i := 1; i < b.N; i++ {
		EncodeDecodeWithBinary(point)
	}
}

func BenchmarkEncodeDecodeWithGob(b *testing.B) {
	point := &DataPoint{Value: 0.1, Timestamp: 1}
	for i := 1; i < b.N; i++ {
		EncodeDecodeWithGob(point)
	}
}
