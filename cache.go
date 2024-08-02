package main

type cache interface {
	Get(key string) string
	Set(key string, val string)
}
