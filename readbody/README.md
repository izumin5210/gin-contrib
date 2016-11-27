readbody
========

Why
---

Safely read a request body from a context without breaking following processes.
Useful for capturing it within a middleware.


What
----

```go
r.Use(func (c *gin.Context) {
	body := readbody.Read(c)

	c.Next()  // Never breaks main processes

	// `body` can be still accessible here
})
```
