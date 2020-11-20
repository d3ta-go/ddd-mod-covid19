#!/usr/bin/env bash
# bash

echo "go test: github.com/d3ta-go/ddd-mod-covid19/modules/covid19/infrastructure/adapter/covid19/who... "
echo "-------------------------------------------------------------------------------"
echo ""

go test -timeout 120s  github.com/d3ta-go/ddd-mod-covid19/modules/covid19/infrastructure/adapter/covid19/who -v -cover

echo ""
echo "-------------------------------------------------------------------------------"
echo "go test: DONE "
echo ""