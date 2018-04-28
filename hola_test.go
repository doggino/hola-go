package hola

import "os"
import "testing"
import "net/http"
import "net/http/httptest"

var mux *http.ServeMux

// This is executed once per test suite
func TestMain(m *testing.M) {
  setUp()
  code := m.Run()
  os.Exit(code)
}

func setUp() {
  mux = http.NewServeMux()
  mux.HandleFunc("/hola", handleRequest)
}

func testHandleRequestBoilerplate(url string, expectedAnswer string, t *testing.T) {
  writer := httptest.NewRecorder()
  request, _ := http.NewRequest("GET", url, nil)
  mux.ServeHTTP(writer, request)

  if writer.Code != 200 {
    t.Errorf("Response code is %v instead of 200.", writer.Code)
  }
  answer := writer.Body.String()
  t.Logf("body: \"%v\"\n", answer)
  if answer != expectedAnswer {
    t.Errorf("Body is wrong: \"%v\", expected: \"%v\"\n", answer, expectedAnswer)
  }
}

func TestHandleRequestNoQuery(t *testing.T) {
  testHandleRequestBoilerplate("/hola", defaultAnswer, t)
}

func TestHandleRequestWeirdQuery(t *testing.T) {
  testHandleRequestBoilerplate("/hola?number=45", defaultAnswer, t)
}

func TestHandelRequestQueryNameEmpty(t *testing.T) {
  testHandleRequestBoilerplate("/hola?name=", defaultAnswer, t)
}

func TestHandleRequestExpectedQuery(t *testing.T) {
  testHandleRequestBoilerplate("/hola?name=perico", "hola, perico!", t)
}

func TestHandleRequestAndFail(t *testing.T) {
  t.Skip("this test must fail.")
  testHandleRequestBoilerplate("/hola?name=perico", "hola, manolo!", t)
}

func TestSkip(t *testing.T) {
  t.Skip("how to skip a test")
}

