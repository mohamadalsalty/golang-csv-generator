package main

import (
    "database/sql"
    "encoding/csv"
    "flag"
    "fmt"
    "log"
    "os"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432 // default port for PostgreSQL
    user     = "root"
    password = "root"
    dbname   = "info"
)

func main() {
    // Define flags for the query and the output CSV file
    query := flag.String("query", "", "The SQL query to execute")
    outputFile := flag.String("output", "output.csv", "The output CSV file")
    flag.Parse()

    if *query == "" {
        log.Fatal("You must provide a query using the -query flag")
    }

    // Construct the connection string
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Open a connection to the database
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Verify the connection to the database
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the database")

    // Perform the query provided as an argument
    rows, err := db.Query(*query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Get the column names from the result
    columns, err := rows.Columns()
    if err != nil {
        log.Fatal(err)
    }

    // Open the output CSV file
    file, err := os.Create(*outputFile)
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()

    // Create a CSV writer
    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write the column names to the CSV file
    if err := writer.Write(columns); err != nil {
        log.Fatal("Cannot write to file", err)
    }

    // Create a slice of interface{}'s to represent each column, and a second slice to contain pointers to each item in the columns slice
    values := make([]interface{}, len(columns))
    valuePtrs := make([]interface{}, len(columns))
    for i := range values {
        valuePtrs[i] = &values[i]
    }

    // Iterate over the rows and write the results to the CSV file
    for rows.Next() {
        err = rows.Scan(valuePtrs...)
        if err != nil {
            log.Fatal(err)
        }

        // Convert the values to strings and write them to the CSV file
        record := make([]string, len(columns))
        for i, val := range values {
            if val != nil {
                record[i] = fmt.Sprintf("%v", val)
            } else {
                record[i] = ""
            }
        }
        if err := writer.Write(record); err != nil {
            log.Fatal("Cannot write to file", err)
        }
    }

    // Check for errors encountered during iteration
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Results have been written to %s\n", *outputFile)
}
