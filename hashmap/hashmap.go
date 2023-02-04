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

func (hahsTable *HashTable) Delete() {

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
	testHashTable.Insert("hola", 222)
	testHashTable.Insert(1, "hola")
	fmt.Println(testHashTable.Search(1))
}
