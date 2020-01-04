#!/bin/bash
$version=$(< VERSION)
$version="v${version}"
git commit -S -m "$1"
git tag -a "${version}" -m "$1"
