# phinvads-fhir

PHIN VADS written in Go. Load `phinvads.dump` into a PostgreSQL database using `pg_restore`, run the app with [air](https://github.com/air-verse/air), and go!

### Dev setup

1. Clone this repo:

    ```bash
    git clone https://github.com/CDCgov/phinvads-fhir.git
    cd phinvads-fhir
    ```

1. Install and run [PostgreSQL](https://www.postgresql.org/download/), or just run `docker compose up`
1. Create an empty database:

    ```bash
    psql -c 'CREATE DATABASE phinvads'
    ```

1. Load the database dump file:

    ```bash
    pg_restore -d phinvads --no-owner --role=$(whoami) phinvads.dump
    ```

1. Install [Go](https://go.dev/doc/install)
1. Install [air](https://github.com/air-verse/air):

    ```bash
    go install github.com/air-verse/air@latest
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
