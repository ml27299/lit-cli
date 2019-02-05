#!/bin/sh

eval "git tag -a v$1 -m \"update\""
eval "goreleaser"