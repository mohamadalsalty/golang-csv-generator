# Setting Up PostgreSQL Database and Running the Go Application

To run the Go application, you'll first need to set up a PostgreSQL database and then execute the Go application using the provided Makefile.

### Creating a PostgreSQL Database and Table

1. **Install PostgreSQL**: If you haven't already, install PostgreSQL on your system. You can download it from [PostgreSQL's official website](https://www.postgresql.org/download/).

2. **Create Database and Table**:
   - Once PostgreSQL is installed, open a terminal and log in to PostgreSQL using the command:
     ```sh
     psql -U postgres
     ```
   - Create a new database:
     ```sql
     CREATE DATABASE info;
     ```
   - Connect to the newly created database:
     ```sh
     \c info
     ```
   - Create a table and insert some data:
     ```sql
     CREATE TABLE users (
         id SERIAL PRIMARY KEY,
         name VARCHAR(100) NOT NULL
     );

     INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Charlie');
     ```

### Running the Go Application with Make

To run the Go application using the Makefile with custom query and output file parameters, follow these steps:

1. **Set Query and Output File**: Open the terminal and navigate to the directory containing the Makefile and Go application. Use the `make` command with `QUERY` and `OUTPUT_FILE` parameters to specify the query and output file:
   ```sh
   make QUERY="SELECT id, name FROM users" OUTPUT_FILE="test.csv"
