# SalesForge API

## Overview

SalesForge API is a service for managing email sequences and steps within those sequences. It provides functionality to add, update, and delete sequences and steps.

## Prerequisites

- Go 1.16+
- PostgreSQL

## Project Structure

- `internal/models`: Contains the data models.
- `internal/persistence`: Contains the repository layer for database interactions.
- `internal/service`: Contains the service layer for business logic.
- `internal/psql`: Contains the PostgreSQL connection setup.
- `config`: Contains configuration files.
- `docker-entrypoint-initdb.d`: Contains SQL scripts for initializing the database.

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
MySql:
  Db: "yourdatabase"
  User: "youruser"
  Pass: "yourpassword"
  Host: "localhost"
  Port: 5432
Logger:
  Level: "info"
  Format: "json"
```

## Running the Service
To run the SalesForge API project, follow these steps:  

1. Clone the repository:  
```git clone ```
2. Navifate to its directory:  
```cd salesforge-api```
3. Build and run the project using the Makefile:  
`make`

### API Endpoints

#### Add Sequence

- **Endpoint**: `/add-sequence`
- **Method**: `POST`
- **Payload**:
  ```json
  {
    "account_id": 12345,
    "sequence_name": "Sequence Name",
    "sequence_open_tracking_enabled": true,
    "sequence_click_tracking_enabled": true,
    "steps": [
      {
        "step_email_subject": "Subject of the email",
        "step_email_body": "Body of the email",
        "wait_days": 1,
        "eligible_start_time": 1633036800,
        "eligible_end_time": 1633123200
      }
    ]
  }
  ```

#### Update Sequence

- **Endpoint**: `/update-sequence`
- **Method**: `POST`
- **Payload**:
  ```json
  {
    "account_id": 12345,
    "sequence_id": 11223,
    "sequence_name": "Updated Sequence Name",
    "sequence_open_tracking_enabled": false,
    "sequence_click_tracking_enabled": true
  }
  ```

#### Update Step

- **Endpoint**: `/update-step`
- **Method**: `POST`
- **Payload**:
  ```json
  {
    "account_id": 12345,
    "step_id": 67890,
    "sequence_id": 11223,
    "step_email_subject": "Updated Subject of the email",
    "step_email_body": "Updated Body of the email"
  }
  ```

#### Delete Step

- **Endpoint**: `/delete-step`
- **Method**: `POST`
- **Payload**:
  ```json
  {
    "account_id": 12345,
    "step_id": 67890,
    "sequence_id": 11223
  }
  ```

##TODO

- **Testing**:
    - Write unit tests for all service methods.
    - Write integration tests for database interactions.
    - Implement end-to-end tests for API endpoints.

- **Healthcheck**:
    - Implement a healthcheck endpoint to monitor the status of the service.
    - Ensure the healthcheck endpoint verifies database connectivity.

- **Logging**:
    - Integrate a logging library to capture and store logs.
    - Ensure logs include relevant information such as request details, errors, and performance metrics.
    - Configure log levels and formats as specified in the `config.yaml`.

- **Error Handling**:
    - Implement comprehensive error handling for all service methods.
    - Ensure meaningful error messages are returned to the client.

- **Configuration**:
    - Validate configuration values on startup.
    - Add support for environment-specific configurations.

- **Security**:
    - Implement authentication and authorization for API endpoints.
    - Ensure sensitive data is encrypted in transit and at rest.

- **Documentation**:
    - Document all API endpoints with examples and expected responses.
    - Provide setup and usage instructions for developers.

- **Performance**:
    - Optimize database queries for better performance.
    - Implement caching where appropriate to reduce load on the database.

## License
This project is licensed under the MIT License.