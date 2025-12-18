# Greenlight API
This is a slightly modified implementation of the _Let's Go Further!_ book project. (by Alex Edwards)

It includes a variety of features such as:
  - CRUD Operations and JSON Error Handling
  - Input and Data Validation
  - SQL migration and PostgreSQL Query Timeouts
  - Optimistic Concurrency Control
  - Pagination, Sorting, Filtering, and Full-Text Search
  - IP-based Rate Limiting
  - Graceful Shutdown and expvar Metrics
  - User Activation Flow, Stateful Token Authentication, and Password Reset
  - Permission-based Authorization
  - CORS support

# Setup
## Local
### Requirements
- Go 1.25
- Make
- [Migrate](https://github.com/golang-migrate/migrate)
- Docker
#### 1. Dependencies
We can install dependencies using this command:
```bash
go mod download
```
#### 2. PostgreSQL
To simplify the PostgreSQL setup we can use docker: 
```bash
docker run --name greenlight-db -e POSTGRES_PASSWORD=example -e POSTGRES_USER=greenlight -e POSTGRES_DB=greenlight -p 5432:5432 -v greenlight-data:/var/lib/postgresql -d postgres
```
This command sets up a PostgreSQL instance that has the same DSN as the default one provided.

(This command is for PostgreSQL version 18 and above due to changes to [PGDATA variable](https://hub.docker.com/_/postgres#pgdata).)

#### 3. Run the Application
We use a Makefile to handle migrations and execution:
```bash
make run/api
```

## Production
We can configure an Ubuntu server (or most Debian-based distros) by executing the shell script located at `/remote/setup/01.sh` like:
```bash
bash /remote/setup/01.sh
```
At a high level the setup script does the following things:

- Update all packages on the server.

- Set the server time zone (in this case it will set it to `Europe/Helsinki`) and install support for locales.

- Create a `greenlight` user on the server, which we can use for day-to-day maintenance and for running our API application (rather than using the  `root` user account). We should also add the `greenlight` user to the `sudo` group, so that it can perform actions as `root` if necessary.

- Copy the `root` user’s `$HOME/.ssh` directory into the `greenlight` user's home directory. This will enable us to authenticate as the `greenlight` user using the same SSH key pair that we used to authenticate as the `root` user. We should also force the `greenlight` user to set a new password the first time they log in.

- Configure firewall settings to only permit traffic on ports `22` (SSH), `80` (HTTP) and `443` (HTTPS). We’ll also install [fail2ban](https://github.com/fail2ban/fail2ban) to automatically temporarily ban an IP address if it makes too many failed SSH login attempts.

- Install PostgreSQL. We’ll also create the `greenlight` database and user, and create a system-wide `GREENLIGHT_DB_DSN` environment variable for connecting to the database.

> [!Note] 
> But to simplify the PostgreSQL setup we can use docker instead: 
> ```bash
> docker run --name greenlight-db -e POSTGRES_PASSWORD=example -e POSTGRES_USER=greenlight -e POSTGRES_DB=greenlight -p 5432:5432 -v greenlight-data:/var/lib/postgresql -d postgres
> ```
> > This commands sets up a PostgreSQL instance that has the same DSN as the default one provided.
> > (This command is for PostgreSQL version 18 and above due to changes to [PGDATA variable](https://hub.docker.com/_/postgres#pgdata).)

- Install the `migrate` tool, using the [pre-built binaries](https://github.com/golang-migrate/migrate/releases) from GitHub.

- Install Caddy by following the [official installation instructions](https://caddyserver.com/docs/install#debian-ubuntu-raspbian) for Ubuntu.

- Reboot the server.

# Future plans:
- Add a Docker Compose configuration for easier deployment
- Generalize the table structure to support generic products beyond just movies.
- Build a modern dashboard using shadcn.
