# hello-go-web

Um projeto simples de API web em Go para demonstração e testes de containerização com Docker.

## Sobre o projeto

Este projeto consiste em uma API web minimalista escrita em Go que retorna uma mensagem JSON simples. É ideal para demonstrar conceitos de:

- Implementação de API web simples em Go
- Estrutura básica de projetos Go
- Implementação de APIs REST simples em Go

## Estrutura do projeto

```
hello-go-web/
├── src/                # Diretório de código-fonte
│   ├── main.go         # Código-fonte principal da aplicação
│   ├── go.mod          # Definição de módulo Go
│   └── go.sum          # Checksums das dependências
└── README.md           # Este arquivo
```

## Requisitos

- Go 1.22 ou superior

## Como executar localmente

1. Clone o repositório:
   ```
   git clone https://github.com/yourusername/hello-go-web.git
   cd hello-go-web
   ```

2. Entre no diretório src:
   ```
   cd src
   ```

3. Execute a aplicação localmente:
   ```
   go run main.go
   ```

4. Acesse a API:
   ```
   curl http://localhost:8080
   ```
   
   Resposta esperada:
   ```json
   {"message":"Hello Go"}
   ```



