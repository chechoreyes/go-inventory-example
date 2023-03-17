# Inventory backend with Golang

## Clean Archicteure

### Layers

1. Repository: Database
2. Business Logic: Services
3. Presentation: API

# Docker

```docker
docker pull mariadb:10.7.4
docker run -d --name mariadb -p 3306:3306 --env MARIADB_ROOT_PASSWORD=rootroot mariadb:10.7.4
```
