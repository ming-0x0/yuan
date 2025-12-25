#!/bin/bash

# Parse arguments
# $1: driver (postgres, mysql)
# $2: command (up, down, create, down-to)
# $3: optional name/version

GOOSE_DRIVER=$1
COMMAND=$2

# Check if driver is provided
if [ -z "$GOOSE_DRIVER" ]; then
    echo "Please provide a database driver (postgres or mysql)."
    echo "Usage: $0 {postgres|mysql} {up|down|create migration_name|down-to version}"
    exit 1
fi

# Check if command is provided
if [ -z "$COMMAND" ]; then
    echo "Please provide a command."
    echo "Usage: $0 $GOOSE_DRIVER {up|down|create migration_name|down-to version}"
    exit 1
fi

GOOSE_DBSTRING="postgres://yuan:password@localhost:5432/yuan?sslmode=disable"
if [ "$GOOSE_DRIVER" == "mysql" ]; then
    GOOSE_DBSTRING="yuan:password@tcp(localhost:3306)/yuan?parseTime=true"
fi

GOOSE_MIGRATION_DIR="infrastructure/database/$GOOSE_DRIVER/migrations"

# Check if Goose is installed
if ! command -v goose &> /dev/null
then
    echo "Goose is not installed. Please install Goose before running this script."
    exit 1
fi

# Run Goose commands
case "$COMMAND" in
    up)
        goose -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER "$GOOSE_DBSTRING" up
        ;;
    down)
        goose -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER "$GOOSE_DBSTRING" down
        ;;
    create)
        if [ -z "$3" ]; then
            echo "Please provide a name for the migration."
            exit 1
        fi
        mkdir -p $GOOSE_MIGRATION_DIR
        goose -s -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER "$GOOSE_DBSTRING" create "$3" sql
        ;;
    down-to)
        if [ -z "$3" ]; then
            echo "Please provide a version to rollback to."
            exit 1
        fi
        goose -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER "$GOOSE_DBSTRING" down-to $3
        ;;
    *)
        echo "Usage: $0 {postgres|mysql} {up|down|create migration_name|down-to version}"
        exit 1
        ;;
esac


