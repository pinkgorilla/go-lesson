package main

import "context"

type key string

var v = key("v")

func ContextValue() {
	pctx := context.Background()
	ctx := context.WithValue(pctx, v, "secret")
	i, ok := ctx.Value(k)
}
