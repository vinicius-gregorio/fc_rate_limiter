# FULL CYCLE - RATE LIMITER

Um middleware de rate limiter desenvolvido em Go para controlar o tráfego de requisições em um serviço web. Ele suporta a limitação de requisições por segundo com base no IP do cliente ou em um token de acesso. O rate limiter utiliza Redis para armazenar e gerenciar os dados de limitação e pode ser facilmente integrado a qualquer aplicação web em Go.


## Requisitos

- [X] O rate limiter deve poder trabalhar como um middleware que é injetado ao servidor web
- [X] O rate limiter deve permitir a configuração do número máximo de requisições permitidas por segundo.
- [X] O rate limiter deve ter ter a opção de escolher o tempo de bloqueio do IP ou do Token caso a quantidade de requisições tenha sido excedida.
- [X] As configurações de limite devem ser realizadas via variáveis de ambiente ou em um arquivo “.env” na pasta raiz.
- [X] Deve ser possível configurar o rate limiter tanto para limitação por IP quanto por token de acesso.
- [X] O sistema deve responder adequadamente quando o limite é excedido:
Código HTTP: 429
- [X] Mensagem: you have reached the maximum number of requests or actions allowed within a certain time frame
- [X] Todas as informações de "limiter” devem ser armazenadas e consultadas de um banco de dados Redis. Você pode utilizar docker-compose para subir o Redis.
- [X] Crie uma “strategy” que permita trocar facilmente o Redis por outro mecanismo de persistência.
- [X] A lógica do limiter deve estar separada do middleware

## Funcionalidades

- Limitação Baseada em IP: Restringe o número de requisições de um único endereço IP dentro de uma janela de tempo definida.
- Limitação Baseada em Token: Aplica limites personalizados com base em um token de acesso enviado no cabeçalho da requisição.
- Configuração por Variáveis de Ambiente:
- Limites de requisição (por IP e por token).
- Duração do bloqueio após exceder os limites.
- Configurações de conexão com o Redis.
- Respostas Claras:
    Código HTTP: 429 Too Many Requests.
    Mensagem: Você atingiu o número máximo de requisições ou ações permitidas dentro de um determinado período.
- Armazenamento no Redis:
    Todos os dados de limitação são armazenados e gerenciados no Redis.
    Estratégia de Armazenamento Pluggável: Possibilidade de trocar o Redis por outro mecanismo de persistência.


## Configuração
A configuração do aplicativo é feita por variáveis de ambiente. Essas variáveis podem ser definidas no arquivo docker-compose.yml ou em um arquivo .env.



Variável	Descrição	Valor Padrão
REDIS_HOST	Endereço do Redis	localhost
REDIS_PORT	Porta do Redis	6379
REDIS_PASSWORD	Senha do Redis	123
RATE_LIMITER_IP_LIMIT	Número máximo de requisições por segundo por IP	10
RATE_LIMITER_TOKEN_LIMIT	Número máximo de requisições por segundo por token	100
RATE_LIMITER_BLOCK_DURATION	Duração do bloqueio (em segundos) após exceder o limite	300


## Como funciona
### Integração do Middleware

O rate limiter é implementado como middleware que:

1. Extrai o API_KEY do cabeçalho da requisição (se presente).
2. Utiliza o IP do cliente caso nenhum token seja fornecido.
3. Verifica o limite de requisições usando o RedisRateLimiter:
    - Se o limite for excedido, responde com HTTP 429.
    - Se permitido, encaminha a requisição para o próximo handler.



## Como usar
### Instalação
1. ` $ go get github.com/vinicius-gregorio/fc_rate_limiter` 
2. Faça o devido import:
```go
import (
    "github.com/vinicius-gregorio/fc_rate_limiter/limiter"
)
``` 

### Exemplos
1. Confira o exemplo em `examples/main.go` 
2. Confira o `docker-compose.yml` para variaveis de ambiente.


## Full Cycle:

### Como executar o projeto?
1. `docker compose up -d` 
2. Faça requisições para `localhost:8080` 