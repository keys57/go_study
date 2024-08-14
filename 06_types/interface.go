package main

type List interface {
	Add(index int, val any) error
	Append(val any)
	Delete(index int) error
}
