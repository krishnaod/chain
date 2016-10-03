package state

import (
	"chain/protocol/bc"
	"chain/protocol/patricia"
)

// PriorIssuances maps an "issuance hash" to the time (in Unix millis)
// at which it should expire from the issuance memory.
type PriorIssuances map[bc.Hash]uint64

// Snapshot encompasses a snapshot of entire blockchain state. It
// consists of a patricia state tree and the issuances memory.
type Snapshot struct {
	Tree      *patricia.Tree
	Issuances PriorIssuances
}

func (s *Snapshot) PruneIssuances(timestampMS uint64) {
	// Delete expired issuance memory from the snapshot.
	for hash, expiryMS := range s.Issuances {
		if timestampMS > expiryMS {
			delete(s.Issuances, hash)
		}
	}
}

// Copy makes a copy of provided snapshot.
func Copy(original *Snapshot) *Snapshot {
	c := &Snapshot{
		Tree:      patricia.Copy(original.Tree),
		Issuances: make(PriorIssuances, len(original.Issuances)),
	}
	for k, v := range original.Issuances {
		c.Issuances[k] = v
	}
	return c
}

// Empty returns an empty state snapshot.
func Empty() *Snapshot {
	return &Snapshot{
		Tree:      new(patricia.Tree),
		Issuances: make(PriorIssuances),
	}
}
