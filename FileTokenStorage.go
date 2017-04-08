package vebra

import (
	"encoding/base64"
	"io/ioutil"
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
func (ts *FileTokenStorage) Save(token string) error {
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))
	ts.token = encodedToken
	return ioutil.WriteFile(ts.tokenFileName, []byte(encodedToken), 0644)
}

// Load loads the persisted token from file
func (ts *FileTokenStorage) Load() (string, error) {
	out, err := ioutil.ReadFile(ts.tokenFileName)
	ts.token = string(out)
	return ts.token, err
}
