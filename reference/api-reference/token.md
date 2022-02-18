# Token

## Pegando sua token

Logo após obter sua token, talvez você queira obter as informações do usuário, veja abaixo para saber mais!

{% content-ref url="pets.md" %}
[pets.md](pets.md)
{% endcontent-ref %}

{% swagger baseUrl="https://sso.specomunica.com.br/connect" method="get" path="/token" summary="Obter a token" %}
{% swagger-description %}
Retorna sua token, a expiração dela, a 

`refresh_token`

 e o tipo da token.

\


É necessário que o 

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

{% swagger-response status="400: Bad Request" description="Algo está errado, verifique os parâmetros" %}

{% endswagger-response %}
{% endswagger %}
