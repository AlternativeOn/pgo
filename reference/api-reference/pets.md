---
description: Aqui você obtêm todas as informações necessárias sobre o usuário.
---

# Userinfo

## Obtendo as informações

{% swagger baseUrl="https://sso.specomunica.com.br/connect" method="get" path="/userinfo" summary="Obter as informações do usuário" %}
{% swagger-description %}
Sua token vai ser usada aqui.
{% endswagger-description %}

{% swagger-parameter in="header" name="Bearer " required="true" type="string" %}
SUA_TOKEN
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

{% hint style="info" %}
**Good to know:** This API method was created using the API Method block, it's how you can build out an API method documentation from scratch. Have a play with the block and you'll see you can do some nifty things like add and reorder parameters, document responses, and give your methods detailed descriptions.
{% endhint %}

## Updating a pet

{% swagger src="https://petstore.swagger.io/v2/swagger.json" path="/pet" method="put" %}
[https://petstore.swagger.io/v2/swagger.json](https://petstore.swagger.io/v2/swagger.json)
{% endswagger %}

{% hint style="info" %}
**Good to know:** This API method was auto-generated from an example Swagger file. You'll see that it's not editable – that's because the contents are synced to an URL! Any time the linked file changes, the documentation will change too.
{% endhint %}
