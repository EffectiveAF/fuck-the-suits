#!/bin/bash

set -euo pipefail

# Run migrations
for file in $*; do
    psql -U ${pg_user:-postgres} -d fuckthesuits -f "$file"
done
