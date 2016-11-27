app_error
=========

Why
---

Introducing a concept of "app error" which is "handled" errors -- like "yeah, I know it could happen and I'm taking good care of."  
App error can be exposed to clients so that they can tell what's happening inside a request, and what exactly causes the request to fail.


What
----

- Must not use `c.AbortWithError` for non-fatal errors
  - Use `c.Error` instead
- Must create custom error (i.e. `errors.New`) for "handled" errors
  - Create a const of an error string
  - Error string must follow the pattern `{status}.{package}.{struct/domain}...`

```go
r.Use(app_error.Wrap())

// or

r.Use(app_error.WrapWithCallback(func (c *gin.Context, body []byte, err error) {
	// notify the error to somewhere...
}))
```

``` go
package api

import ...

const ERROR_NAME_EMPTY = "422.api.sample.name.empty"  // This is the "app error"

func Sample(c *gin.Context) {
	// ...

	if somethingWentWrong {
		err := errors.New(AWESOME_REQUEST_ERROR_FOO_BAR)  // Create custom error
		c.Error(err)                                      // Must not use c.AbortWithError
		return
	}

	// ...

	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
```


Reference
---------

### Status codes

#### 4xx

- 400 BadRequest
- 401 Unauthorized
- 402 PaymentRequired
- 403 Forbidden
- 404 NotFound
- 405 MethodNotAllowed
- 406 NotAcceptable
- 407 ProxyAuthRequired
- 408 RequestTimeout
- 409 Conflict
- 410 Gone
- 411 LengthRequired
- 412 PreconditionFailed
- 413 RequestEntityTooLarge
- 414 RequestURITooLong
- 415 UnsupportedMediaType
- 416 RequestedRangeNotSatisfiable
- 417 ExpectationFailed
- 418 Teapot
- 422 UnprocessableEntity
- 423 Locked
- 424 FailedDependency
- 426 UpgradeRequired
- 428 PreconditionRequired
- 429 TooManyRequests
- 431 RequestHeaderFieldsTooLarge
- 451 UnavailableForLegalReasons

#### 5xx

- 500 InternalServerError
- 501 NotImplemented
- 502 BadGateway
- 503 ServiceUnavailable
- 504 GatewayTimeout
- 505 HTTPVersionNotSupported
- 506 VariantAlsoNegotiates
- 507 InsufficientStorage
- 508 LoopDetected
- 510 NotExtended
- 511 NetworkAuthenticationRequired
