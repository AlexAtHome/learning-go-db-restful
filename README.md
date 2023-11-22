# database access + RESTful API tutorials

- https://go.dev/doc/tutorial/database-access
- https://go.dev/doc/tutorial/web-service-gin

## How to run

This project requires:
- podman
- podman-compose
- go >=1.21.3

```bash
podman compose up --build -d
go run .
```

If you're running the project for the first time, you need to populate the MySQL database.

```bash
podman exec -it mariadb bash

# In the mariadb container
mariadb -u admin -p
use recordings;
source /root/create-tables.sql
quit
exit
```
