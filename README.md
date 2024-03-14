
motoristas: GET trucker/api/v1

```
[
    {
        "uuid": "",
        "name": "",
        "avatar": "",
    }
]
```

dados bancários: GET baking/api/v1/{trucker_uuid}

```
{
    "bank": "",
    "agency": "",
    "account": "",
}
```

contrato integração bff:

```
[
    {
        "uuid": "",
        "name": "",
        "avatar": "",
        "banking": {
            "bank": "",
            "agency": "",
            "account": "",
        }
    }
]
```

Expor um endpoint capaz de entragar uma lista de motoristas respeitando o seguinte contrato:

- contrato integração bff

Fluxo:

1. Obter a lista de motoristas
2. Obter dados bancários de cada motoristas
3. Entregar a listagem de motoristas no endpoint criado, respeitando o contrato previamente definido
