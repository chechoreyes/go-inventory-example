# Inventory backend with Golang

## Clean Archicteure

### Layers

1. Repository: Database
2. Business Logic: Services
3. Presentation: API

## Docker

```docker
docker pull mariadb
docker run -d --name mariadb -p 3306:3306 --env MARIADB_ROOT_PASSWORD=rootroot mariadb
```

## Unit tests: Mockery

```cmd
go get github.com/vektra/mockery/v3@lastest
```

### Use (recursively)

```cmd
go install github.com/vektra/mockery/v2@v2.20.0
```
