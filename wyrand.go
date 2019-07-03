package wyhash

type WySource struct {
	seed uint64
}

//
func (r *WySource) Seed(seed int64) {
	r.seed = uint64(seed)
}

func (r *WySource) Int63() int64 {
	return int64(r.Uint64())
}

func (r *WySource) Uint64() uint64 {
	r.seed += wyp0
	return mum(r.seed^wyp1, r.seed)
}
