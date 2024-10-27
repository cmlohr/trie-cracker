package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// TrieNode represents each node in the trie 
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd	bool // end of a valid word
}

// trie represents the whole trie structure 
type Trie struct {
	root *TrieNode 
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{childre: make(map[rune]*TrieNode)}}
}

// insert adds a word to the trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists node.children[char]; !exists {
		
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}


//search checks if a word exists in the trie
func(t *Trie) Search (word string) bool {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}


// md5hash generates the md5 hash of a string
func MD5Hash(text string) string {
	hash := md5.sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main () {
	trie := newTrie()

	// placeholder word insertion
	words := []string{"password", "pass", "word", "crack"}
	for _, word := range words {
		trie.Insert(word)
	}

	// search
	fmt.Println("searching for 'pass':", trie.Search("pass")) // true
	fmt.Println("searching for 'hello':", trie.Search("hello")) //false

	//placeholder md5 hashing
	targetHash := MD5Hash("password")
	fmt.Println("MD5 of 'password':", targetHash)
