package pgo

import (
	"encoding/json"
	"errors"
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

type Solucoes []struct {
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

type Userinfo struct {
	Sub           string `json:"sub"`
	AuthTime      int    `json:"auth_time"`
	Idp           string `json:"idp"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	IntegrationID string `json:"integration_id"`
	Amr           string `json:"amr"`
	Schools       string `json:"schools"`
}

type School struct {
	ID            string   `json:"id"`
	IntegrationID string   `json:"integration_id"`
	UserID        string   `json:"user_id"`
	Name          string   `json:"name"`
	Roles         []string `json:"roles"`
	TimeZone      string   `json:"time_zone"`
	URL           string   `json:"url"`
}

type UserClass []struct {
	Value  string   `json:"value"`
	Label  string   `json:"label"`
	Turmas []Turmas `json:"turmas"`
}
type Turmas struct {
	NomeTurma   string `json:"nomeTurma"`
	TurmaValida bool   `json:"turmaValida"`
	NomeSerie   string `json:"nomeSerie"`
}

type Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func Login(username string, password string) (string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://sso.specomunica.com.br/connect/token", strings.NewReader("username="+username+"&password="+password+"&grant_type=password&client_id=hubpsd&client_secret=DA5730D8-90FF-4A41-BFED-147B8E0E2A08&scope=openid%20offline_access%20integration_info"))

	if err != nil {
		return "", errors.New("Não foi possível criar a requesição:" + err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return "", errors.New("Não foi possível enviar a requisão:" + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Não foi possível ler a resposta:" + err.Error())
	}

	//Check if response is valid
	if res.StatusCode != 200 {
		//Unmarshal error response
		var errResp Error
		json.Unmarshal(body, &errResp)
		return "", errors.New("Não foi possível fazer a autenticação: " + errResp.ErrorDescription)
	}
	//Unmarshal response
	var token Token
	json.Unmarshal(body, &token)

	return token.AccessToken, nil
}

func GetUserInfo(token string) (string, string, string, error) {
	userinforesponse, err := tokenRequest("https://sso.specomunica.com.br/connect/userinfo", "GET", token)
	if err != nil {
		return "", "", "", errors.New("Um erro aconteceu:" + err.Error())
	}
	var userinfo Userinfo
	json.Unmarshal([]byte(userinforesponse), &userinfo)
	var school School
	json.Unmarshal([]byte(userinfo.Schools), &school)
	return school.ID, school.UserID, school.Roles[0], nil

} //Retorna o id da escola, o id do usuário e o papel do usuário

func GetUserName(token string) (string, error) {
	//Retorna o nome do usuário
	userinforesponse, err := tokenRequest("https://sso.specomunica.com.br/connect/userinfo", "GET", token)
	if err != nil {
		return "", errors.New("Um erro aconteceu:" + err.Error())
	}
	var usrname Userinfo
	json.Unmarshal([]byte(userinforesponse), &usrname)
	return usrname.Name, nil
}

func GetClass(token string, userId string) (string, error) {
	getclassresponse, err := tokenRequest("https://apihub.positivoon.com.br/api/NivelEnsino?usuarioId="+userId, "GET", token)
	if err != nil {
		return "", err
	}
	//unmarshal response
	var userclass UserClass
	json.Unmarshal([]byte(getclassresponse), &userclass)
	return userclass[0].Value, nil

}

func GetRawSolutions(token string, role string, class string, schoolID string) (string, error) {
	getRawSolutions, err := tokenRequest("https://apihub.positivoon.com.br/api/Categoria/Solucoes/Perfil/"+role+"?NivelEnsino="+class+"&IdEscola="+schoolID, "GET", token)
	if err != nil {
		return "", err
	}
	return getRawSolutions, nil
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
}

func tokenRequest(url string, method string, token string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", errors.New("Não foi possível criar a requesição:" + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		return "", errors.New("Não foi possível enviar a requisão:" + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Não foi possível ler a resposta:" + err.Error())
	}

	return string(body), nil
}
