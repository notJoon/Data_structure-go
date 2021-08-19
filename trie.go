/*
package main 
import "fmt"
*/

/*
	Tries : Tree structure which is storing string 

	Time Complexity : 
	 	- Insert or Search : O(m) (m : length of the word)
*/

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize] *Node
	isEnd bool 
}

// Trie represent a trie and jas a pointer to the root node 
type Trie struct {
	root *Node
}

// InitTrie will create new Trie 
func InitTrie() *Trie {
	result := &Trie{root : &Node{}}
	return result
}
// Insert will take in a word and add it to the tries
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'  // mapping alphabets to a -> [0], b -> [1], ..., z -> [25]
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true 
}

// Search will take in a word and return true is that word is included in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

/*
func main() {
	test := InitTrie()

	toAdd := []string {
		"aragon",
		"aragorn",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		test.Insert(v)
	}
	fmt.Println(test.Search("ooregon"))
}
*/