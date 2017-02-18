package maildir

import "fmt"

func init () {
	m, err = New("_obj/Maildir", true)
	if err != nil {
		fmt.Println(err)
	}
	_ = m
}

var (
	err error
	m *Maildir

)
// Fuzz passes the data to the mock connection
// Data is random input generated by go-fuzz, note that in most cases it is invalid.
// The function must return 1 if the fuzzer should increase priority of the given input during subsequent
// fuzzing (for example, the input is lexically correct and was parsed successfully); -1 if the input must
// not be added to corpus even if gives new coverage; and 0 otherwise
func Fuzz(data []byte) int {
	m.Child(string(data), false)
	return 1
}