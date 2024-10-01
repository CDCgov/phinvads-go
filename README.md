# phinvads-go

PHIN VADS written in Go. Load `phinvads.dump` into a PostgreSQL database using `pg_restore`, run the app with [air](https://github.com/air-verse/air), and go!

## Dev setup

Clone the repo:

```bash
git clone https://github.com/CDCgov/phinvads-go.git
cd phinvads-go
```

Create a `.env` file:

```bash
cp .env.sample .env
```

### direnv setup and configuration

1. [Install direnv (`brew install direnv`)](https://direnv.net/docs/installation.html)
2. [Add a hook for your shell](https://direnv.net/docs/hook.html)
3. Restart your terminal
4. Run `direnv allow` in the project directory

### Database setup

#### Running Postgres in Docker

Download and install [Docker](https://www.docker.com/products/docker-desktop/) if you don't already have it installed.

1. Place `phinvads.dump` into the top level of your project directory (`/phinvads-go`)
2. Navigate to the project directory and start your PostgreSQL database with `make start`
3. Run `make refresh` to create the `phinvads` database and load in the data

When you want to shut down your database environment, run `make stop`.

#### Running Postgres locally

1. Install and run [PostgreSQL](https://www.postgresql.org/download/)
1. Create an empty database:

    ```bash
    psql -c 'CREATE DATABASE phinvads'
    ```

1. Load the database dump file:

    ```bash
    pg_restore -d phinvads --no-owner --role=$(whoami) phinvads.dump
    ```

### Application setup

1. Install [Go](https://go.dev/doc/install)
1. Install [air](https://github.com/air-verse/air):

    ```bash
    go install github.com/air-verse/air@latest
    ```

1. Install [templ](https://github.com/a-h/templ)

    ```bash
    go install github.com/a-h/templ/cmd/templ@latest
    ```

1. Install [mkcert](https://github.com/FiloSottile/mkcert)
1. Create your own self-signed certs for local development:  

    ```bash
    cd tls
    mkcert -install
    mkcert localhost
    cd ..
    ```

1. Run the app! If you are only working on backend code, you can just run a simple live reload with air:

    ```bash
    air
    ```

1. Air will also work for the frontend, but you will have to refresh your browser every time you make a change. To get automatic browser reloads, run the app this way:

    ```bash
    templ generate --watch --proxy="http://localhost:4000"
    # Then, in a separate terminal window, run air:
    air -c .air-with-proxy.toml
    ```
