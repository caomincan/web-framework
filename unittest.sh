#!/bin/bash
set -euo pipefail

PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

GOPATH="${GOPATH:-$HOME/go}"
export GOPATH
export PATH="/usr/local/bin:$GOPATH/bin:$PATH"

if [[ "$(uname -s)" == "Darwin" ]]; then
  # macOS: 没有 mapfile/readarray（默认 bash 3.2），用数组赋值
  PKGS=($(go list ./... | grep -v '/vendor/'))
else
  # Linux (Debian): 有 mapfile/readarray
  mapfile -t PKGS < <(go list ./... | grep -v '/vendor/')
fi

go version
echo "GOPATH: $GOPATH"
echo "PATH: $PATH"

# tools should already be in /usr/local/bin from base image
command -v gocovmerge >/dev/null 2>&1 || { echo "missing gocovmerge in PATH"; exit 1; }
command -v gocover-cobertura >/dev/null 2>&1 || { echo "missing gocover-cobertura in PATH"; exit 1; }

idx=0
COV_FILE="merge.out"
# go get github.com/wadey/gocovmerge
# go install github.com/wadey/gocovmerge
for pkg in "${PKGS[@]}"; do
  echo "#$idx: test $pkg and generate coverage profile"
  go test -v -cover --coverpkg=./... -coverprofile=profile.out $pkg
  if [ $? -eq 0 ]; then
      echo "test $pkg success"
  else
      echo "test $pkg failed"
      exit 1
  fi
  if [ $idx -eq 0 ];then
    echo "first profile"
    cp profile.out $COV_FILE
  else
    echo "merge profile"
    cp $COV_FILE tmp.out
    gocovmerge profile.out tmp.out > $COV_FILE
  fi
  ((idx+=1))
done
# output result
go tool cover -html=$COV_FILE -o coverage.html
# output xml for ci
gocover-cobertura < "$COV_FILE" > "$PROJECT_DIR/cobertura-coverage.xml"

# clean up
rm -f "$PROJECT_DIR"/*.out tmp.out profile.out 2>/dev/null || true