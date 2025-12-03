# MariaDB-Server
A lightweight setup that runs a MariaDB server inside a Docker container alongside a companion service that inserts data into the database every second upto 10 minutes. Useful for testing, load simulation, real-time dashboards, and development environments that need a continuously changing dataset.

# Setup
After cloning the repository:
  1. Create a directory named secrets and add the following files:
     |-secrets
       |-database
         |-db_password.txt
         |-db_root_password.txt
  2. Add a single line text to each of these files. Ex. db_password.txt -> mariadb_password, db_root_password.txt -> mariadb_admin
  3. Save it and run
```bash
docker compose up --build
```
