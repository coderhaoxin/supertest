package supertest

import "github.com/parnurzeal/gorequest"
import "github.com/pkg4go/urlx"
import "encoding/json"
import "net/http"
import "reflect"
import "strings"
import "errors"
import "fmt"

type Agent struct {
	host    string
	path    string
	method  string
	asserts [][]interface{}
	agent   *gorequest.SuperAgent
}

func Request(host string) *Agent {
	r := &Agent{}
	r.host = host
	r.agent = gorequest.New()
	return r
}

func (r *Agent) Get(path string) *Agent {
	host, _ := urlx.Resolve(r.host, path)
	r.agent.Get(host)
	return r
}

func (r *Agent) Post(path string) *Agent {
	host, _ := urlx.Resolve(r.host, path)
	r.agent.Post(host)
	return r
}

func (r *Agent) Put(path string) *Agent {
	host, _ := urlx.Resolve(r.host, path)
	r.agent.Put(host)
	return r
}

func (r *Agent) Delete(path string) *Agent {
	host, _ := urlx.Resolve(r.host, path)
	r.agent.Delete(host)
	return r
}

func (r *Agent) Patch(path string) *Agent {
	host, _ := urlx.Resolve(r.host, path)
	r.agent.Patch(host)
	return r
}

func (r *Agent) Head(path string) *Agent {
	host, _ := urlx.Resolve(r.host, path)
	r.agent.Head(host)
	return r
}

func (r *Agent) Options(path string) *Agent {
	host, _ := urlx.Resolve(r.host, path)
	r.agent.Options(host)
	return r
}

func (r *Agent) Set(param, value string) *Agent {
	r.agent.Set(param, value)
	return r
}

func (r *Agent) SetBasicAuth(username, password string) *Agent {
	r.agent.SetBasicAuth(username, password)
	return r
}

func (r *Agent) AddCookie(cookie *http.Cookie) *Agent {
	r.agent.AddCookie(cookie)
	return r
}

func (r *Agent) AddCookies(cookies []*http.Cookie) *Agent {
	r.agent.AddCookies(cookies)
	return r
}

func (r *Agent) Type(ts string) *Agent {
	r.agent.Type(ts)
	return r
}

func (r *Agent) Query(q interface{}) *Agent {
	r.agent.Query(q)
	return r
}

func (r *Agent) Send(data interface{}) *Agent {
	r.agent.Send(data)
	return r
}

func (r *Agent) Expect(args ...interface{}) *Agent {
	r.asserts = append(r.asserts, args)
	return r
}

func (r *Agent) End(cbs ...func(response gorequest.Response, bodyString string, errors []error)) {
	r.agent.End(func(res gorequest.Response, body string, errs []error) {

		contentType := res.Header.Get("Content-Type")
		status := res.StatusCode

		for _, assert := range r.asserts {
			if len(assert) == 1 {
				v := assert[0]

				if getType(v) == "int" {
					// status
					checkStatus(v, status)
				} else {
					// body
					checkBody(v, body, contentType)
				}
			} else if len(assert) == 2 {

				if getType(assert[0]) == "int" {
					// Expect(200, `body`)
					checkStatus(assert[0], status)
					checkBody(assert[1], body, contentType)
				} else if getType(assert[0]) == "string" {
					// Expect("Content-Type", "application/json")
					checkHeader(res.Header, assert[0], assert[1])
				} else {
					panic(errors.New("Unknown Expect behavior"))
				}
			} else {
				panic(errors.New("Expect only accept one or two args"))
			}
		}

		if len(cbs) > 0 {
			cbs[0](res, body, errs)
		}

	})
}

func getType(v interface{}) string {
	return reflect.ValueOf(v).Kind().String()
}

func checkStatus(status interface{}, actual int) {
	expect := status.(int)
	if expect != actual {
		panic(fmt.Errorf("Expected status: [%d], but got: [%d]", expect, actual))
	}
}

func checkHeader(header http.Header, key, val interface{}) {
	k := key.(string)
	actual := header.Get(k)
	expect := val.(string)
	if actual != expect {
		panic(fmt.Errorf("Expected header [%s] to equal: [%s], but got: [%s]", k, expect, actual))
	}
}

func checkBody(tobe interface{}, body, contentType string) {
	// only support text, json
	var expect string

	if strings.HasPrefix(contentType, "application/json") {
		// json TODO: more content types
		if getType(tobe) == "string" {
			expect = tobe.(string)
		} else {
			buf, err := json.Marshal(tobe)
			if err != nil {
				panic(err)
			}

			expect = string(buf[0:len(buf)])
		}

		if trim(expect) != trim(body) {
			panic(fmt.Errorf("Expected body:\n%s\nbut got:\n%s", trim(expect), trim(body)))
		}
	} else if strings.HasPrefix(contentType, "text/") {
		// text
		expect = tobe.(string)

		if expect != body {
			panic(fmt.Errorf("Expected body:\n%s\nbut got:\n%s", expect, body))
		}
	} else {
		panic(fmt.Errorf("content-type: %s not supported", contentType))
	}
}

func trim(str string) string {
	return strings.Replace(strings.Replace(strings.Replace(str, "\n", "", -1), "\t", "", -1), " ", "", -1)
}
