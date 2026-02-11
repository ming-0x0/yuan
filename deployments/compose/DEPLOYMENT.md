# Deployment Guide

## Server Setup (One-time)

Trên server dev, bạn **chỉ cần** 2 files:

### 1. Tạo thư mục deployment
```bash
ssh user@dev-server
mkdir -p /opt/yuan
cd /opt/yuan
```

### 2. Copy 2 files cần thiết

**Option A: Copy thủ công**
```bash
# Từ máy local
scp deployments/compose/docker-compose.dev.yaml user@dev-server:/opt/yuan/
scp deployments/compose/.env.dev user@dev-server:/opt/yuan/
```

**Option B: Tạo trực tiếp trên server**

Tạo file `docker-compose.dev.yaml`:
```bash
cat > /opt/yuan/docker-compose.dev.yaml << 'EOF'
# Copy nội dung từ deployments/compose/docker-compose.dev.yaml
EOF
```

Tạo file `.env.dev`:
```bash
cat > /opt/yuan/.env.dev << 'EOF'
ENV=dev
DOCKER_IMAGE=your-dockerhub-username/yuan:dev
MYSQL_ROOT_PASSWORD=devRootPassword123
MYSQL_USER=yuan
MYSQL_PASSWORD=devPassword123
MYSQL_DATABASE=yuan
LOG_LEVEL=info
EOF
```

### 3. Cập nhật `.env.dev`
```bash
# Sửa DOCKER_IMAGE thành Docker Hub username thực của bạn
nano /opt/yuan/.env.dev
```

### 4. Cài đặt Docker (nếu chưa có)
```bash
# Ubuntu/Debian
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER

# Logout và login lại để áp dụng group
```

## Deployment Process

### Automatic (via GitHub Actions)

Khi push code lên branch `develop`:
1. ✅ GitHub Actions build Docker image
2. ✅ Push image lên Docker Hub
3. ✅ Copy docker-compose.dev.yaml và .env.dev lên server
4. ✅ Pull image mới nhất
5. ✅ Restart containers
6. ✅ Verify health check

**Không cần làm gì thêm!**

### Manual Deployment

Nếu muốn deploy thủ công:

```bash
# 1. Build và push image (từ máy local)
docker build -f build/package/standalone/Dockerfile -t your-username/yuan:dev .
docker push your-username/yuan:dev

# 2. Deploy trên server
ssh user@dev-server
cd /opt/yuan

# Pull image mới
docker pull your-username/yuan:dev

# Restart services
docker compose -f docker-compose.dev.yaml --env-file .env.dev down
docker compose -f docker-compose.dev.yaml --env-file .env.dev up -d

# Verify
docker compose -f docker-compose.dev.yaml ps
docker compose -f docker-compose.dev.yaml logs -f app
```

## Server Requirements

### Minimum
- **OS**: Ubuntu 20.04+ / Debian 11+ / CentOS 8+
- **RAM**: 2GB
- **Disk**: 10GB
- **Docker**: 20.10+
- **Docker Compose**: 2.0+

### Ports Required
- `8081`: Application
- `3307`: MySQL
- `8082`: Adminer (optional)

## File Structure on Server

```
/opt/yuan/
├── docker-compose.dev.yaml    # Docker Compose config
└── .env.dev                   # Environment variables

# Volumes (tự động tạo bởi Docker)
/var/lib/docker/volumes/
└── compose_mysql_dev_data/    # MySQL data
```

**Không cần source code!** ✅

## Updating Deployment

### Update Docker Image Only
```bash
# Trên server
cd /opt/yuan
docker compose -f docker-compose.dev.yaml pull app
docker compose -f docker-compose.dev.yaml up -d app
```

### Update Configuration
```bash
# Sửa .env.dev hoặc docker-compose.dev.yaml
nano /opt/yuan/.env.dev

# Restart
docker compose -f docker-compose.dev.yaml --env-file .env.dev up -d
```

## Rollback

```bash
# Pull version cũ
docker pull your-username/yuan:dev-<old-commit-sha>

# Update .env.dev
DOCKER_IMAGE=your-username/yuan:dev-<old-commit-sha>

# Restart
docker compose -f docker-compose.dev.yaml --env-file .env.dev up -d
```

## Monitoring

### View Logs
```bash
# All services
docker compose -f docker-compose.dev.yaml logs -f

# App only
docker compose -f docker-compose.dev.yaml logs -f app

# MySQL only
docker compose -f docker-compose.dev.yaml logs -f mysql
```

### Check Status
```bash
docker compose -f docker-compose.dev.yaml ps
```

### Resource Usage
```bash
docker stats
```

## Troubleshooting

### Container won't start
```bash
# Check logs
docker compose -f docker-compose.dev.yaml logs app

# Check if port is in use
sudo lsof -i :8081

# Restart all services
docker compose -f docker-compose.dev.yaml restart
```

### Database connection issues
```bash
# Check MySQL is healthy
docker compose -f docker-compose.dev.yaml ps mysql

# Connect to MySQL
docker exec -it yuan-mysql-dev mysql -u yuan -p

# Check database exists
docker exec -it yuan-mysql-dev mysql -u yuan -p -e "SHOW DATABASES;"
```

### Out of disk space
```bash
# Clean up unused images
docker image prune -a

# Clean up unused volumes
docker volume prune

# Clean up everything
docker system prune -a --volumes
```

## Security Checklist

- [ ] Change default passwords in `.env.dev`
- [ ] Use strong MySQL root password
- [ ] Configure firewall (allow only necessary ports)
- [ ] Use SSH key authentication (disable password auth)
- [ ] Keep Docker updated
- [ ] Regular backups of MySQL data
- [ ] Monitor logs for suspicious activity

## Backup & Restore

### Backup MySQL
```bash
docker exec yuan-mysql-dev mysqldump -u yuan -p yuan > backup.sql
```

### Restore MySQL
```bash
cat backup.sql | docker exec -i yuan-mysql-dev mysql -u yuan -p yuan
```

### Backup Volume
```bash
docker run --rm -v compose_mysql_dev_data:/data -v $(pwd):/backup \
  alpine tar czf /backup/mysql-backup.tar.gz /data
```
