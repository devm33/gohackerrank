package contacts

import "fmt"

type Trie struct {
	count    int
	children map[string]Trie
}

func add(root Trie, val string) {

}

func find(root Trie, sub string) {

}

func main() {
	var N int
	fmt.Scan(&N)
	var root Trie
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
