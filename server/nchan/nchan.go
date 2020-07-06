package nchan

// WriteString 可以无堵塞向chan写入数据
func WriteString(c chan string, d string) bool {
	select {
	case c <- d:
		return true
	default:
		return false
	}
}

//CoverString 可以覆盖之前的数据,写入数据
func CoverString(c chan string, d string) {
	if WriteString(c, d) {
		return
	}
	<-c
	WriteString(c, d)
}

//ReadString 无堵塞读取数据
func ReadString(c chan string) (string, bool) {
	var result string
	select {
	case result = <-c:
		return result, true
	default:
		return "", false
	}
}

//AllString 读取所有数据
func AllString(c chan string) []string {
	var results []string = []string{}
	for {
		if result, ok := ReadString(c); ok {
			results = append(results, result)
		} else {
			break
		}
	}
	return results
}
