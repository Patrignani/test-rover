# test-rover

Esta aplicação foi desenvolvida para simular envios de comandos para rovers por meio de um arquivo de texto.

## 📋 Arquitetura

O projeto monorepo contém dois pontos de entrada (API e prompt), podendo incluir outros (gRPC, RabbitMQ, Cloud Functions, etc.). A camada core corresponde aos objetos e regras de negócio da aplicação.

Neste exemplo, está sendo utilizado apenas um arquivo de CI/CD do GitHub, mas, se houver a necessidade de deploy, cada projeto de entrada terá seu próprio YAML.

## 📋 Informações sobre o arquivo de texto

O arquivo de texto deve conter, no mínimo, 3 linhas e sempre ter um número ímpar de linhas. A primeira linha representa as coordenadas do Planalto onde o rover está localizado. A segunda e terceira linha são informações sobre o rover, que podem se repetir N vezes para cada rover. Exemplo:

Linha 1) 5 5 > Eixos do planalto atual

Linha 2) 1 2 N > Dois primeiros números separados por espaço são os eixos atuais do rover, a sigla representa a coordenada

Linha 3) MRL > Essa linha representa a instrução passada para o rover

## 📋 Informações sobre as coordenadas

As coordenadas serão representadas por siglas

```
	North = "N"
	East  = "E"
	South = "S"
	West  = "W"
```
## 📋 Informações sobre as instruções

As instruções são representadas por siglas

```
  Left  = "L"
  Right = "R"
  Move  = "M"
```
### 📄 Comandos das instruções

Left - Movimenta o rover a 90° para a esquerda.

Right - Movimenta o rover a 90° para a direita.

Move - Rover se movimenta para frente do eixo referente à coordenada.

Observação: o rover nunca ficará em um eixo negativo e nunca ficará em uma posição maior que o eixo passado do planalto.

### 📋 Pré-requisitos

Golang 1.20 ou superior

### 🔧 Instalação

https://go.dev/doc/install

Caso utilize o Visual Studio Code com a extensão do Go, siga os seguintes passos para utilizar múltiplos módulos:

Ctrl+Shift+P

Preferences: Open User Settings (Json)

Adicione o seguinte código no JSON:
```json
  "gopls": {
        "experimentalWorkspaceModule": true,
    }
```    

## ⚙️ Executando a API

Para executar a API, é necessário acessar o diretório correspondente:

```
cd ./api/
```

Em seguida, execute o seguinte comando para iniciar a API:

```
 go run main.go
```

Para testar a API, é necessário ter um arquivo de texto com as informações necessárias. A aplicação já contém um arquivo de teste na raiz.

Para enviar uma solicitação POST no Windows, é preciso estar no diretório raiz do projeto ou alterar o caminho do arquivo no comando:

```
Invoke-WebRequest -Uri http://localhost:8080/upload -Method POST -Body (Get-Content -Path "./commands_rovers.txt" -Raw)
```

No Linux, o comando para enviar uma solicitação POST é:

```
curl -X POST -d "$(cat ./commands_rovers.txt)" http://localhost:8080/upload
```

## ⚙️ Executando o Prompt

Para executar o prompt, é necessário acessar o diretório correspondente:

```
cd ./prompt/
```

Em seguida, execute o seguinte comando:

```
go run main.go commands_rovers.txt
```

## ⚙️ Executando os testes
Para executar os testes, é necessário acessar o diretório core:

```
cd ./core/
```

Em seguida, execute o seguinte comando:

```
go test ./...
```
