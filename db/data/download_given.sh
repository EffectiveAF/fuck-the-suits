#!/bin/bash
# Steve Phillips / elimisteve
# 2021.02.01

set -euo pipefail

# Download .txt "CSV"-ish file for today

## Default to today's date if no date specified
date=${1:-$(date +%Y%m%d)}

filename="CNMSshvol${date}.txt"
wget "http://regsho.finra.org/$filename" -O raw/"$filename"
