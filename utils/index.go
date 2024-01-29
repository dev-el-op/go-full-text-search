package utils

type Index map[string][]int

func (index Index) Add(documents []Document) {
	for _, document := range documents {
		for _, token := range analyze(document.Text) {
			ids := index[token]

			if ids != nil && ids[len(ids)-1] == document.ID {
				continue
			}

			index[token] = append(ids, document.ID)
		}
	}
}

func Intersection(a []int, b []int) []int {
	maxLength := len(a)

	if len(b) > maxLength {
		maxLength = len(b)
	}

	intersection := make([]int, 0, maxLength)

	var i, j int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			intersection = append(intersection, a[i])
			i++
			j++
		}
	}

	return intersection

}

func (index Index) Search(query string) []int {
	var results []int

	for _, token := range analyze(query) {

		if ids, ok := index[token]; ok {
			if results == nil {
				results = ids
			} else {
				results = Intersection(results, ids)
			}
		} else {
			return nil
		}

	}

	return results
}
