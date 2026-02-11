# Docker Compose Setup for Yuan

This directory contains Docker Compose configurations for running the Yuan application in different environments.

## Available Configurations

### 1. Base Configuration (`docker-compose.yaml`)
- **Purpose**: Default configuration for basic setup
- **Services**: MySQL database + Application
- **Port**: Application runs on `8080`
- **Database Port**: MySQL on `3306`

### 2. Local Development (`docker-compose.local.yaml`)
- **Purpose**: Local development with hot-reload
- **Services**: MySQL database + Application (development mode)
- **Features**:
  - Hot-reload support with volume mounting
  - Debug port exposed (`2345`)
  - Source code mounted for live changes
  - Debug logging enabled
- **Ports**:
  - Application: `8080`
  - Debugger: `2345`
  - MySQL: `3306`

### 3. Dev Environment (`docker-compose.dev.yaml`)
- **Purpose**: Development/staging environment
- **Services**: MySQL database + Application + Adminer
- **Features**:
  - **Uses pre-built Docker image** (no source code needed on server!)
  - Environment variables support
  - Auto-restart policies
  - Health checks
  - Adminer for database management
- **Ports**:
  - Application: `8081`
  - MySQL: `3307`
  - Adminer: `8082`

> **Note**: Dev environment pulls pre-built Docker images from Docker Hub. You only need `docker-compose.dev.yaml` and `.env.dev` on the server - **no source code required**! See [DEPLOYMENT.md](DEPLOYMENT.md) for details.


## Usage

### Local Development
```bash
# Start local environment
docker-compose -f docker-compose.local.yaml --env-file .env.local up

# Start in detached mode
docker-compose -f docker-compose.local.yaml --env-file .env.local up -d

# View logs
docker-compose -f docker-compose.local.yaml logs -f

# Stop
docker-compose -f docker-compose.local.yaml down

# Stop and remove volumes
docker-compose -f docker-compose.local.yaml down -v
```

### Dev Environment
```bash
# Start dev environment
docker-compose -f docker-compose.dev.yaml --env-file .env.dev up

# Start in detached mode
docker-compose -f docker-compose.dev.yaml --env-file .env.dev up -d

# Rebuild and start
docker-compose -f docker-compose.dev.yaml --env-file .env.dev up --build

# View logs
docker-compose -f docker-compose.dev.yaml logs -f

# Stop
docker-compose -f docker-compose.dev.yaml down
```

### Base Configuration
```bash
# Start base environment
docker-compose up

# Start in detached mode
docker-compose up -d

# Stop
docker-compose down
```

## Environment Variables

### Local (.env.local)
- `ENV=local`
- `LOG_LEVEL=debug`
- `MYSQL_PASSWORD=password`

### Dev (.env.dev)
- `ENV=dev`
- `LOG_LEVEL=info`
- `MYSQL_PASSWORD=devPassword123`
- `MYSQL_ROOT_PASSWORD=devRootPassword123`

You can customize these by editing the respective `.env.*` files.

## Database Access

### Local Environment
- **Host**: `localhost`
- **Port**: `3306`
- **Database**: `yuan`
- **User**: `yuan`
- **Password**: `password`

### Dev Environment
- **Host**: `localhost`
- **Port**: `3307`
- **Database**: `yuan`
- **User**: `yuan`
- **Password**: `devPassword123`
- **Adminer UI**: http://localhost:8082

## Useful Commands

```bash
# Execute commands in running container
docker-compose exec app sh

# View database logs
docker-compose logs mysql

# Restart specific service
docker-compose restart app

# Remove all containers and volumes
docker-compose down -v

# Build without cache
docker-compose build --no-cache

# Scale services (if needed)
docker-compose up --scale app=3
```

## Troubleshooting

### Port Already in Use
If you get a port conflict error, either:
1. Stop the conflicting service
2. Change the port mapping in the docker-compose file

### Database Connection Issues
1. Ensure MySQL is healthy: `docker-compose ps`
2. Check logs: `docker-compose logs mysql`
3. Verify environment variables are set correctly

### Application Not Starting
1. Check application logs: `docker-compose logs app`
2. Ensure database is ready (health check passed)
3. Verify configuration files are mounted correctly

## Network

All services run on isolated Docker networks:
- Local: `yuan-local-network`
- Dev: `yuan-dev-network`
- Base: `yuan-network`

## Volumes

Persistent data is stored in Docker volumes:
- Local: `mysql_local_data`, `go_modules`
- Dev: `mysql_dev_data`
- Base: `mysql_data`

To remove volumes: `docker-compose down -v`
