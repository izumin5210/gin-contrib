recovery
========

Why
---

Recover from errors and deal with them.


What
----

```go
r.Use(recovery.Wrap())

// or

r.Use(recovery.WrapWithCallback(func (c *gin.Context, body []byte, err interface{}) {
	// notify the error to somewhere...
}))
```
