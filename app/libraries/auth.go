package libraries

import (
	"fmt"
	"net/url"
	"time"
)

type (
	UserCredential struct {
		Username string
		Password string
		Ttl      int
	}
)

func (uc *UserCredential) CreateToken() (string, bool) {
	uName := Env("AUTH_USERNAME")
	uPsswd := Env("AUTH_PASSWORD")
	prefix := Env("PREFIX_AUTH_KEY")

	if uc.Username != uName {
		return "Invalid Username.", false
	}

	if uc.Password != uPsswd {
		return "Invalid Password.", false
	}

	stringToHash := url.QueryEscape("#" + uc.Username + "@|@" + uc.Password + "#")

	generateToken, err := HashPassword(stringToHash)
	if err != nil {
		return err.Error(), false
	}

	keyName := prefix + generateToken
	value := `{"creds":"` + uName + `:` + uPsswd + `","live":true,"created_at": "` + time.Now().String() + `"}`

	fmt.Println("Key Name : " + keyName)
	rdb := Redis()
	ttl := time.Duration(uc.Ttl) * time.Second
	rdb.Set(keyName, value, time.Duration(ttl))

	return generateToken, true
}

func CheckAuthKey(requestedKey string) bool {
	prefix := Env("PREFIX_AUTH_KEY")
	searchKey := prefix + requestedKey
	rdb := Redis()
	getKey := rdb.Get(searchKey)
	if err := getKey.Err(); err != nil {
		return false
	}

	key, err := getKey.Result()
	if err != nil {
		fmt.Printf("unable to GET user token. error: %v", err)
		return false
	}

	if key == "" {
		return false
	}

	return true
}
