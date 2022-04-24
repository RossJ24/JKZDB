package db

import "github.com/boltdb/bolt"

type JKZDB struct {
	db bolt.DB
	// Maps Key Type to Bucket (Index, primary, etc.)
	RangeMap map[string][]byte
	// Map Indexes to Buckets
	IndexedFields map[string][]byte
}
