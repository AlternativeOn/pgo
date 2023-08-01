package pgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

/* TYPES DO LOGIN */
type LoginUsuario struct {
	Sucesso  bool   `json:"sucesso"`  //true ou false
	Mensagem string `json:"mensagem"` //mensagem do resultado do login
	Dados    Dados  `json:"dados"`    //pode retornar uma string
}

type Dados struct {
	AccessToken         string    `json:"access_token"`          //Token de acesso e atividades no Studos.
	AccessTokenParceiro string    `json:"access_token_parceiro"` //Usado para as comunicações
	Alias               string    `json:"alias"`
	ExpiresIn           int       `json:"expires_in"`
	TokenType           string    `json:"token_type"`
	RefreshToken        string    `json:"refresh_token"`
	Scope               string    `json:"scope"`
	Schools             []Schools `json:"schools"`
}

type Schools struct {
	ID            string      `json:"id"`
	IntegrationID interface{} `json:"integration_id"`
	UserID        string      `json:"user_id"`
	Name          string      `json:"name"`
	Roles         []string    `json:"roles"`
}

/* TYPE DO LOGIN PRIMITIVO */

type LoginPrimitivoDadosEscola struct {
	ID            string   `json:"id"`
	IntegrationID string   `json:"integration_id"`
	UserID        string   `json:"user_id"`
	Name          string   `json:"name"`
	Roles         []string `json:"roles"`
	TimeZone      string   `json:"time_zone"`
	URL           string   `json:"url"`
}

