package v3

type WySource struct {
	seed uint64
}

func (r *WySource) Seed(seed int64) {
	r.seed = uint64(seed)
}

func (r *WySource) Uint64() uint64 {
	r.seed += wyp0
	return wymum(r.seed^wyp1, r.seed)
}