package loadbalancing

import (
	"hash/fnv"
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
	hash := ch.hashKey(key)
	index := int(hash) % length
	return index
}
