# env-templater
### Tool to make .env file use template.

This program takes three arguments from the command line:

- `source_file` - the name of the source .env file.
- `dest_file` - the name of the new file that will be created or overwritten.
- `prefix` - the prefix to remove from lines starting with this prefix.

The program opens the source file and creates a new file. It then iterates over each line of the source file, checks if it starts with the specified prefix, removes that prefix from the line if present, and writes the modified line to the new file. After completing its work, the program reports that it has finished.


## Go Install
```bash
go install github.com/sergoslav/env-templater@latest
```

## Usages Example

Source `.env.template` file:
```env
APP_NAME="My App"
#DB Contigurations
dev-DB_HOST="localhost"
prod-DB_HOST="production.db.host"
dev-DB_DATABASE="project-dev"
prod-DB_DATABASE="project-prod"
DB_USERNAME=user
DB_PASSWORD=password
```

Run:
```bash
$ env-templater .env.template .env dev-
```

Result `.env` file:
```env
APP_NAME="My App"
#DB Contigurations
DB_HOST="localhost"
prod-DB_HOST="production.db.host"
DB_DATABASE="project-dev"
prod-DB_DATABASE="project-prod"
DB_USERNAME=user
DB_PASSWORD=password
```
