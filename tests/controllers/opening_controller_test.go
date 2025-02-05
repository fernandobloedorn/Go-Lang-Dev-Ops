package controllers_test 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// func TestMain(m *testing.M) {

// }

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{}"))
}

func TestShowOpening(t *testing.T) {
	req := httptest.NewRequest("GET", "/opening?id=2", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()

	defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Errorf("expected error to be nil got %v", err)
    }

	assert.Nil(t, err)

    // if string(data) != "ABC" {
    //     t.Errorf("expected ABC got %v", string(data))
    // }

	// fmt.Println(resp.StatusCode)
	fmt.Println(data)

	if !strings.Contains(data.String(), "id") {
        t.Errorf(`response body "%s" does not contain "ID"`, wr.Body.String(),)
    }

	assert.Equal(t, resp.StatusCode, http.StatusOK, "Show opening should return status 200")

}

func TestShowOpenings(t *testing.T) {
	req := httptest.NewRequest("GET", "/openings", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()

	fmt.Println(resp.StatusCode)

	assert.Equal(t, resp.StatusCode, http.StatusOK, "Show openings should return status 200")

}