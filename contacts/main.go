package main

import "fmt"

type Trie struct {
	count    int
	value    string
	children map[uint8]Trie
}

func makeLeafTrie(value string) Trie {
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
	fmt.Printf("Adding %s to node, value %s\n", val, root.value)
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
		if len(common_prefix) == 0 { // no common prefix, add under empty value
			rem_trie := Trie{cur.count, rem_cur_value[1:], cur.children}
			new_trie := makeLeafTrie(add_value[1:])
			cur.value = ""
			cur.children = make(map[uint8]Trie)
			cur.children[rem_cur_value[0]] = rem_trie
			cur.children[add_value[0]] = new_trie
			cur.count += 1
		} else if len(rem_cur_value) == 0 { // add under cur
			add(cur, rem_add_value)
		} else if len(rem_add_value) == 0 { // add rem cur under
			// TODO last case
		} else { // some common prefix, need to split
			rem_trie := Trie{cur.count, rem_cur_value[1:], cur.children}
			cur.value = common_prefix
			cur.children = make(map[uint8]Trie)
			cur.children[rem_cur_value[0]] = rem_trie
			cur.children[rem_add_value[0]] = makeLeafTrie(rem_add_value[1:])
			cur.count += 1
		}
	} else {
		fmt.Printf("First byte %s not found in node (value %s)\n", string(first_byte), root.value)
		root.children[first_byte] = makeLeafTrie(add_value)
	}
}

func findCount(root Trie, sub string) int {
	fmt.Printf("Finding count for %s\n", sub)
	cur := root
	var ok bool
	for i := 0; i < len(sub); i++ {
		// TODO need to match on cur.value here before looking at children
		cur, ok = root.children[sub[i]]
		if !ok {
			return 0
		}
	}
	return cur.count
}

func printTrie(val uint8, root Trie, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("%s -> value=%s count=%d \n", string(val), root.value, root.count)
	for v, c := range root.children {
		printTrie(v, c, indent+1)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	var root = makeLeafTrie("")
	root.count = 0
	var command, arg string
	for i := 1; i <= N; i++ {
		fmt.Scanf("%s %s", &command, &arg)
		if command == "add" {
			add(root, arg)
			printTrie(0, root, 0)
		} else if command == "find" {
			fmt.Println(findCount(root, arg))
		}
	}
}
