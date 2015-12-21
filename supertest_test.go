package supertest

import "testing"
import "strings"
import "fmt"

// Basic

const HOST = "http://httpbin.org"

func TestGet(t *testing.T) {
	Request(HOST).
		Get("/get").
		Query("name=test").
		Expect(200).
		Expect("Content-Type", "application/json").
		End()
}

func TestPost(t *testing.T) {
	Request(HOST).
		Post("/post").
		Send(`{"name":"test"}`).
		Expect(200).
		Expect("Content-Type", "application/json").
		End()
}

func TestCheckStatus(t *testing.T) {
	defer checkError("Expected status: [204], but got: [200]")

	Request(HOST).
		Get("/get").
		Expect(204).
		End()
}

func TestCheckHeader(t *testing.T) {
	defer checkError("Expected header [name] to equal: [supertest], but got: [test]")

	Request(HOST).
		Get("/response-headers").
		Query("name=test").
		Expect("name", "supertest").
		End()
}

const TEXT_BODY = `User-agent: *
Disallow: /deny
`

func TestCheckBody_Text(t *testing.T) {
	Request(HOST).
		Get("/robots.txt").
		Expect(200).
		Expect(TEXT_BODY).
		End()
}

func TestCheckBody_Text_Error(t *testing.T) {
	defer checkError("Expected body:\nerror\nbut got:\n" + TEXT_BODY)

	Request(HOST).
		Get("/robots.txt").
		Expect(200).
		Expect("error").
		End()
}

func TestCheckBody_Json_String(t *testing.T) {
	body := `{
		"Content-Length": "68",
		"Content-Type": "application/json"
	}`

	Request(HOST).
		Get("/response-headers").
		Expect(200).
		Expect(body).
		End()
}

func TestCheckBody_Json_Map(t *testing.T) {
	body := map[string]string{
		"Content-Length": "68",
		"Content-Type":   "application/json",
	}

	Request(HOST).
		Get("/response-headers").
		Expect(200).
		Expect(body).
		End()
}

func TestCheckBody_Json_Struct(t *testing.T) {
	type Body struct {
		ContentLength string `json:"Content-Length"`
		ContentType   string `json:"Content-Type"`
	}

	body := Body{
		ContentLength: "68",
		ContentType:   "application/json",
	}

	Request(HOST).
		Get("/response-headers").
		Expect(200).
		Expect(body).
		End()
}

func checkError(suffix string) {
	err := recover()

	if err == nil {
		panic("test failed")
		return
	}

	str := fmt.Sprintf("%v", err)
	if !strings.HasSuffix(str, suffix) {
		panic(fmt.Sprintf("test failed - error is:\n%s\nbut expected error suffix:\n%s\n", str, suffix))
	}
}
