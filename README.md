
# License Key Auth System

A license key system built with Vue.js and Golang



## Features

- Really Fast ⚡️
- Tailwind Admin
- GoFiber backend
- Ready to go Authentication


## Frontend Setup

Install my-project with npm

```sh
  cd frontend
  npm install
  npm run dev
```
    
## Backend Setup

Install my-project with npm

```sh
  cd backend
  go run . --development
```
    
## Environment Variables

To run this project locally, you will need to fill your `backend/.env.local` file with the following informatiom

`FRONTEND_URL` (default: `localhost:5173`)

`PORT` (default: `8090`)

`JWT_SECRET` (default: `my-secret`)

`POSTGRES_URL` (default: `postgresql://<username>:<password>@127.0.0.1:<port>/<db_name>`
