package loadbalancing

import (
	"hash/fnv"

	"github.com/nanoDFS/Master/utils/crypto"
)

type ConsistentHashing struct {
}

func NewConsistentHashing() *ConsistentHashing {
	return &ConsistentHashing{}
}

func (ch *ConsistentHashing) hashKey(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (ch *ConsistentHashing) GetIndex(key string, length int) int {
	fullyHashedString := crypto.HashSHA256(key)
	index := int(ch.hashKey(fullyHashedString)) % length
	return index
}
