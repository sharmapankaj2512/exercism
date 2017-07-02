package accumulate

const testVersion = 1

func Accumulate(xs []string, f func(string) string) []string {
	ys := make([]string, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}
