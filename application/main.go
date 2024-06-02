// Package main does all the things
package main

import (
	"jamiekelly/lifts/internal/driver"
)

var runtime driver.Port

func init() {
	runtime = driver.Lambda{}
	runtime.Init()
}

func main() {
	runtime.Run()
}
