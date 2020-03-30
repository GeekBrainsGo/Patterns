package singleton

var single *singleton

func Single() *singleton {
	if single == nil {
		single = &singleton{}
	}
	return single
}
