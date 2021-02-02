#!/bin/bash
# Steve Phillips / elimisteve
# 2021.02.01

set -euo pipefail

# Download .txt "CSV"-ish file for today

filename="CNMSshvol$(date +%Y%m%d).txt"
wget "http://regsho.finra.org/$filename" -O raw/"$filename"
