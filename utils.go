package adventOfCode23

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Input(day int) (error, []string) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return errors.New("cannot load env"), nil
	}

	url := fmt.Sprintf("https://adventofcode.com/2023/day/%v/input", day)
	sessionCookie := os.Getenv("SESSION")
	if sessionCookie == "" {
		fmt.Println("SESSION environment variable is not set")
		return errors.New("session is not in env"), nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return errors.New("request creation error"), nil
	}

	// Add the session cookie to the request
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("error making request: %q", err)), nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return errors.New(fmt.Sprintf("Error reading response body: %q", err)), nil
	}

	var result []string
	var tmp []byte

	for _, v := range body {
		if v == byte('\n') {
			result = append(result, string(tmp))
			tmp = []byte{}
		} else {
			tmp = append(tmp, v)
		}
	}

	if len(tmp) > 0 {
		result = append(result, string(tmp))
	}
	return nil, result
}
