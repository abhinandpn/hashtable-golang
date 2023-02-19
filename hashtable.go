package main

import (
	"fmt"
)

// Array size
const arraysize = 7

// Hash Table structure
// Hash table is a Array
type Hashtable struct {
	array [arraysize]*Bucket
}

// Bucket Structure
// Bucket is a LinkedList
type Bucket struct {
	head *Bucketnode
}

// Bucket Node structure
// Bucket node is a Linked list
type Bucketnode struct {
	key  string
	next *Bucketnode
}

// TODO: Hash table
// insert
func (h *Hashtable) Insert(key string) {
	Index := hash(key)
	h.array[Index].insert(key)
}

// Search
// return bool if key is exist or not
func (h *Hashtable) Search(key string) bool {
	Index := hash(key)
	return h.array[Index].Search(key)

}

// Delete
func (h *Hashtable) Delete(key string) {
	Index := hash(key)
	h.array[Index].Delete(key)
}

// TODO: Bucket
// Insert
func (b *Bucket) insert(k string) {
	if !b.Search(k) {
		newnode := &Bucketnode{key: k}
		newnode.next = b.head
		b.head = newnode
	} else {
		fmt.Println(k, "Alredy exist")
	}
}

// Search

func (b *Bucket) Search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete
func (b *Bucket) Delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prvNode := b.head
	for prvNode.next != nil {
		if prvNode.next.key == k {
			// Delete
			prvNode.next = prvNode.next.next
		}
		prvNode = prvNode.next
	}
}

// Hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		// sum += int(v)
		sum = sum + int(v)
	}
	return sum % arraysize
}

// init
// create a bucket slot of hash tables
func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func main() {
	hash := Init()
	list := []string{
		"Abhi",
		"Arun",
		"Mohan",
		"Akshay",
		"Athul",
	}
	for _, v := range list {
		hash.Insert(v)
	}
}
