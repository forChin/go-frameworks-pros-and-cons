## Frameworks

| Name                                      | Stars         | Description   | Pros | Cons |
| -------------                             |:-------------:|:-------------:|-------------|-------------|
| [Gin](https://github.com/gin-gonic/gin)   | 63.7k         |Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.| <ul><li>Most popular.</li><li>Runs comparatively faster compared to other Golang frameworks.</li><li>A lot of tools are out of the box.</li><li>Good examples online and decent community.</li></ul> |<ul><li>Gin doesn't support [regexps in its routes](https://github.com/gin-gonic/gin/issues/229)</li><li>Gin's logger is not no-op (if efficient logging matters, which it should if you aim for performance)</li><li>Not so idiomatic.</li><li>Gin is too much magic.</li></ul>  |
| [beego](https://github.com/beego/beego)                                  | 29k     | beego is an open-source, high-performance web framework for the Go programming language. |<ul><li>Provides integrated ORM, built-in cache handler, logger, etc.</li><li>No need for a third-party library. Very unlikely though!</li><li>MVC Architecture.</li><li>It is mostly used to develop enterprise web applications.</li></ul> | <ul><li>**Overwhelming.**</li><li>Non-idiomatic code.</li><li>Less community support.</li></ul>
| [kit](https://github.com/go-kit/kit)                                  | 24k     | Go kit is a programming toolkit for building microservices (or elegant monoliths) in Go. | <ul><li>Separation of concerns : Go- kit nicely separates the business logic, endpoint and transport,Low-level libraries.</li><li>Adapters for common infrastructure components</li></ul> | <ul><li>Heavy use of interfaces</li><li>Overhead of adding API : To create a simple api you need to add a lot of code Function in the interface, Implementation, Endpoint factory function, Transport function, Request encoder, Request decoder, Response encoder and Response decoder, Add the endpoint to the server, Add the endpoint to the client etc.</li><li>**Overwhelming.**</li></ul>|
| [echo](https://github.com/labstack/echo)| 24k | High performance, minimalist Go web framework| <ul><li>Highly optimized HTTP router.</li><li>Many built-in middleware to use, or define your own. Middleware can beset at root, group or route level.</li><li>minimalist Go web framework.</li></ul>| <ul><li></li></ul>|
| [fiber](https://github.com/gofiber/fiber) | 22.8k| Fiber is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind. | <ul><li>Extreme performance: JSON serialization, database access and server-side template composition.</li><li>Rich routing.</li><li>Main focus of minimalism.</li></ul>| <ul><li>No support for http2 (in TODO)</li></ul>|
| [go-zero](https://github.com/zeromicro/go-zero) | 21k| A cloud-native Go microservices framework with cli tool for productivity.| <ul><li>Builtin concurrency control, adaptive circuit breaker, adaptive load shedding, auto trigger, auto recover</li><li>plenty of builtin microservice management and concurrent toolkits</li></ul>| <ul><li>Less community support.</li><li>Hard to find info.</li><li>**Overwhelming.**</li></ul>|
| [kratos](https://github.com/go-kratos/kratos)| 19.1k | Your ultimate Go microservices framework for the cloud-native era. | <ul><li>Cover the various utilities for business development.</li></ul> | <ul><li>Less community support.</li><li>Hard to find info.</li><li>**Overwhelming.**</li></ul> |
| [fasthttp](https://github.com/valyala/fasthttp)| 18.6k| Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http| <ul><li>Fast</li></ul>| <ul><li>fasthttp was design for some high performance edge cases. Unless your server/client needs to handle thousands of small to medium requests per seconds and needs a consistent low millisecond response time fasthttp might not be for you.</li><li>Not easy to use.</li><li>Fasthttp doesn't provide ServeMux</li><li>fasthttp is mainly used for requests.</li><li>fasthttp doesn't support HTTP/2.0 (in TODO)</li></ul> |
| [httprouter](https://github.com/julienschmidt/httprouter)| 14.6k| A high performance HTTP request router that scales well| <ul><li>It scales well even with very long paths and a large number of routes.</li></ul>| <ul><li>Poor functionality.</li><li>**Infrequent code updates.**</li><li>A high performance HTTP request router that scales well</li></ul>|
| [chi](https://github.com/go-chi/chi)| 12.4k| lightweight, idiomatic and composable router for building Go HTTP services| <ul><li>Designed for modular/composable APIs - middlewares, inline middlewares, route groups and sub-router mounting</li><li>100% compatible with net/http - use any http or middleware pkg in the ecosystem that is also compatible with net/http</li><li>A good base middleware stack</li><li>Good examples online and decent community.</li></ul>| <ul><li>Two downsides of chi are that it doesn't automatically handle OPTIONS requests, and it doesn't set an Allow header in 405 responses. If you're building a web application or a private API, then those things probably aren't a big problem — but if you're building a public API it's something to be aware of. Like httprouter, it also doesn't support host-based routes.</li><li>As a note of caution, there have been 5 major version increments for chi in the past 6 years — most of which include breaking changes. History isn't necessarily a predictor of the future, but does suggest that backward-compatibility and not making breaking changes is less of a priority for chi than some of the other routers on this shortlist. In contrast httprouter and gorilla/mux haven't made any breaking changes in this time.</li></ul> |

