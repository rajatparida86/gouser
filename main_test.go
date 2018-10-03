package main_test

import (
	"encoding/json"

	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rajatparida86/gouser"
)

var a main.App

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users (
  id SERIAL,
  name TEXT NOT NULL,
  age NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  CONSTRAINT users_pkey PRIMARY KEY (id)
)`

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(
		"gouser",
		"gouser",
		"gouser_db",
		"localhost",
		"5432")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}
func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART 1")
}
func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected empty array. Got %s", body)
	}
}
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("Expected %v. Got %v", expected, actual)
	}
}
func TestNonExistentUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user/11", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

	var r map[string]string
	json.Unmarshal(response.Body.Bytes(), &r)
	if r["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", r["error"])
	}
}

func TestCreateUser(t *testing.T) {
	clearTable()
	payload := []byte(`{"name":"test user","age":2}`)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["name"] != "test user" {
		t.Errorf("Expected user to be 'test user'. Got '%v'", m["name"])
	}
	if m["age"] != 2.0 {
		t.Errorf("Expected user age to be '2'. Got '%v'", m["age"])
	}
	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected user ID to be '1'. Got '%v'", m["id"])
	}
}
