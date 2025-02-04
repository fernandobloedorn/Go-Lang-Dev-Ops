package tests

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/opening?id=2", UpperCaseHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))

}
