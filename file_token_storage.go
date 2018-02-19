package api

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

// FileTokenStorage implements the TokenStorage interface.
// It is one possible implementation for storing user credentials.
type FileTokenStorage struct {
	token         string
	tokenFileName string
}

// SetFileName sets the file name used to store the Token in
func (ts *FileTokenStorage) SetFileName(path string) {
	ts.tokenFileName = path
}

// Save persists the Token to a file
func (ts *FileTokenStorage) Save(token Token) error {
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token.tokenString))
	ts.token = encodedToken
	return ioutil.WriteFile(ts.tokenFileName, []byte(encodedToken), 0644)
}

// Load loads the persisted Token from file
func (ts *FileTokenStorage) Load() (*Token, error) {
	if _,err := os.Stat(ts.tokenFileName); err == nil {
		out, err := ioutil.ReadFile(ts.tokenFileName)
		ts.token = string(out)
		return NewToken(ts.token), err
	}
	token := NewToken("")
	token.Invalidate()
	return token, nil
}
