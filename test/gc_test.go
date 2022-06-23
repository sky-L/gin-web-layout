package test

import (
	"os"
	"runtime/trace"
	"testing"
)

func TestGc(t *testing.T) {
	f, _ := os.Create("trace.out")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()

	for n := 1; n < 10000; n++ {
		allocate()
	}
}

func allocate() {
	_ = make([]byte, 1<<20)
}
