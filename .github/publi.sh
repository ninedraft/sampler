#!/usr/bin/env bash

repo=${1}
tag=${2#"refs/tags/"}

curl "https://proxy.golang.org/$repo/@v/$tag.info" --fail
