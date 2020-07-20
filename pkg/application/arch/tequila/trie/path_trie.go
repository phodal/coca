// https://github.com/dghubble/trie
//The MIT License (MIT)
//
//Copyright (c) 2014 Dalton Hubble
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.

package trie

import "strings"

type PathTrie struct {
	segmenter StringSegmenter // key segmenter, must not cause heap allocs
	Value     string
	Children  map[string]*PathTrie
}

// PathTrieConfig for building a path trie with different segmenter
type PathTrieConfig struct {
	Segmenter StringSegmenter
}

// NewPathTrie allocates and returns a new *PathTrie.
func NewPathTrie() *PathTrie {
	return &PathTrie{
		segmenter: PathSegmenter,
	}
}

// NewPathTrieWithConfig allocates and returns a new *PathTrie with the given *PathTrieConfig
func NewPathTrieWithConfig(config *PathTrieConfig) *PathTrie {
	segmenter := PathSegmenter
	if config != nil && config.Segmenter != nil {
		segmenter = config.Segmenter
	}

	return &PathTrie{
		segmenter: segmenter,
	}
}

func (trie *PathTrie) Put(key string) {
	node := trie
	for part, i := trie.segmenter(key, 0); part != ""; part, i = trie.segmenter(key, i) {
		child, _ := node.Children[part]
		if child == nil {
			if node.Children == nil {
				node.Children = map[string]*PathTrie{}
			}
			child = NewPathTrie()
			node.Children[part] = child
		}

		child.Value = strings.ReplaceAll(part, "/", "")
		node = child
	}
	// does node have an existing Value?
	//isNewVal := node.Value == nil
	//return isNewVal
}
