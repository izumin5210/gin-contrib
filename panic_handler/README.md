panic_handler
=============

Why
---

Middleware for https://github.com/creasty/panicsync


What
----

```go
r.Use(panic_handler.Wrap())

// or

r.Use(panic_handler.WrapWithCallback(func (c *gin.Context, body []byte, info *panicsync.Info) {
	info.Print()

	// notify the error to somewhere...
}))
```

```go
func Sample(c *gin.Context) {
	ph := panic_handler.Get(c)

	go func() {
		defer ph.Sync()

		// ...
	}()

	// ...
}
```
