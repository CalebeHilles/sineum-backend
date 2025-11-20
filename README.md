# Sineum Backend API

REST API for blog management built with Golang and PostgreSQL.
This API uses GORM as ORM and the database is hosted on Neon DB.

## Base URL

**Local:** `http://localhost:8000`  
**Production:** `[https://seu-dominio.railway.app](https://sineum-backend-production.up.railway.app)`

## Endpoints

| Method | Route | Description |
|--------|-------|-------------|
| `GET` | `/blogs` | List all blogs |
| `GET` | `/blogs/{id}` | Get blog by ID |
| `POST` | `/blogs` | Create a new blog |
| `PUT` | `/blogs/{id}` | Update a blog |
| `DELETE` | `/blogs/{id}` | Delete a blog |

## Blog Structure

| Field | Type | Limit |
|-------|------|-------|
| ID | serial | - |
| Title | varchar | 200 characters |
| Description | text | - |
| Content | text | - |

## Upcoming Features

- Authentication
- Categories
- Social features (follow, likes, comments)
