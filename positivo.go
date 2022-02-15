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

type Solutions []struct {
	ID       string     `json:"id"`
	Nome     string     `json:"nome"`
	Cor      string     `json:"cor"`
	Ordem    int        `json:"ordem"`
	Ativo    bool       `json:"ativo"`
	Solucoes []Solucoes `json:"solucoes"`
}
type Solucoes struct {
	ID               string `json:"id"`
	Nome             string `json:"nome"`
	Descricao        string `json:"descricao"`
	Arquivo          string `json:"arquivo"`
	Link             string `json:"link"`
	Ativo            bool   `json:"ativo"`
	TipoRenderizacao string `json:"tipoRenderizacao"`
	Slug             string `json:"slug"`
	Ordem            int    `json:"ordem"`
	DataCadastro     string `json:"dataCadastro"`
	Novo             bool   `json:"novo"`
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
} //Retorna a url de acesso ao portal do aluno

func GetMessages(token string) string {
	return "https://web.specomunica.com.br/link/feed?jwt=" + token
}

func RequestSupport(token string) string {
	return "https://web.specomunica.com.br/link/atendimento?jwt=" + token
}

func GetAgenda(token string) string {
	return "https://web.specomunica.com.br/link/agenda?jwt=" + token
}

func GetBooks(token string) (string, error) {
	return tokenRequest("https://livro-digital-estante.prd.positivoon.com.br/v3/livros?busca=&componenteCurricular=&nivelEnsino=&serie=&volume=", "GET", token)
} //Retorna um json com os livros do aluno

func GetUserinfo(token string) (string, error) {
	userinforesponse, err := tokenRequest("http://sso.specomunica.com.br/connect/userinfo", "GET", token)
	if err != nil {
		return "Um erro aconteceu:", err
	}
	return userinforesponse, nil

} //Retorna um json com os dados do usuário

func tokenRequest(url string, method string, token string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "Não foi possível criar a requesição:", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		return "Não foi possível enviar a requisão:", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "Não foi possível ler a resposta:", err
	}

	return string(body), nil
}
