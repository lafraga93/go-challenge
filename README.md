
### Desafio Golang

##### O objetivo desse desafio é propor e realizar melhorias no programa exemplo contido nesse repositório.

##### O programa é muito simples e basicamente expõe um endpoint GET responsável por entregar uma lista de motoristas com seus respectivos dados bancários. Os motoristas e seus respectivos dados bancários são obtidos a partir de outras apis.

##### Na refatoração, serão considerados os seguintes pontos:

* Boas práticas de codificação
* Organização, arquitetura e design de código (camadas, arquitetura hexagonal, princípios SOLID, DRY, KISS, etc)
* Tratamento de erros
* Testes de unidade
* Performance da aplicação

##### Contratos e referência técnica:

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
