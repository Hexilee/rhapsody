package rady

import "reflect"

/*
FILE is a tag to bind a path with a file

Usage:

	type UserController struct {
		Controller 	`prefix:"api/v1"`
		FILE		`path:"static" file:"./index.html"`
	}

	type Root struct {
		UserController
	}

	func main() {
		CreateApplication(new(Root)).Run()
	}
 */
type FILE struct {
}

// STATIC is a tag bind a prefix with a directory
type STATIC struct {
}

var (
	FileType = reflect.TypeOf(FILE{})
	StaticType = reflect.TypeOf(STATIC{})
)