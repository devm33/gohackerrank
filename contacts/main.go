package contacts

import "fmt"

type Trie struct {
	count    int
	value    string
	children map[uint8]Trie
}

func makeTrie(value string) Trie {
	return Trie{1, value, make(map[string]Trie)}
}

func getCommonPrefix(a, b string) (string, string, string) {
	// Returns prefix and then the suffixes of the inputs
	var shorter_length = len(a)
	if shorter_length > len(b) {
		shorter_length = len(b)
	}
	var i = 0
	for i < shorter_length && a[i] == b[i] {
		i++
	}
	return a[:i], a[i:], b[i:]
}

func add(root Trie, val string) {
	var first_byte = val[0]
	root.count++
	t, ok := root.children[first_byte]
	if ok {
		new_t_value, rem_t_value, added_value = getCommonPrefix(t.value, val[1:])
		if new_t_value == // TODO this and more cases here
		else if len(new_t_value) > 0 {
			var rem_trie = makeTrie(rem_value[1:])
			rem_trie.children = t.children
			t.children = make(map[string]Trie)
			t.children[rem_value[0]] = rem_trie
			t.children[added_value[0]] = makeTrie(added_value[1:])
		} else {

		}

		// TODO go deeper

	} else {
		root.children[first_byte] = makeTrie(val[1:])
	}
}

func find(root Trie, sub string) int {
	return 0
}

func main() {
	var N int
	fmt.Scan(&N)
	var root = makeTrie("")
	root.count = 0
	var command, arg string
	for i := 1; i <= N; i++ {
		fmt.Scanf("%s %s", &command, &arg)
		if command == "add" {
			add(root, arg)
		} else if command == "find" {
			find(root, arg)
		}
	}
}
