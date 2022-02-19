# Início Rápido

## Pegue a sua chave da API

Suas solicitações da API são autenticadas usando a chave da API. Qualquer requisição que não inclua uma chave da API retornará um erro.

Você pode pegar a chave da API de uma maneira simples, apenas faça uma requisição como explicado abaixo:



{% swagger baseUrl="https://sso.specomunica.com.br" method="get" path="/connect/token" summary="Obter a chave da API" %}
{% swagger-description %}
Envie sua requisição para obter a sua token (ou chave da API). Note que é necessário ter um cadastro no Positivo On para fazer isso.

\


Também é necessário que o 

`Content-Type`

seja 

`application/x-www-form-urlencoded`

 para os parâmetros abaixo.
{% endswagger-description %}

{% swagger-parameter in="body" name="username" required="true" type="string" %}
Seu usuário do Positivo On
{% endswagger-parameter %}

{% swagger-parameter in="body" name="password" required="true" type="string" %}
A sua senha do Positivo On
{% endswagger-parameter %}

{% swagger-parameter in="body" name="grant_type" required="true" type="string" %}
O valor desse parâmetro é 

`password`

.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="client_id" required="true" type="string" %}
O valor desse parâmetro é 

`hubpsd`

.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="client_secret" type="string" required="true" %}
O valor desse parâmetro é 

`DA5730D8-90FF-4A41-BFED-147B8E0E2A08`

.
{% endswagger-parameter %}

{% swagger-parameter in="body" name="scope" type="string" required="true" %}
O valor desse parâmetro é 

`openid%20offline_access%20integration_info`

.
{% endswagger-parameter %}

{% swagger-response status="200" description="Login realizado com sucesso!" %}
```javascript
{
    "access_token": "Uma token muito longa",
    "expires_in": 43200,
    "token_type": "Bearer",
    "refresh_token": "Uma outra token, porem menor",
    "scope": "integration_info offline_access openid"
}
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="Algo está errado, verifique se os parâmetros estão corretos" %}

{% endswagger-response %}
{% endswagger %}

### Obtendo as informações do usuário

Talvez você também queira pegar as informações do usuário, elas serão bem importantes em quase todas as outras requisições que for fazer

{% swagger baseUrl="https://sso.specomunica.com.br" method="get" path="/userinfo" summary="Obter as informações do usuário" %}
{% swagger-description %}
Sua token será usada aqui. Nos próximos exemplos ela não será mais lembrada, mas lembre-se que ela sempre vai ser necessária para qualquer outra requisição.
{% endswagger-description %}

{% swagger-parameter in="header" name="Bearer " required="true" type="string" %}
Sua token vem aqui
{% endswagger-parameter %}

{% swagger-response status="200" description="Tudo está ok!" %}
```javascript
{
    "sub": "ID Único",
    "auth_time": 1645217796,
    "idp": "local",
    "name": "Rebecca Silva",
    "username": "rs2005",
    "email": "",
    "integration_id": "ID da integração",
    "amr": "pwd",
    "schools": "{\"id\":\"ID único da escola\",\"integration_id\":\"ID da integração (escola)\",\"user_id\":\"ID da escola\",\"name\":\"Nome de sua escola\",\"roles\":[\"ALUNO\"],\"time_zone\":\"E. South America Standard Time\",\"url\":\"psdXXXX.specomunica.com.br\"}"
}
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="Algo está errado, verifique a token" %}

{% endswagger-response %}
{% endswagger %}

## Instalando a biblioteca

Talvez uma das melhores maneiras de interagir com a API é com a biblioteca `pgo`. Atualmente ela so está disponível em Golang.

{% tabs %}
{% tab title="Go" %}
```
# Instale via go get
go get -u github.com/alternativeon/pgo
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
**Precisamos de você!** Se souber programar em outra linguagem, você pode nos ajudar criando uma biblioteca para a API do Positivo.
{% endhint %}

## Faça sua primeira requisição

Para fazer sua primeira requisição, enviar uma solicitação autenticada para o endpoint `NivelEnsino`. Isso vai lhe mostrar em qual turma o usuário está.

{% swagger method="get" path="/api/NivelEnsino" baseUrl="https://apihub.positivoon.com.br" summary="Obtêm as turma em que o usuário está matriculado." %}
{% swagger-description %}
Essa request retorna em qual turma e série o usuário está.
{% endswagger-description %}

{% swagger-parameter in="path" name="usuarioid" type="int" required="true" %}
Esse é o ID do seu usuário, não sabe como obter? Olhe na sessão "Informações do Usuário".
{% endswagger-parameter %}

{% swagger-response status="200: OK" description="A informação foi obtida com sucesso." %}
```javascript
[
    {
        "value": "EF1",
        "label": "Ensino Fundamental Anos Iniciais",
        "turmas": [
            {
                "nomeTurma": "11502",
                "turmaValida": true,
                "nomeSerie": "5º ano"
            }
        ]
    }
]
```
{% endswagger-response %}

{% swagger-response status="400: Bad Request" description="Verifique o ID do usuário" %}
```javascript
{
    "type": "https://tools.ietf.org/html/rfc7231#section-6.5.1",
    "title": "One or more validation errors occurred.",
    "status": 400,
    "traceId": "00-19b58062ef670a45b0d2829a85a1899e-8dc602620e03f34a-00",
    "errors": {
        "usuarioId": [
            "The value '' is invalid."
        ]
    }
}
```
{% endswagger-response %}

{% swagger-response status="401: Unauthorized" description="Verifique se a token está presente." %}

{% endswagger-response %}
{% endswagger %}

Dê uma olhada em como você pode usar esse método usando nossas biblioteca, ou via `curl`:

{% tabs %}
{% tab title="curl" %}
```
curl --location --request GET 'https://apihub.positivoon.com.br/api/NivelEnsino?usuarioID=XXXX' \
--header 'Authorization: Bearer sua_token_aqui 
```
{% endtab %}

{% tab title="Go" %}
```go
// Lembre-se de importar o modulo com import "github.com/alternativeon/pgo"
token, err := pgo.Login("seu usuário", "sua senha")
	if err != nil {
		panic(err)
	}
resultado, err := pgo.GetClass(token) //Retorna em json a turma do usuário
	if err != nil { //Essa função também vai fazer uma requisição automaticamente para obter o id do usuário. 
		panic(err)
	}
fmt.Println(resultado)
```
{% endtab %}
{% endtabs %}
