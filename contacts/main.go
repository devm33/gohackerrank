package contacts

import "fmt"

type Trie struct {
	count    int
	value    string
	children map[string]Trie
}

func makeTrie(value string) Trie {
	return Trie{1, value, make(map[string]Trie)}
}

func add(root Trie, val string) {
	var first_letter = val[0:1]
	root.count++
	t, ok := root.children[first_letter]
	if ok {
		// TODO go deeper
	} else {
		root.children[first_letter] = makeTrie(val[1:])
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
