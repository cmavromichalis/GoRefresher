package main

type Cipher interface {
	validate(string) bool
	calculate(string, string, bool) (string, error)
	encode(string, string) (string, error)
	decode(string, string) (string, error)
}
