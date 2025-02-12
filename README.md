# SalesForge API

## Overview

SalesForge API is a service for managing email sequences and steps within those sequences. It provides functionality to add, update, and delete sequences and steps.

## Prerequisites

- Go 1.22.5+
- PostgreSQL
- Docker

## Project Structure

- `internal/models`: Contains the data models.
- `internal/persistence`: Contains the repository layer for database interactions.
- `internal/service`: Contains the service layer for business logic.
- `internal/psql`: Contains the PostgreSQL connection setup.
- `config`: Contains configuration files.

## Database Setup

1. Ensure PostgreSQL is installed and running.
2. Create a database for the project.
3. Run the `init.sql` script to create the necessary tables and prepopulate them with sample data.

## Configuration

Create a configuration file `config.yaml` inside `config` directory with the following structure:

```yaml
ServiceName: "salesforge-api"
Environment: "dev"
Server:
  AppServerPort: 8080
  HealthcheckPort: 8081
  JWTAuthentication: false
Psql:
  Db: "postgres"
  User: "yourusername"
  Pass: "yourpassword"
  Host: "localhost"
  Port: 5432
TestDB: #For running functional tests
  Db: "postgres"
  User: "yourusername"
  Pass: "yourpassword"
  Host: "localhost"
  Port: 5433
Logger:
  Level: "info" #debug
  Format: "json" #console
```

## Running the Service
To run the SalesForge API project, follow these steps:  

1. Clone the repository:  
```git clone ```
2. Navifate to its directory:  
```cd salesforge-api```
3. Build and run the project using the Makefile:  
`make`

## Running Tests
To run the tests, use `make test` command:
```bash
make test
```
You will need test database credentials in the configuration file.

### API Endpoints

#### Add Sequence

- **Endpoint**: `/v1/sequence`
- **Method**: `POST`
- **Payload**:
  ```json
  {
    "account_id": 6789,
    "sequence_name": "New Welcome Sequence!",
    "sequence_open_tracking_enabled": true,
    "sequence_click_tracking_enabled": false,
    "steps": [
        {
            "step_email_subject": "Welcome to our service",
            "step_email_body": "Thank you for joining us!",
            "wait_days": 1,
            "eligible_start_time": 1737621878,
            "eligible_end_time": 1737631081
        },
        {
            "step_email_subject": "Getting Started",
            "step_email_body": "Here are some tips to get started.",
            "wait_days": 2,
            "eligible_start_time": 1737751081,
            "eligible_end_time": 1737791222
        }
    ]
  }
  ```

#### Update Sequence

- **Endpoint**: `/v1/sequence`
- **Method**: `PUT`
- **Payload**:
  ```json
  {
    "account_id": 6789,
    "sequence_id": 2,
    "sequence_open_tracking_enabled": false,
    "sequence_click_tracking_enabled": true
  }
  ```

#### Update Step

- **Endpoint**: `/v1/step`
- **Method**: `PUT`
- **Payload**:
  ```json
  {
    "account_id": 1,
    "step_id": 1,
    "sequence_id": 1,
    "step_email_subject": "Welcome to our service!",
    "step_email_body": "Thank you for joining us."
  }
  ```

#### Delete Step

- **Endpoint**: `/v1/step`
- **Method**: `DELETE`
- **Payload**:
  ```json
  {
    "account_id": 12345,
    "step_id": 67890,
    "sequence_id": 11223
  }
  ```

## TODO
- **UUID**:
    - Consider using UUIDs for unique identifiers instead of integers.
    - Consider adding support for UUID generation in the database.

- **Testing**:
    - Consider implementing end-to-end tests for API endpoints.

## License
This project is licensed under the MIT License.