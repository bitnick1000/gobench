package md5

import (
	"crypto/md5"
	// "encoding/hex"
	"io"
	"testing"
)

func Benchmark_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := md5.New()
		io.WriteString(h, "message digest")
		h.Sum(nil)
		// io.WriteString(h, randString())
		// id := hex.EncodeToString(h.Sum(nil))
	}
}
