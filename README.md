# database access + RESTful API tutorials

- https://go.dev/doc/tutorial/database-access
- https://go.dev/doc/tutorial/web-service-gin

## How to run

```bash
podman compose up --build -d
podman exec -it mariadb bash

# In the mariadb container
mariadb -u admin -p
use recordings;
source /root/create-tables.sql
quit
exit

export DBUSER=admin
export DBPASS=qwerty
go run .
```
