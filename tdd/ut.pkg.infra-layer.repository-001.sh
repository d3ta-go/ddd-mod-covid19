#!/usr/bin/env bash
# bash

echo "go test: github.com/d3ta-go/ddd-mod-covid19/modules/covid19/infrastructure/repository... "
echo "-------------------------------------------------------------------------------"
echo ""

go test -timeout 120s  github.com/d3ta-go/ddd-mod-covid19/modules/covid19/infrastructure/repository -v -cover

echo ""
echo "-------------------------------------------------------------------------------"
echo "go test: DONE "
echo ""