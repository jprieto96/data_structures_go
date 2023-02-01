package main

const HASH_TABLE_ROWS = 100

type HashTable struct {
	array [HASH_TABLE_ROWS]*Bucket
}

type Bucket struct {
	node *BucketNode
}

type BucketNode struct {
	key  string
	next *BucketNode
}

func insert() bool {
	return false
}

func delete() bool {
	return false
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % HASH_TABLE_ROWS
}
