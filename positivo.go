package pgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func Login(username string, password string) (string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://sso.specomunica.com.br/connect/token", strings.NewReader("username="+username+"&password="+password+"&grant_type=password&client_id=hubpsd&client_secret=DA5730D8-90FF-4A41-BFED-147B8E0E2A08&scope=openid%20offline_access%20integration_info"))

	if err != nil {
		return "Não foi possível criar a requesição:", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return "Não foi possível enviar a requisão:", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "Não foi possível ler a resposta:", err
	}

	//Check if response is valid
	if res.StatusCode != 200 {
		//Unmarshal error response
		var errResp Error
		json.Unmarshal(body, &errResp)
		return "Não foi possível fazer a autenticação: " + errResp.ErrorDescription, err
	}
	//Unmarshal response
	var token Token
	json.Unmarshal(body, &token)

	return token.AccessToken, nil
}

func GetHomework(token string) string {
	return "https://plus-app.studos.com.br/auth/psd?jwt=" + token + "&redirect=/student/central"
}
