#!/bin/bash
version=v$(< VERSION)
git commit -S -m $version
git tag -a $version -m $version
