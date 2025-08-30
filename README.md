# envload

Lightweight environment variable loader for Go.

## Usage

```go
package main

import (
    "log"
    "os"
    "github.com/youruser/envload"
)

func main() {
    err := envload.File(".env")
    if err != nil {
        log.Fatal(err)
    }
    
    dbURL := os.Getenv("DATABASE_URL")
    // Use your environment variables...
}
```

## Features

- Loads key=value pairs from files
- Ignores comments (`#`) and empty lines  
- Strips optional quotes from values
- Lightweight with no dependencies

## File Format

```
# Database configuration
DATABASE_URL=postgres://localhost/mydb
API_KEY="your-secret-key"
DEBUG=true
```