package nibblesort

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {

	x := uint64(1)

	for i := 0; i < 10000; i++ {
		x = nextrng(x)

		y := Sort(x)

		s := suint64(x)
		sort.Sort(&s)

		if y != uint64(s) {
			t.Errorf("failed to sort x=%x; got=%x, want %x", x, y, uint64(s))
		}
	}
}

type suint64 uint64

func (s suint64) Len() int           { return 16 }
func (s suint64) Less(i, j int) bool { return (s>>uint(i*4))&0x0f < (s>>uint(j*4))&0x0f }

func (s *suint64) Swap(i, j int) {
	*s = *s & ^(0x0f<<uint(i*4)) & ^(0x0f<<uint(j*4)) | (((*s >> uint(i*4)) & 0x0f) << uint(j*4)) | (((*s >> uint(j*4)) & 0x0f) << uint(i*4))
}

func BenchmarkSort(b *testing.B) {

	x := uint64(0x0ddc0ffeebadf00d)

	for i := 0; i < b.N; i++ {
		Sort(x)
		x = nextrng(x)
	}

}

func BenchmarkQuick(b *testing.B) {

	x := uint64(0x0ddc0ffeebadf00d)

	for i := 0; i < b.N; i++ {
		s := suint64(x)
		sort.Sort(&s)
		x = nextrng(x)
	}
}

func nextrng(x uint64) uint64 {
	x ^= x >> 12 // a
	x ^= x << 25 // b
	x ^= x >> 27 // c
	return x * 2685821657736338717

}