type LoginPrimitvoDadosUsuario struct {
	Sub           string `json:"sub"`
	AuthTime      int    `json:"auth_time"`
	Idp           string `json:"idp"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	IntegrationID string `json:"integration_id"`
	Amr           string `json:"amr"`
	LoginPrimitivoDadosEscola       string `json:"schools"`
}

type LoginPrimitivoDadosSerie []struct {
	Value  string   `json:"value"`
	Label  string   `json:"label"`
	LoginPrimitivoDadosTurma []Turmas `json:"turmas"`
}

type LoginPrimitivoDadosTurma struct {
	NomeTurma   string `json:"nomeTurma"`
	TurmaValida bool   `json:"turmaValida"`
	NomeSerie   string `json:"nomeSerie"`
}

/* TYPES DOS LIVROS */
type Livro struct {
	ComponenteCurricular string `json:"componenteCurricular"`
	Volume               string `json:"volume"`
	Arquivos             []struct {
		Tipo      string `json:"tipo"`
		IDArquivo string `json:"idArquivo"`
		Caminho   string `json:"caminho"`
	} `json:"arquivos"`
}

type InfoLivro struct {
	ComponenteCurricular string `json:"componenteCurricular"`
	Volume               string `json:"volume"`
	Tipo                 string `json:"tipo"`
	URL                  string `json:"caminho"`
}

/* TYPES DA LIBRARY */

type Token struct { //Usado para retornar a token do usuário
	Token         string
	TokenParceiro string
	IdEscola      string
	NomeEscola    string
	IdUsuario     string
}

type TokenPrimitiva struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type Recursos struct {
	Mensagens   string
	Agenda      string
	Atendimento string
	Studos      string
}

type ErroPrimitivo struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

var indexDaEscola = 0 //usado apenas quando necessário.

func Login(username string, password string) (*Token, error) {
	/* PRIMEIRA PARTE DO LOGIN
	 * Aqui é coletado o usuário e senha do usuário, e então é retornado a primeira token necessária
	 */
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("grant_type", "password")
	_ = writer.WriteField("password", password)
	_ = writer.WriteField("username", username)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("POST", "https://apihub.positivoon.com.br/api/login/token", payload)

	if err != nil {
		return nil, errors.New("Não foi possível criar a requesição:" + err.Error())
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)

	if err != nil {
		return nil, errors.New("Não foi possível enviar a requisão:" + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Não foi possível ler a resposta:" + err.Error())
	}

	//Verifica se a respostas está ok
	//Se o login estiver errado, o site retorna uma string na struct Dados, impossibilitando a leitura do JSON pelo go.
	//A melhor solução é verificar se o status é 200
	//Outros códigos HTTP de possivel retorno é 401 (usuário/senha errada) ou 500
	//Implemente no front-end uma verificação do código HTTP.
	if res.StatusCode != 200 {
		return nil, errors.New("Não foi possível fazer a autenticação!\nStatus HTTP: " + fmt.Sprint(res.StatusCode))
	}
	//Lê a resposta
	var respDoPrimeiroLogin LoginUsuario
	err = json.Unmarshal(body, &respDoPrimeiroLogin)
	if err != nil {
		panic(err)
	}
	dadosDoPrimeiroLogin := respDoPrimeiroLogin.Dados
	primeiraToken := dadosDoPrimeiroLogin.AccessToken
	//fmt.Println(dadosDoPrimeiroLogin)

	/* SEGUNDA PARTE DO LOGIN
	* Sim, o login é divido em duas partes, para obter a token na escola selecionada.
	* Aqui eu achei melhor escolher o index da escola, já que no meu caso funcionaria melhor a escola com o index 1
	(no caso que um professor com duas escolar usar), mas caso não tenha nenhum, usar o index 0.
	* É possivel mudar a seleção do index
	*/
	quantidadeEscolas := len(dadosDoPrimeiroLogin.Schools)
	//fmt.Println("QE:", quantidadeEscolas)
	if quantidadeEscolas >= 2 {
		indexDaEscola = 1
	}

	escolaUsuario := dadosDoPrimeiroLogin.Schools[indexDaEscola]
	fmt.Println(escolaUsuario.ID)

	// Segunda request, para trocar a token
	payload = &bytes.Buffer{}
	writer = multipart.NewWriter(payload)
	_ = writer.WriteField("grant_type", "change_school")
	_ = writer.WriteField("access_token", primeiraToken)
	_ = writer.WriteField("school_id", escolaUsuario.ID)
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req, err = http.NewRequest("POST", "https://apihub.positivoon.com.br/api/login/token", payload)

	if err != nil {
		return nil, errors.New("Não foi possível criar a requesição:" + err.Error())
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, err = client.Do(req)

	if err != nil {
		return nil, errors.New("Não foi possível enviar a requisão:" + err.Error())
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Não foi possível ler a resposta:" + err.Error())
	}

	//Verifica se a respostas está ok
	//Se o login estiver errado, o site retorna uma string na struct Dados, impossibilitando a leitura do JSON pelo go.
	//A melhor solução é verificar se o status é 200
	//Outros códigos HTTP de possivel retorno é 401 (usuário/senha errada) ou 500
	//Implemente no front-end uma verificação do código HTTP.
	if res.StatusCode != 200 {
		return nil, errors.New("Não foi possível fazer a autenticação!\nStatus HTTP: " + fmt.Sprint(res.StatusCode))
	}

	var respDoSegundoLogin LoginUsuario
	err = json.Unmarshal(body, &respDoSegundoLogin)
	if err != nil {
		panic(err)
	}
	dadosDoSegundoLogin := respDoSegundoLogin.Dados

	token := &Token{
		Token:         dadosDoSegundoLogin.AccessToken,
		TokenParceiro: dadosDoSegundoLogin.AccessTokenParceiro,
		IdEscola:      escolaUsuario.ID,
		NomeEscola:    escolaUsuario.Name,
		IdUsuario:     escolaUsuario.UserID,
	}
	res.Body.Close()
	return token, nil
}

func LegacyLogin(username string, password string) (*TokenPrimitiva, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("POST", "https://sso.specomunica.com.br/connect/token", strings.NewReader("username="+username+"&password="+password+"&grant_type=password&client_id=hubpsd&client_secret=DA5730D8-90FF-4A41-BFED-147B8E0E2A08&scope=openid%20offline_access%20integration_info"))
	if err != nil {
		return nil, errors.New("Não foi possível criar a requesição:" + err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Não foi possível enviar a requisão:" + err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Não foi possível ler a resposta:" + err.Error())
	}

	//Verifica se a resposta é valida
	if res.StatusCode != 200 {
		//Decompressa a mensagem de erro
		var errResp ErroPrimitivo
		json.Unmarshal(body, &errResp)
		return nil, errors.New("Não foi possível fazer a autenticação: " + errResp.ErrorDescription)
	}
	//Decompressa a resposta
	var token TokenPrimitiva
	json.Unmarshal(body, &token)

	res.Body.Close()
	tokenLegada := &TokenPrimitiva {
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn: token.ExpiresIn,
		TokenType: token.TokenType
	}
	return tokenLegada, nil
}

func ObterRecursos(idEscola string, userToken string) *Recursos {
	msg := &Recursos{
		Mensagens:   "https://web.escolaapp.com/link/feed?contextoId=" + idEscola + "&jwt=" + userToken,
		Agenda:      "https://web.escolaapp.com/link/agenda?contextoId=" + idEscola + "&jwt=" + userToken,
		Atendimento: "https://web.escolaapp.com/link/atendimento?contextoId=" + idEscola + "&jwt=" + userToken,
		Studos:      "https://plus-app.studos.com.br/auth/psd?jwt=" + userToken,
	}

	return msg
}

func ObterLivros(token string) ([]InfoLivro, error) {
	JsonBody, err := tokenRequest("https://livro-digital-estante.prd.positivoon.com.br/v3/livros", "GET", token)
	if err != nil {
		return nil, err
	}

	LivrosParsados, err := ExtrairInfoLivros(JsonBody)
	if err != nil {
		return nil, err
	}

	return LivrosParsados, nil

	/* Exemplo de como ler []InfoLivros:
	bookInfos, err := pgo.ObterLivros(token)
	if err != nil {
		//cuide do erro
	}
	for _, book := range bookInfos {
		fmt.Println("Componente Curricular:", book.ComponenteCurricular)
		fmt.Println("Volume:", book.Volume)
		fmt.Println("Tipo:", book.Tipo)
		fmt.Println("URL:", book.URL)
		fmt.Println()
	}
	*/

}

func tokenRequest(url string, method string, token string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, errors.New("Não foi possível criar a requesição:" + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		return nil, errors.New("Não foi possível enviar a requisão:" + err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Não foi possível ler a resposta:" + err.Error())
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Não foi possível fazer a autenticação!\nStatus HTTP: " + fmt.Sprint(res.StatusCode))
	}

	res.Body.Close()
	return body, nil
}

func ExtrairInfoLivros(jsonData []byte) ([]InfoLivro, error) {
	var books []InfoLivro

	var rawBooks []map[string]interface{}
	err := json.Unmarshal(jsonData, &rawBooks)
	if err != nil {
		return nil, err
	}

	for _, rawBook := range rawBooks {
		arquivos := rawBook["arquivos"].([]interface{})
		if len(arquivos) > 0 {
			arquivo := arquivos[0].(map[string]interface{})
			book := InfoLivro{
				ComponenteCurricular: rawBook["componenteCurricular"].(string),
				Volume:               rawBook["volume"].(string),
				Tipo:                 arquivo["tipo"].(string),
				URL:                  arquivo["caminho"].(string),
			}
			books = append(books, book)
		}
	}

	return books, nil
}
