
[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]
[![GoDoc][doc-img]][doc-url]

### supertest

* Based on [parnurzeal/gorequest](https://github.com/parnurzeal/gorequest)
* Inspired by [visionmedia/supertest](https://github.com/visionmedia/supertest)

### APIs

* [GoDoc](http://godoc.org/github.com/haoxins/supertest)

### Examples

```go
import . "github.com/haoxins/supertest"
import "testing"

func TestGet1(t *testing.T) {
  Request("http://httpbin.org", t).
    Get("/get").
    Query("name=test").
    Expect(200).
    Expect("Content-Type", "application/json").
    End()
}

func TestGet2(t *testing.T) {
  Request("http://example.com").
    Get("/hello").
    Expect(200). // status
    Expect("Content-Type", "application/json"). // header
    Expect(`{"name":"hello"}`). // body
    // or
    // Expect(map[string]string{"name": "hello"}).
    End()
}
```

### License
MIT

[travis-img]: https://img.shields.io/travis/haoxins/supertest.svg?style=flat-square
[travis-url]: https://travis-ci.org/haoxins/supertest
[license-img]: https://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[doc-img]: https://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/haoxins/supertest
