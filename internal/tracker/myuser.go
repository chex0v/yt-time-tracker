package tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	UserMePath = "/users/me?fields=id,name,login,fullName,email,online"
)

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Login    string `json:"login,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email,omitempty"`
	Online   bool   `json:"online,omitempty"`
}

func (c Client) MyUserInfo() (User, error) {

	req, err := http.NewRequest("GET", c.Url+UserMePath, nil)

	req.Header.Add("Authorization", c.headerToken())

	if err != nil {
		log.Fatal(err)
	}

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("status code for user info error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return User{}, err
	}

	target := User{}

	err = json.Unmarshal(body, &target)

	if err != nil {
		return User{}, err
	}

	return target, nil

}
