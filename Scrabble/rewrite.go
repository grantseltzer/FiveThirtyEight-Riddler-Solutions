package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// ByLength interface for sort
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	//  Pull in all strings from enable file (from STDIO)
	//  and populate slice with them
	scanner := bufio.NewScanner(os.Stdin)
	allStrings := []string{}
	for scanner.Scan() {
		allStrings = append(allStrings, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// Sort and reverse allStrings
	sort.Sort(ByLength(allStrings))
	last := len(allStrings) - 1
	for i := 0; i < len(allStrings)/2; i++ {
		allStrings[i], allStrings[last-i] = allStrings[last-i], allStrings[i]
	}

	// Calculate delimiting indexes of words of different lengths
	startingIndexes := []int{0}
	previousLength := 28
	lengths := []int{len(allStrings[0])}
	for i, currentWord := range allStrings {
		if len(currentWord) < previousLength {
			startingIndexes = append(startingIndexes, i)
			lengths = append(lengths, len(currentWord))
			previousLength = len(currentWord)

		}
	}

	fmt.Println(lengths)

	// Seperate words into slices by length
	sliceOfSlices := [][]string{}
	for x := range startingIndexes {
		if x == len(startingIndexes)-1 {
			sliceOfSlices = append(sliceOfSlices, allStrings[startingIndexes[x]:])
		} else {
			sliceOfSlices = append(sliceOfSlices, allStrings[startingIndexes[x]:startingIndexes[x+1]])
		}
	}

	var x *avlNode
	sliceOfReferenceTrees := []*avlNode{}
	for _, indivSlice := range sliceOfSlices {
		x = avlNodePtr()
		for _, valueInSlice := range indivSlice {
			x = insert(x, valueInSlice)
		}
		sliceOfReferenceTrees = append(sliceOfReferenceTrees, x)
	}

	inOrder(sliceOfReferenceTrees[0], func(node *avlNode) {
		fmt.Println(node.Key, node.Height)
	})

}

func avlNodePtr() *avlNode { return &avlNode{} }

// func checkOnInsert(root *avlNode, key string, sliceOfRefs []*avlNode) *avlNode {
// 	index := 28 - len(key)
// }

func containedIn(word string, strings []string) bool {
	for _, t := range strings {
		if word == t {
			return true
		}
	}
	return false
}

func splitFromFrontAndBack(word string) (string, string) {
	return word[0 : len(word)-1], word[1:len(word)]
}
