#!/bin/bash
set -e

# Load environment variables from .env file
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo ".env file not found. Please create a .env file with the necessary variables."
    exit 1
fi

# Check if Goose is installed
if ! command -v goose &> /dev/null
then
    echo "Goose is not installed. Please install Goose before running this script."
    exit 1
fi

# Run Goose commands
case "$1" in
    up)
        goose -dir $MIGRATION_DIR mysql "$DB_USER:$DB_PASS@tcp($DB_HOST:$DB_PORT)/$DB_NAME" up
        ;;
    down)
        goose -dir $MIGRATION_DIR mysql "$DB_USER:$DB_PASS@tcp($DB_HOST:$DB_PORT)/$DB_NAME" down
        ;;
    create)
        if [ -z "$2" ]; then
            echo "Please provide a name for the migration."
            exit 1
        fi
        goose -s -dir $MIGRATION_DIR create "$2" sql
        ;;
    down-to)
        if [ -z "$2" ]; then
            echo "Please provide a version to rollback to."
            exit 1
        fi
        goose -dir $MIGRATION_DIR mysql "$DB_USER:$DB_PASS@tcp($DB_HOST:$DB_PORT)/$DB_NAME" down-to $2
        ;;
    *)
        echo "Usage: $0 {up|down|create migration_name|down-to version}"
        exit 1
        ;;
esac
