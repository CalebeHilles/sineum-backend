# Sineum Backend API

REST API for blog management built with Golang and PostgreSQL.
This API uses GORM as ORM and the database is hosted on Neon DB.

## Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| `GET` | `/blogs` | Lista todos os blogs |
| `GET` | `/blogs/{id}` | Lista um blog pelo ID |
| `POST` | `/blogs` | Posta um blog |
| `PUT` | `/blogs/{id}` | Edita um blog |
| `DELETE` | `/blogs/{id}` | Deleta um blog |

## Campos de cada post

| Campo | Tipo | Limite |
|-------|------|--------|
| ID | serial | - |
| Título | varchar | 200 caracteres |
| Descrição | text | - |
| Corpo do texto | text | - |

## Futuras features palpáveis

- Autenticação
- Categorias

## Futuras features

- Funções sociais (seguir, like, comentários)
