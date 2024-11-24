### Steps to update db

1. Create or changes migrations files in ```sql/schemas```, using goose.
2. Go to schema folder ```cd sql/schema``` 
3. Run command ```goose postgres "DB_URL" up``` to impact db.
4. Run command ```go run sql/scripts/update_combined.go```