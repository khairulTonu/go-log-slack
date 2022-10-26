package goslack

const (
	Success = iota + 1
	Warning
	Alert
)

var StatusMap = map[int]string{
	Success: "#00663a",
	Warning: "#eda200",
	Alert:   "#9e0505",
}
