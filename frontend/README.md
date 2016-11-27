frontend
========

Why
---

It's very useful when you're creating SPA, wherein a server always responds with index.html even when a route for a path doesn't defined.


What
----

Fallback to a handler when no handler for a path is defined in a router.

```go
r.Use(frontend.Wrap(func (c *gin.Context) {
	c.HTML(http.StatusOK, "index", nil)
}))
```
