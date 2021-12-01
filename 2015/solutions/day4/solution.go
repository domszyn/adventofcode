package day4

import (
	"crypto/md5"
	"fmt"
)

func FindHashWithPrefix(prefix string) int {
	key := "ckczppom"

	for i := 1; ; i++ {
		input := fmt.Sprintf("%s%d", key, i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input)))

		if hash[:len(prefix)] == prefix {
			return i
		}
	}
}
