# Docker Compose Local Development

## Quick Start

### 1. Start services
```bash
# From project root
make docker-local-up

# Or manually
docker compose -f deployments/compose/docker-compose.local.yaml --env-file deployments/compose/.env.local up -d
```

### 2. View logs
```bash
docker compose -f deployments/compose/docker-compose.local.yaml logs -f app
```

### 3. Stop services
```bash
make docker-local-down
```

## Services

### MySQL
- **Container**: `mysql-local`
- **Port**: `3306`
- **Database**: `yuan`
- **User**: `yuan`
- **Password**: `password`

### Application
- **Container**: `yuan-app-local`
- **Port**: `8080` (application)
- **Port**: `2345` (debugger)
- **Hot-reload**: ✅ Enabled (source code mounted)

## Features

✅ **Hot-reload**: Code changes automatically reflected (no rebuild needed)
✅ **Debug support**: Delve debugger on port 2345
✅ **MySQL ready**: Database initialized with schema
✅ **Go modules cache**: Faster rebuilds with volume caching

## Environment Variables

Edit `.env.local` to customize:

```bash
ENV=local
MYSQL_HOST=mysql-local
MYSQL_PORT=3306
MYSQL_ROOT_PASSWORD=rootpassword
MYSQL_USER=yuan
MYSQL_PASSWORD=password
MYSQL_DATABASE=yuan
LOG_LEVEL=debug
```

## Development Workflow

1. **Start containers**: `make docker-local-up`
2. **Edit code**: Changes auto-reload
3. **View logs**: `docker compose logs -f app`
4. **Access MySQL**: `docker exec -it mysql-local mysql -u yuan -p`
5. **Stop**: `make docker-local-down`

## Troubleshooting

### Port already in use
```bash
# Check what's using port 3306 or 8080
lsof -i :3306
lsof -i :8080

# Stop conflicting service or change port in .env.local
```

### Container won't start
```bash
# View logs
docker compose -f deployments/compose/docker-compose.local.yaml logs

# Rebuild
make docker-local-rebuild
```

### MySQL connection failed
```bash
# Wait for MySQL to be healthy
docker compose -f deployments/compose/docker-compose.local.yaml ps

# Check MySQL logs
docker logs mysql-local
```

## Useful Commands

```bash
# Rebuild and restart
make docker-local-rebuild

# Access app container
docker exec -it yuan-app-local sh

# Access MySQL
docker exec -it mysql-local mysql -u yuan -p

# View all logs
docker compose -f deployments/compose/docker-compose.local.yaml logs -f

# Clean up volumes
docker compose -f deployments/compose/docker-compose.local.yaml down -v
```
