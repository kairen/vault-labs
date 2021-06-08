package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

var (
	staticToken = os.Getenv("VAULT_TOKEN")
	vault_addr  = os.Getenv("VAULT_ADDR")
	username    = os.Getenv("VAULT_USER")
	password    = os.Getenv("VAULT_PASSOWRD")
)

func main() {
	client, err := api.NewClient(&api.Config{Address: vault_addr})
	if err != nil {
		fmt.Println(err)
		return
	}

	client.SetToken(staticToken)
	if len(username) > 0 && len(password) > 0 {
		userToken, err := userLogin()
		if err != nil {
			fmt.Println(err)
			return
		}
		client.SetToken(userToken)
	}

	secret, err := client.Logical().Read("secret/data/api/config")
	if err != nil {
		fmt.Println(err)
		return
	}

	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return
	}
	fmt.Printf("DB username: %v\n", m["db_username"])
	fmt.Printf("DB password: %v\n", m["db_password"])
}

func userLogin() (string, error) {
	// create a vault client
	client, err := api.NewClient(&api.Config{Address: vault_addr})
	if err != nil {
		return "", err
	}

	// to pass the password
	options := map[string]interface{}{
		"password": password,
	}
	path := fmt.Sprintf("auth/userpass/login/%s", username)

	// PUT call to get a token
	secret, err := client.Logical().Write(path, options)
	if err != nil {
		return "", err
	}

	token := secret.Auth.ClientToken
	return token, nil
}
