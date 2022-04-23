package db

import boltdb "github.com/boltdb/bolt"

type JKZDB struct {
	db boltdb.DB
	// Maps Key to Bucket
	RangeMap map[string][]byte
	// Map Indexes to Buckets
	IndexedFields map[string][]byte
}