<p align="center">
  <a href="https://gofiber.io">
    <img alt="Fiber" height="125" src=".github/logo.svg">
  </a>
  <br>
</p>
<p align="center">
  Esta é uma API feita em <a href="https://go.dev/">Go</a> como teste técnico para a posição de Backend Pleno na <a href="https://transfeera.com/">Transfeera</a>. A aplicação usa <a href="https://www.postgresql.org/">Postgresql</a> como banco de dados, testes unitários e de integração usando a biblioteca <a href="https://github.com/stretchr/testify">Testify</a>, containerização com <a href="https://www.docker.com/">Docker</a> e possui instruções de instalação, configuração e execução do projeto, além da documentação da API logo abaixo.
</p>

# ⚙️ Instalação

**É necessário ter Go, Docker e Docker Compose instalados na sua máquina.**

Clone o projeto para sua máquina:

```bash
git clone https://github.com/lbcosta/transfeera.backend.developer.test.git
```

Na raíz do projeto, crie um arquivo **.env** com os seguintes valores:

```
POSTGRES_HOST=db
POSTGRES_HOST_SEED=localhost
POSTGRES_PORT=5432
POSTGRES_USER=transfeera
POSTGRES_PASSWORD=transfeera.t3st
POSTGRES_DB=transfeera
```

# ⚡️ Inicialização

**Opção 1**

Para inicializar os testes e subir os containers, basta executar na raíz do projeto o script:

```bash
./scripts/run.sh
```

> ⚠️ Se os testes falharem, a execução irá ser cancelada. Se forem bem sucedidos, o script irá subir os containers necessários para o uso da aplicação e irá travar o terminal enquanto a aplicação estiver executando.

**Opção 2**

Para inicializar apenas os containers da aplicação _(e não rodar os testes)_, basta usar o **docker-compose**. Inicializar a aplicação desse modo, não irá travar o terminal:

```bash
docker-compose up --build -d
```

# 🌱 Seeding

Para popular o banco de dados que foi inicializado com a aplicação, basta rodar o seguinte script:

```bash
go run scripts/seed.go
```

> ⚠️ Provavelmente o terminal usado para subir a aplicação estará travado. Então é necessário abrir outro terminal para executar o script acima.

# 📃 Documentação da API

## 🔗 Endpoints

### 🔍 Listagem de Recebedores

**URL**: `/api/v1/beneficiaries[?filter=X][&page=Y]` <br>
**Paramêtros**: `filter` - Filtro de busca, que pode ser um valor de _Status_ ("Rascunho" ou "Validado"), _Nome_, _Tipo de Chave PIX_ ("cpf", "cnpj", "email", "telefone", "chave\*aleatoria") ou o _Valor da chave PIX_ <br>
**Method**: GET<br>
**Request Body**: _Sem request_ <br>
**Response**: Objeto JSON contendo metadados sobre a busca e os dados buscados<br>

Exemplo de **Response**:

<img src="https://img.shields.io/badge/Status-200-green">

```json
{
  "status": "success",
  "code": 200,
  "metadata": {
    "total_count": 39,
    "page": 1,
    "per_page": 10,
    "total_pages": 4
  },
  "data": [
    {
      "status": "Validado",
      "name": "John Doe",
      "document_number": "12345678907",
      "email": "johndoe@example.com",
      "pix_key_type": "CPF",
      "pix_key_value": "12345678907",
      "bank": "ABC Bank",
      "agency": "1234",
      "account": "56789"
    },
    {
      "status": "Validado",
      "name": "Jane Doe",
      "document_number": "98765432107",
      "email": "janedoe@example.com",
      "pix_key_type": "CPF",
      "pix_key_value": "98765432107",
      "bank": "XYZ Bank",
      "agency": "5678",
      "account": "12345"
    }
    // ...
  ]
}
```

Exemplos de possíveis **Erros**:

<img src="https://img.shields.io/badge/Status-400-red">

```json
{
  "status": "invalid_input",
  "code": 400,
  "error": "The requested page does not exist."
}
```

### ✏️ Criação de Novo Recebedor

**URL**: `/api/v1/beneficiaries` <br>
**Method**: POST<br>
**Request Body**: Objeto JSON com informações do Recebedor<br>
**Response**: Objeto criado<br>

Exemplo de **Request Body**:

```json
{
  "name": "Leonardo Costa",
  "document_number": "04788380340",
  "email": "lbcosta.dev@gmail.com",
  "pix_key_type": "EMAIL",
  "pix_key_value": "lbcosta.dev@gmail.com"
}
```

Exemplo de **Response**:

<img src="https://img.shields.io/badge/Status-200-green">

```json
{
  "status": "Rascunho",
  "name": "Leonardo Costa",
  "document_number": "04788380340",
  "email": "lbcosta.dev@gmail.com",
  "pix_key_type": "EMAIL",
  "pix_key_value": "lbcosta.dev@gmail.com",
  "bank": "TransfeeraBank",
  "agency": "1234-5",
  "account": "987654-3"
}
```

Exemplos de possíveis **Erros**:

<img src="https://img.shields.io/badge/Status-400-red">

```json
{
  "status": "invalid_input",
  "code": 400,
  "error": "error on the following fields: Email"
}
```

### 🔄 Edição de Recebedores

**URL**: `/api/v1/beneficiaries/:id` <br>
**Paramêtros**: `id` - Id do recebedor a ser editado<br>
**Method**: PATCH<br>
**Request Body**: Objeto JSON com as informações a serem editadas<br>
**Response**: Objeto após ser editado<br>

Exemplo de **Request Body**:

```json
{
  "document_number": "04788380241",
  "email": "lbcosta@gmail.com"
}
```

Exemplo de **Response**:

<img src="https://img.shields.io/badge/Status-200-green">

```json
{
  "status": "Rascunho",
  "name": "Leonardo Costa",
  "document_number": "04788380241",
  "email": "lbcosta@gmail.com",
  "pix_key_type": "EMAIL",
  "pix_key_value": "lbcosta.dev@gmail.com",
  "bank": "TransfeeraBank",
  "agency": "1234-5",
  "account": "987654-3"
}
```

Exemplos de possíveis **Erros**:

<img src="https://img.shields.io/badge/Status-422-red">

```json
{
  "status": "error",
  "code": 422,
  "error": "record not found"
}
```

<img src="https://img.shields.io/badge/Status-422-red">

```json
{
  "status": "error",
  "code": 422,
  "error": "beneficiaries with Status=Validado should not update some fields"
}
```

### 🗑️ Exclusão de Recebedores

**URL**: `/api/v1/beneficiaries` <br>

**Method**: DELETE <br>

**Request Body**: Objeto JSON com uma lista dos IDs dos recebedores a serem excluídos <br>

**Response**: _Sem response - Status 204: No Content_ <br>

Exemplo de **Request Body**

```json
{
  "ids": [2, 3, 8, 12, 27]
}
```

Exemplos de possíveis **Erros**:

<img src="https://img.shields.io/badge/Status-422-red">

```json
{
  "status": "error",
  "code": 422,
  "error": "resource not found"
}
```
