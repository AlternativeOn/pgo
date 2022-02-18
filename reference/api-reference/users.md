---
description: >-
  Para obter as soluções, você irá precisar de fazer duas requests, como a
  seguir
---

# Soluções

{% swagger method="get" path="?usuarioid=" baseUrl="https://apihub.positivoon.com.br/api/NivelEnsino" summary="" %}
{% swagger-description %}
Essa request retorna em qual turma e série o usuário está.
{% endswagger-description %}

{% swagger-parameter in="path" name="id" type="int" required="true" %}
Esse é o ID do seu usuário, pode ser obtido na página Userinfo
{% endswagger-parameter %}

{% swagger-parameter in="header" name="Bearer" type="string" required="true" %}
Sua token de acesso a plataforma.
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

{% swagger-response status="401: Unauthorized" description="" %}

{% endswagger-response %}
{% endswagger %}

## Creating users

{% swagger src="https://petstore.swagger.io/v2/swagger.json" path="/user/createWithList" method="post" %}
[https://petstore.swagger.io/v2/swagger.json](https://petstore.swagger.io/v2/swagger.json)
{% endswagger %}

{% swagger src="https://petstore.swagger.io/v2/swagger.json" path="/user/createWithArray" method="post" %}
[https://petstore.swagger.io/v2/swagger.json](https://petstore.swagger.io/v2/swagger.json)
{% endswagger %}
