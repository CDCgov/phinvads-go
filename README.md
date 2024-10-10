# phinvads-go

PHIN VADS written in Go. Load `phinvads.dump` into a PostgreSQL database using `pg_restore`, run the app with [air](https://github.com/air-verse/air), and go!

## Getting Started

Clone down the repository:

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

In this section, we'll:

1. Start PostgreSQL in a container or locally
2. Load data into PostgreSQL
3. Start the app

#### Running PostgreSQL in a Container

You may choose to manage your database using Docker. `phinvads-go` provides a [Docker Compose](https://docs.docker.com/compose/) file to allow for a streamlined database setup. Please download and install [Rancher Desktop](https://rancherdesktop.io/) if you don't already have it installed.

Note: If you'd prefer to install PostgreSQL locally, please skip the rest of this section and refer to the [Local Database Setup](#local-database-setup) section instead.

To start your database, run this command:

`make startdb`

Next, please follow these steps in order to load in the data:

1. Place `phinvads.dump` into the top level of your `/phinvads-go` project directory (Please let an engineering team member know if you need a copy of `phinvads.dump`)
2. Navigate to the project directory and start your PostgreSQL database with `make startdb` (if the container isn't already running)
3. Run `make refreshdb` to create the `phinvads` database and load in the data

#### Running the App

In order to run the app you'll need to have completed the following steps:

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

Once all dependencies have been installed, run this command:

```bash
make startapp
```

The application is now running and available in your browser at [http://localhost:4000](http://localhost:4000/). As files are modified and saved, your application will automatically rebuild.

If you'd like to use live reloading in your browser when doing frontend work, you can make use of the templ proxy. The proxy can be accessed in your browser at [http://localhost:7331](http://localhost:7331). Changes to `*.templ` files will be automatically reflected in the browser.

You can stop running the app with `ctrl + c`.

### Local Database Setup

If you'd prefer to not rely on Docker and instead run PostgreSQL locally, you will need to:

1. Install and run [PostgreSQL](https://www.postgresql.org/download/)
1. Create an empty database:

    ```bash
    psql -c 'CREATE DATABASE phinvads'
    ```

1. Load the database dump file:

    ```bash
    pg_restore -d phinvads --no-owner --role=$(whoami) phinvads.dump
    ```
    
### Linting

Pull requests to the `main` branch are required to pass linting checks. To run these locally, [install `golangci-lint`](https://golangci-lint.run/welcome/install/#local-installation). Then you can `golangci-lint run` from the project directory.
