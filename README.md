## Build and run
Open a powershell console and execute ```go build && ./rss_project```

### Steps to update db

1. Create or changes migrations files in ```sql/schemas```, using goose.
2. Go to schema folder ```cd sql/schema``` 
3. Execute migration: Run command ```goose postgres "DB_URL" up``` to impact db.
4. Create file to insert values in ```sql/queries```

###### Get Go types based of db 

5. Back to root path ```cd ../../```
6. Update combined file ```go run sql/scripts/update_combined.go```. 
7. Generate Go types: Run command ```sqlc generate```