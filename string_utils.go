package canvas

// Portions taken from https://gobyexample.com/collection-functions

// Unique takes a string slice and returns a new slice with duplicate values removed
func Unique(vs []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range vs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Index the position of the string in the slice if found or -1 if the string is not found
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns true if the slice contains the provided value
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns true if one of the strings in the slice satisfies the predicate f.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all of the strings in the slice satisfy the predicate f.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying the function f to each string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// RemoveAt removes the item at a given index
func RemoveAt(vs []string, i int) []string {
	if i < 0 {
		return vs
	}
	if i > len(vs)-1 {
		return vs
	}
	return append(vs[:i], vs[i+1:]...)
}

// Remove removes the item at a given index
func Remove(vs []string, s string) []string {
	for i, v := range vs {
		if v == s {
			vs = RemoveAt(vs, i)
		}
	}
	return vs
}
