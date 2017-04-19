#!/bin/bash
set -e

echo 'mode: atomic' > coverage.txt
for pkg in $(go list ./... | grep -v vendor); do
  go test -v -race -cover -covermode=atomic -coverprofile=coverage.txt.tmp "$pkg"
  if [[ -f coverage.txt.tmp ]]; then
    tail -n +2 coverage.txt.tmp >> coverage.txt
    rm -f coverage.txt.tmp
  fi
done

bash <(curl -s https://codecov.io/bash) -t 644395c8-662b-4bbf-803e-bbc4500f7c84
rm -f coverage.txt
