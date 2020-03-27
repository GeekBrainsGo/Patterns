package singleton

var single *singleton

var Single = &singleton{}

type singleton struct {
	Number int
}

func Single() *singleton {
	if single == nil {
		single = &singleton{

		}
	}
	return single
}