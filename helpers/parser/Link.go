package parser
type Link struct {
	Dest string
	Source string
}

func RemoveLink(s []Link, i int) []Link{
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
    return s[:len(s)-1]
}

func AppendUniqueLinks(slice []Link, i []Link) []Link {
	if len(i) > 1 {
	
		for _, val := range i {
			var l  []Link
			l = append(l, val)
			slice = AppendUniqueLinks(slice, l)
		}
		
		return slice
	}
		
	if len(i) == 0 {
		return slice
	}

	for _, ele := range slice {
        if ele == i[0] {
            return slice
        }
    }

	return append(slice, i[0])
}

func RemoveUniqueLinks(slice []Link, i []Link) []Link {
	if len(i) > 1 {
	
		for _, val := range i {
			var l  []Link
			l = append(l, val)
			slice = AppendUniqueLinks(slice, l)
		}
		
		return slice
	}
		
	if len(i) == 0 {
		return slice
	}

	for index, ele := range slice {
        if ele == i[0] {
            return RemoveLink(slice, index)
        }
    }

	return slice
}