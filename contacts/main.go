package contacts

import "fmt"

type Trie struct {
	count    int
	value    string
	children map[uint8]Trie
}

func makeTrie(value string) Trie {
	return Trie{1, value, make(map[uint8]Trie)}
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
	first_byte := val[0]
	add_value := val[1:]
	root.count++
	cur, ok := root.children[first_byte]
	if ok {
		/* cases at this point:
		cur.value and add_value
		- have no common prefix
		- cur.value is a prefix of add_value
		- add_value is a prefix of cur.value
		- have some common prefix
		*/
		common_prefix, rem_cur_value, rem_add_value := getCommonPrefix(cur.value, add_value)
		if len(common_prefix) == 0 {
			rem_trie := Trie{cur.count, rem_cur_value[1:], cur.children}
			new_trie := makeTrie(add_value[1:])
			cur.children = make(map[uint8]Trie)
			cur.children[rem_cur_value[0]] = rem_trie
			cur.children[add_value[0]] = new_trie
			cur.count = cur.count + 1
		} else if len(rem_cur_value) == 0 {

		} else if len(rem_add_value) == 0 {

		} else {
			rem_trie := makeTrie(rem_cur_value[1:])
			rem_trie.children = cur.children
			cur.children = make(map[uint8]Trie)
			cur.children[rem_cur_value[0]] = rem_trie
			cur.children[rem_add_value[0]] = makeTrie(rem_add_value[1:])
		}
	} else {
		root.children[first_byte] = makeTrie(add_value)
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
