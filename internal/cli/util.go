package cli

func contains(list []string, name string) bool {
	for _, v := range list {
		if v == name {
			return true
		}
	}
	return false
}

func index(list []string, name string) int {
	for i, v := range list {
		if v == name {
			return i
		}
	}
	return -1
}

func remove(list []string, name string) []string {
	index := index(list, name)
	if index < 0 {
		return list
	}
	list[index] = list[len(list)-1]
	list[len(list)-1] = ""
	return list[:len(list)-1]
}
