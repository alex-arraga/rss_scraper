## Build and run
Open a powershell console and execute ```go build && ./rss_project```

### Steps to update db

1. Create or changes migrations files in ```internal/database/sql/schemas```, using goose.
2. Go to schema folder ```cd internal/database/sql/schemas``` 
3. Execute migration: Run command ```goose postgres "DB_URL" up``` to impact db.
4. Create file to insert values in ```internal/database/sql/queries```

###### Get Go types based of db 

5. Back to root path
6. Update combined file ```go run internal/database/sql/scripts/update_combined.go```. 
7. Generate Go types: Run command ```sqlc generate```