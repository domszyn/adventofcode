package day20

func GetAnswers() (int, int) {
	return ParseMap(true).findMinDistance(), ParseMap(false).findMinDistanceWithLevels()
}
