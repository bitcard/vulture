package mst

import (
	"crypto"
	"io"
)

type NodeStore interface {
	Get([]byte) *Node
	Put(*Node) []byte
	Remove([]byte)
	Copy() NodeStore
	Size() uint
}

// Only used for computing the hash for a given node
type HashableNode Node

func hashToWrite(hash []byte, hashSize int) []byte {
	if hash == nil {
		return make([]byte, hashSize)
	}
	return hash
}

func (n *HashableNode) Write(w io.Writer) error {
	// TODO: Test this if it's ever used in something production related. It really shouldn't be since
	// this node store is mainly meant for testing. Persistent node stores will need to thing more
	// about the serialization format.
	err := putUint32(n.level, w)
	if err != nil {
		return err
	}

	// Write key/values first
	err = putUint32(uint32(len(n.children)), w)
	if err != nil {
		return err
	}

	for _, child := range n.children {
		err = child.key.Write(w)
		if err != nil {
			return err
		}
		err = child.value.Write(w)
		if err != nil {
			return err
		}
	}

	// Assumes all hash sizes are exactly the same
	hashSize := 0
	if n.low != nil {
		hashSize = len(n.low)
	} else {
		for _, child := range n.children {
			if child.high != nil {
				hashSize = len(child.high)
				break
			}
		}
	}

	// Don't waste space on zeroed out hashes
	if hashSize > 0 {
		if n.low != nil {
			_, err = w.Write(hashToWrite(n.low, hashSize))
			if err != nil {
				return err
			}
		}

		for _, child := range n.children {
			_, err = w.Write(hashToWrite(child.high, hashSize))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type LocalNodeStore struct {
	dict map[string]*Node
	hash crypto.Hash
}

func NewLocalNodeStore(hash crypto.Hash) *LocalNodeStore {
	return &LocalNodeStore{dict: map[string]*Node{}, hash: hash}
}

func (ns *LocalNodeStore) Get(k []byte) *Node {
	return ns.dict[string(k)]
}

func (ns *LocalNodeStore) Put(n *Node) []byte {
	wn := HashableNode(*n)
	k := HashWritable(&wn, ns.hash)
	ns.dict[string(k)] = n
	return k
}

func (ns *LocalNodeStore) Remove(k []byte) {
	delete(ns.dict, string(k))
}

func (ns *LocalNodeStore) Copy() NodeStore {
	newDict := map[string]*Node{}
	for k, v := range ns.dict {
		newDict[k] = v
	}
	return &LocalNodeStore{newDict, ns.hash}
}

func (ns *LocalNodeStore) Size() uint {
	return uint(len(ns.dict))
}

func FindMissingNodes(ns NodeStore, hash []byte) [][]byte {
	if hash == nil {
		return [][]byte{}
	}
	n := ns.Get(hash)
	if n == nil {
		return [][]byte{hash}
	}
	missingNodes := [][]byte{}
	missingNodes = append(missingNodes, FindMissingNodes(ns, n.low)...)
	for _, child := range n.children {
		missingNodes = append(missingNodes, FindMissingNodes(ns, child.high)...)
	}
	return missingNodes
}
