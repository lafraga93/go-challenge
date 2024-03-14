
### Desafio Golang

##### Implementar um endpoint HTTP GET responsável por entregar uma lista de motoristas com seus respectivos dados bancários.

###### Contrato de integração (response do endpoint criado):

```
[
    ...
    {
        "uuid": "082b918b-1fe3-4bb4-a928-50a2cd4234b2",
        "name": "Renata Silveira",
        "banking": {
            "bank": "Banco do Brasil",
            "agency": "52320",
            "account_number": "10524-6"
        }
    },
    ...
]
```

###### Endpoints externos que serão utilizados:

* Lista de motoristas: `GET` `https://go-challenge-truckers.free.beeceptor.com/api/v1/`

* Dados bancários de motoristas: `GET` `https://go-challenge-banking.free.beeceptor.com/api/v1/trucker/{uuid}`
