package day18

const AllKeysMask = (1 << 26) - 1

type SearchState struct {
	LastKeys string
	KeyMask  uint32
}

func (s SearchState) HasKey(key rune) bool {
	return (s.KeyMask & (1 << (key - 'a'))) != 0
}

func addKey(mask uint32, key rune) uint32 {
	return mask | (1 << (key - 'a'))
}

func (s SearchState) HasAllKeys() bool {
	return s.KeyMask == AllKeysMask
}
