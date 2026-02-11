---
description: Deploy application to dev environment
---

# Deploy Dev Environment Workflow

This workflow guides you through deploying the Yuan application to the dev environment using Docker Compose.

## Prerequisites

- Docker and Docker Compose installed
- Access to the project repository
- Environment variables configured in `.env.dev`

## Steps

### 1. Navigate to project root
```bash
cd /Users/hieupm/Documents/yuan
```

### 2. Verify environment configuration
Check that `.env.dev` file exists and contains the correct configuration:
```bash
cat deployments/compose/.env.dev
```

If the file doesn't exist, copy from the example:
```bash
cp deployments/compose/.env.example deployments/compose/.env.dev
```

Then edit `.env.dev` to set appropriate values for dev environment.

### 3. Stop any running dev containers
```bash
// turbo
make docker-dev-down
```

### 4. Pull latest code changes (if applicable)
```bash
git pull origin main
```

### 5. Build and start dev environment
```bash
// turbo
make docker-dev-rebuild
```

This command will:
- Build the Docker image with production target
- Start MySQL database on port 3307
- Start the application on port 8081
- Start Adminer (database management UI) on port 8082

### 6. Verify services are running
```bash
// turbo
docker compose -f deployments/compose/docker-compose.dev.yaml ps
```

All services should show status as "Up" or "healthy".

### 7. Check application logs
```bash
docker compose -f deployments/compose/docker-compose.dev.yaml logs -f app
```

Press `Ctrl+C` to exit log viewing.

### 8. Verify application health
```bash
// turbo
curl http://localhost:8081/health
```

Expected response: HTTP 200 OK (or appropriate health check response)

### 9. Access services

- **Application**: http://localhost:8081
- **Adminer (Database UI)**: http://localhost:8082
  - Server: `mysql`
  - Username: `yuan`
  - Password: `devPassword123` (or as configured in `.env.dev`)
  - Database: `yuan`

## Troubleshooting

### If services fail to start:

1. Check logs for errors:
```bash
make docker-dev-logs
```

2. Verify port availability:
```bash
lsof -i :8081
lsof -i :3307
lsof -i :8082
```

3. Check Docker resources:
```bash
docker system df
```

4. Clean up and retry:
```bash
make docker-clean
make docker-dev-up
```

### If database connection fails:

1. Wait for MySQL to be fully ready (can take 30-60 seconds on first start)
2. Check MySQL logs:
```bash
docker compose -f deployments/compose/docker-compose.dev.yaml logs mysql
```

3. Verify database initialization:
```bash
docker exec -it yuan-mysql-dev mysql -u yuan -pdevPassword123 -e "SHOW DATABASES;"
```

## Rollback

If deployment fails and you need to rollback:

1. Stop current deployment:
```bash
make docker-dev-down
```

2. Checkout previous working version:
```bash
git checkout <previous-commit-hash>
```

3. Rebuild and deploy:
```bash
make docker-dev-rebuild
```

## Post-Deployment Verification

- [ ] Application is accessible at http://localhost:8081
- [ ] Database is accessible via Adminer at http://localhost:8082
- [ ] Health check endpoint returns success
- [ ] Application logs show no errors
- [ ] Database migrations completed successfully

## Cleanup

To stop the dev environment:
```bash
make docker-dev-down
```

To remove all data and volumes:
```bash
docker compose -f deployments/compose/docker-compose.dev.yaml down -v
```
