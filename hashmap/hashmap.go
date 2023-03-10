package main

import "fmt"

const HASH_TABLE_ROWS = 100

type HashTable struct {
	array [HASH_TABLE_ROWS]*Bucket
}

type Bucket struct {
	head *BucketNode
}

type BucketNode struct {
	key   string
	value interface{}
	next  *BucketNode
}

func Init() *HashTable {
	hashTable := &HashTable{}
	for i := range hashTable.array {
		hashTable.array[i] = &Bucket{}
	}
	return hashTable
}

func (hashTable *HashTable) Insert(key interface{}, value interface{}) {
	keyStr := fmt.Sprintf("%v", key)
	index := hash(keyStr)
	hashTable.array[index].insert(keyStr, value)
}

func (hashTable *HashTable) Search(key interface{}) *BucketNode {
	keyStr := fmt.Sprintf("%v", key)
	index := hash(keyStr)
	return hashTable.array[index].search(keyStr)
}

func (hashTable *HashTable) Delete(key interface{}) {
	keyStr := fmt.Sprintf("%v", key)
	index := hash(keyStr)
	hashTable.array[index].delete(keyStr)
}

func (bucket *Bucket) insert(key string, value interface{}) {
	newNode := &BucketNode{}
	newNode.key = key
	newNode.value = value
	newNode.next = bucket.head
	bucket.head = newNode
}

func (bucket *Bucket) search(key string) *BucketNode {
	currentNode := bucket.head
	for currentNode != nil {
		if currentNode.key == key {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

func (bucket *Bucket) delete(key string) {
	previousNode := bucket.head
	currentNode := previousNode.next
	if previousNode.key == key {
		bucket.head = currentNode
		return
	}

	for currentNode != nil {
		if currentNode.key == key {
			previousNode.next = currentNode.next
			return
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % HASH_TABLE_ROWS
}

func main() {
	testHashTable := Init()
	testHashTable.Insert("hola", 111)
	testHashTable.Insert("trewef", 222)
	fmt.Println(testHashTable.Search("hola"))
	testHashTable.Delete("hola")
	fmt.Println(testHashTable.Search("hola"))
}
