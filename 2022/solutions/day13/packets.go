package day13

type Packets []Packet

func (a Packets) Len() int           { return len(a) }
func (a Packets) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Packets) Less(i, j int) bool { return a[i].Compare(a[j]) == RightOrder }
