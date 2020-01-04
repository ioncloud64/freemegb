#!/bin/bash
$version=$(< VERSION)
git commit -S -m "$1"
git tag -a "v$version" -m "$1"
