package singleton

var Single *singleton

type singleton struct {
	Number int
}

func init() {
	Single = &singleton{}
}