package pokecache

import (
	"time"
)

type CacheEntry struct {
	val     []byte
	created time.Time
}
