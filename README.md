### Steps to update db

1. Create or changes migrations files in ```sql/schemas```, using goose.
2. Go to schema folder ```cd sql/schema``` 
3. Execute migration. Run command ```goose postgres "DB_URL" up``` to impact db.
4. Back to root path ```cd ../../```
5. Update combined file ```go run sql/scripts/update_combined.go```. Generate Go types based on the schemas ```sqlc generate```