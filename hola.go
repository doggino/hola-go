package hola

import (
	"fmt"
	"net/http"
)

func Serve() {
	fmt.Println("running...")
	http.Handle("/hola", http.HandlerFunc(handleRequest))
	err := http.ListenAndServe("127.0.0.1:8001", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("finished.")
}

const defaultAnswer = "hola, mundo!"

func getAnswer(names []string, ok bool) string {

	if !ok {
		return defaultAnswer
	}

	name := names[0]
	if len(name) < 1 {
		return defaultAnswer
	}

	return fmt.Sprintf("hola, %v!", names[0])
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	u := req.URL
	values := u.Query()
	names, ok := values["name"]
	fmt.Println(names)
	fmt.Println(len(names))

	answer := getAnswer(names, ok)

	w.Write([]byte(answer))
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		handleGet(w, req)
	default:
	}
}
