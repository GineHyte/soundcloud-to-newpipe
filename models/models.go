package models

var CLEAR = "\033[2J"
var RESET = "\033[0m"
var BOLD = "\033[1m"
var RED = "\033[31m"
var GREEN = "\033[32m"
var YELLOW = "\033[33m"
var BLUE = "\033[34m"
var MAGENTA = "\033[35m"

type Args struct {
	UserId string
	Token  string
	Output string
}
