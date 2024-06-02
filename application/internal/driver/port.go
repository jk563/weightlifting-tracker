// Package driver details the various ways of running the application
package driver

// Port for drivers
type Port interface {
	Init()
	Run()
}
