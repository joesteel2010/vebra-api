package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
)

// FileTokenStorage implements the TokenStorage interface.
// It is one possible implementation for storing user credentials.
type FileTokenStorage struct {
	token         string
	tokenFileName string
}

// SetFileName sets the file name used to store the token in
func (ts *FileTokenStorage) SetFileName(path string) {
	ts.tokenFileName = path
}

// Save persists the token to a file
func (ts *FileTokenStorage) Save(token token) error {
	log.Printf("Saving token [%s]", token.tokenString)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token.tokenString))
	log.Printf("Saving token (encoded) [%s]", encodedToken)
	ts.token = encodedToken
	return ioutil.WriteFile(ts.tokenFileName, []byte(encodedToken), 0644)
}

// Load loads the persisted token from file
func (ts *FileTokenStorage) Load() (*token, error) {
	if _,err := os.Stat(ts.tokenFileName); err == nil {
		out, err := ioutil.ReadFile(ts.tokenFileName)
		ts.token = string(out)
		log.Printf("Loaded token [%s]", ts.token)
		return Token(ts.token), err
	}
	token := Token("")
	token.Invalidate()
	return token, nil
}
