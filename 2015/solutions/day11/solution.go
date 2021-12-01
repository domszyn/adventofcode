package day11

func GetNewPassword(pwd string) string {
	var iteration int
	for {
		pwd = increment(pwd)
		iteration++
		if isValid(pwd) {
			break
		}
	}

	return pwd
}
