package pkg

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
)

func GenerateSeed(v string) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, v); err != nil {
		return 0, fmt.Errorf("failed to write string: %v", err)
	}

	var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))
	return int64(seed), nil
}
