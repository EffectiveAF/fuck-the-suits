#!/bin/bash
# Steve Phillips / elimisteve
# 2021.02.01

set -euo pipefail

# Download all .txt "CSV"-ish files for every day from the current year

for month in `seq -w 1 12`; do
    for day in `seq -w 1 31`; do
        filename="CNMSshvol$(date +%Y)$month$day.txt"
        wget "http://regsho.finra.org/$filename" -O raw/"$filename"
    done
done
