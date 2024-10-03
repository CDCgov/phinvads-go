# phinvads-go

PHIN VADS written in Go. Load `phinvads.dump` into a PostgreSQL database using `pg_restore`, run the app with [air](https://github.com/air-verse/air), and go!

## Getting Started

Firstly, clone down the repository:

```bash
git clone https://github.com/CDCgov/phinvads-go.git
cd phinvads-go
```

Create a `.env` file:

```bash
cp .env.sample .env
```

### Installing and Configuring `direnv`

We use [`direnv`](https://direnv.net/) to automatically set environment variables on a per-project basis. The `.envrc` file found in the project's directory tells `direnv` to read in the environment variables specified in the `.env` file.

Follow these steps to set up `direnv` on your machine:

1. [Install direnv (`brew install direnv`)](https://direnv.net/docs/installation.html)
2. [Add a hook for your shell](https://direnv.net/docs/hook.html)
3. Restart your terminal
4. Run `direnv allow` in the project directory

When navigating into the `phinvads-go` project directory, you should see something like this in your terminal, which lets you know that the environment variables are loaded:

```bash
direnv: export +DB_HOST +DB_NAME +DB_PASSWORD +DB_PORT +DB_USER +HOST +PORT
```

### Local Development

There are two ways to run `phinvads-go`: with Docker or without Docker. This section will provide instructions to run the project using either method.

#### Running the App With Docker

`phinvads-go` makes use of [Docker Compose](https://docs.docker.com/compose/) to allow for a streamlined development setup. Please download and install [Docker](https://www.docker.com/products/docker-desktop/) if you don't already have it installed.

##### Note About Live Reloading

The Docker Compose setup does not support [live reloading](https://templ.guide/commands-and-tools/live-reload-with-other-tools) (automatic browser refreshing) when static files are changed. If you are doing frontend work and would like this feature, it may be preferable to use the `make startlocal` command to run the app locally. If not, feel free to skip this section and move on to [Running the App](#running-the-app).

It's possible to use both the `make startlocal` command together with the Dockerized database if you'd still like to manage your database using Docker.

```bash
docker compose up -d db
make startlocal
```

The application will be available in your browser at [http://localhost:7331](http://localhost:7331/) with live reloading enabled.

##### Running the App

With Docker installed, we can start the application:

```bash
make start
```

The application is now running and available in your browser at [http://localhost:4000](http://localhost:4000/). Next, we will need to load the data set into the database.

##### Loading Data into PostgreSQL

While the app is accessible, it won't have all of the data it needs to function correctly until we load the data into the PostgreSQL database.

Please follow these steps in order to load in the data:

1. Place `phinvads.dump` into the top level of your `/phinvads-go` project directory (Please let an engineering team member know if you need a copy of `phinvads.dump`)
2. Navigate to the project directory and start your PostgreSQL database with `make start` (if the application isn't already running)
3. Run `make refresh` to create the `phinvads` database and load in the data

##### Shutting Down the App

Running `make stop` will shut down all running containers. The data loaded into your database will persist next time you run `make start`.

#### Running the App Without Docker

`phinvads-go` can also run locally outside of the Docker Compose environment. Please follow the steps listed here if you'd like to run the application without Docker:

##### Database Setup

1. Install and run [PostgreSQL](https://www.postgresql.org/download/)
1. Create an empty database:

    ```bash
    psql -c 'CREATE DATABASE phinvads'
    ```

1. Load the database dump file:

    ```bash
    pg_restore -d phinvads --no-owner --role=$(whoami) phinvads.dump
    ```

##### Application Setup

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
    make startlocal
    ```
