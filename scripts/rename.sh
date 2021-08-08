#!/usr/bin/env bash

EXCLUDE_DIRS=(
    "! -path *.git/*"
)

find . -type f ${EXCLUDE_DIRS[@]} | xargs sed -i "s/bar/${1}/"
grep  ${1} -r --exclude-dir=.git
