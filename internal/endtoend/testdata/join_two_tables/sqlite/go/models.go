// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package querytest

import ()

type Bar struct {
	ID int64
}

type Baz struct {
	ID int64
}

type Foo struct {
	BarID int64
	BazID int64
}