
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

Request("http://example.com")
  .Get("/hello")
  .Expect(200) // status
  .Expect("Content-Type", "application/json") // header
  .Expect(`{"name":"hello"}`) // body
  // or
  // .Expect(map[string]string{"name": "hello"})
```

### License
MIT

[travis-img]: https://img.shields.io/travis/haoxins/supertest.svg?style=flat-square
[travis-url]: https://travis-ci.org/haoxins/supertest
[license-img]: https://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[doc-img]: https://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/haoxins/supertest
