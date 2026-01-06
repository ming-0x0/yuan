#!/bin/bash

# Migration directory
GOOSE_DRIVER="postgres"
GOOSE_MIGRATION_DIR="infrastructure/database/postgres/migrations"
GOOSE_DBSTRING="postgres://yuan:password@localhost:5432/yuan?sslmode=disable"

# Check if Goose is installed
if ! command -v goose &> /dev/null
then
    echo "Goose is not installed. Please install Goose before running this script."
    exit 1
fi

# Run Goose commands
case "$1" in
    up)
        goose -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER $GOOSE_DBSTRING up
        ;;
    down)
        goose -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER $GOOSE_DBSTRING down
        ;;
    create)
        if [ -z "$2" ]; then
            echo "Please provide a name for the migration."
            exit 1
        fi
        goose -s -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER $GOOSE_DBSTRING create "$2" sql
        ;;
    down-to)
        if [ -z "$2" ]; then
            echo "Please provide a version to rollback to."
            exit 1
        fi
        goose -dir $GOOSE_MIGRATION_DIR $GOOSE_DRIVER $GOOSE_DBSTRING down-to $2
        ;;
    *)
        echo "Usage: $0 {up|down|create migration_name|down-to version}"
        exit 1
        ;;
esac
