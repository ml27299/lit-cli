#!/bin/sh
export GITHUB_TOKEN=e6b1c967009056061b8347e9c8d30858f4621b47

eval "git tag -a v$1 -m \"update\""
eval "goreleaser"