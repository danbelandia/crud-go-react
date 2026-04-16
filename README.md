# 👥 CRUD de Usuarios - Full Stack (Go + React)

Un sistema integral de gestión de usuarios desarrollado con una arquitectura limpia. Este proyecto demuestra la integración de una API RESTful construida en Go con una interfaz dinámica en React, utilizando SQLite como motor de base de datos ligero.

## 🚀 Tecnologías Utilizadas

### Backend
* **Lenguaje:** Go (Golang)
* **Base de Datos:** SQLite (`modernc.org/sqlite`)
* **Arquitectura:** Hexagonal / Puertos y Adaptadores
* **Características:** * 
  * Middlewares personalizados (CORS, Autenticación por roles)
  * Paginación nativa

### Frontend
* **Librería:** React 18
* **Herramienta de Construcción:** Vite
* **Estilos:** CSS
* **Características:**
  * Renderizado condicional
  * Manejo de estado centralizado
  * Componentización de UI (Formulario y Tabla separados)

---

## ⚙️ Requisitos Previos

Asegúrate de tener instalado en tu máquina local:
* [Go](https://go.dev/dl/) (v1.21 o superior)
* [Node.js](https://nodejs.org/) (v18 o superior) y npm

---

## 🛠️ Instalación y Ejecución Local

Para correr este proyecto en tu entorno local, necesitas levantar ambos servidores (Backend y Frontend) en terminales separadas.

### 1. Levantar el Backend (Go)
La base de datos SQLite se creará automáticamente en la raíz al ejecutar el servidor por primera vez.
```bash
# Navega a la raíz del proyecto
cd crud-stefanini

# Ejecuta el servidor principal
go run main.go
