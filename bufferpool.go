// Package bufferpool is a wrapper around sync.Pool for bytes.Buffer objects.
package bufferpool

import (
	"bytes"
	"sync"

	"github.com/chronos-tachyon/assert"
)

var gPool = sync.Pool{
	New: func() interface{} {
		buf := new(bytes.Buffer)
		buf.Grow(256)
		return buf
	},
}

// Get returns an empty bytes.Buffer with room for at least 256 bytes.
func Get() *bytes.Buffer {
	return gPool.Get().(*bytes.Buffer)
}

// Put returns the given bytes.Buffer to the pool.
func Put(buf *bytes.Buffer) {
	assert.NotNil(&buf)
	buf.Reset()
	buf.Grow(256)
	gPool.Put(buf)
}
