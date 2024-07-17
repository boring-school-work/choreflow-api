# Choreflow API

## Run the API Locally

### Prerequisites

- Install Go v1.22.1+

### Install Dependencies

Navigate to the project direcotory and run the following command:

```bash
go mod tidy
```

### Create Database

1. Run/source/import the sql file at `db/chores_mgt.sql` on your sql database.
2. Update your environment variables to reflect your sql configuration.

List of environment variables to create:

- `DBAPIHOST`: MYSQL host (e.g. localhost).
- `DBAPIPORT`: MYSQL port (e.g. 3306).
- `DBAPIUSER`: MYSQL user that has access to the database.
- `DBAPIPASS`: MYSQL user password.
- `DBAPINAME`: Must be `chores_mgt`.

### Run Project

Navigate to the project directory and run the command below:

```bash
go run .
```

Visit http://localhost:3211/choreflow/api/v1/swagger/ui to test the api.



