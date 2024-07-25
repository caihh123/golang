package main

import (
	"testing"

	"github.com/google/btree"
)

type Item struct {
	Key   int
	Score int
}

func (i *Item) Less(item btree.Item) bool {
	data, ok := item.(*Item)
	if !ok {
		return false
	}
	if data.Score != i.Score {
		return data.Score > i.Score
	}
	return data.Key > i.Key
}

func TestTree(t *testing.T) {
	tree := btree.New(4)

	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	t.Log("length", tree.Len())
	t.Log("max", tree.Max())
	tree.Ascend(func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}

func TestTreeDelete(t *testing.T) {
	tree := btree.New(32)
	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	tree.Ascend(func(item btree.Item) bool {
		t.Log(item)
		return true
	})
	t.Log("删除2")
	tree.Delete(&Item{Key: 2, Score: 2})
	if tree.Get(&Item{Key: 2, Score: 2}) != nil {
		t.Error()
		return
	}
	tree.Ascend(func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}

func TestTreeAscend(t *testing.T) {
	tree := btree.New(32)
	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	tree.Ascend(func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}

func TestTreeAscendLessThan(t *testing.T) {
	tree := btree.New(32)
	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	tree.AscendLessThan(tree.Max(), func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}

func TestTreeAscendRange(t *testing.T) {
	tree := btree.New(32)
	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	tree.AscendRange(tree.Min(), tree.Max(), func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}

func TestTreeAscendGreaterOrEqual(t *testing.T) {
	tree := btree.New(32)
	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	tree.AscendGreaterOrEqual(tree.Min(), func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}

func TestTreeDescend(t *testing.T) {
	tree := btree.New(32)
	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	tree.Descend(func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}

func TestTreeDescendRange(t *testing.T) {
	tree := btree.New(32)
	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	tree.DescendRange(tree.Max(), tree.Min(), func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}

func TestTreeDescendGreaterThan(t *testing.T) {
	tree := btree.New(32)
	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	tree.DescendGreaterThan(tree.Min(), func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}

func TestTreeDescendLessOrEqual(t *testing.T) {
	tree := btree.New(32)
	for i := 1; i < 10; i++ {
		tree.ReplaceOrInsert(&Item{
			Key:   i,
			Score: i,
		})
	}
	tree.DescendLessOrEqual(tree.Max(), func(item btree.Item) bool {
		t.Log(item)
		return true
	})
}
