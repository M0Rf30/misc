package misc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

// JSONIdentity is a struct to store nested key=value of IPFS json config
type JSONIdentity struct {
	Identity struct {
		PeerID  string
		PrivKey string
	}
}

// ReadLines reads a whole file into memory and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// UserHomeDir return the home path of current user
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// Exists returns whether the given file or directory exists or not
func Exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// ParseJSON returns a parsed json file
func (j *JSONIdentity) ParseJSON(path string) error {
	var err error
	config, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(config, &j)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
