package singleton

var Single1 *singleton

type singleton struct {
	Number int
}

func init() {
	Single1 = &singleton{}
}
