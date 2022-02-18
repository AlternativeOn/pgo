---
description: Aqui você obtêm todas as informações necessárias sobre o usuário.
---

# Userinfo

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
