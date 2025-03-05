package loadbalancing

import (
	"fmt"
	"hash/fnv"

	"github.com/nanoDFS/Master/utils/crypto"
)

type ConsistentHashing struct {
}

func (ch ConsistentHashing) hashKey(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (ch ConsistentHashing) GetIndex(opts Opts) (int, error) {
	fullyHashedString := crypto.HashSHA256(opts.Key)
	if opts.Length == 0 {
		return 0, fmt.Errorf("opts length can't be zero")
	}
	index := int(ch.hashKey(fullyHashedString)) % opts.Length
	return index, nil
}
