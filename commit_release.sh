#!/bin/bash
$version=$(< VERSION)
git commit -S -a v$version -m $1
