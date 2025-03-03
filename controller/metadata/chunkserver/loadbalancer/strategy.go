package loadbalancing

type Opts struct {
	Key    string
	Length int
}
type LoadBalancingStrategy interface {
	GetIndex(opts Opts) int
}
