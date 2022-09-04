package main

import "testing"

// var args [2]string
// args[0] = "Golang"
// args[1] = "Node"

func BenchmarkEchoCli2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoCli2()
	}
}
