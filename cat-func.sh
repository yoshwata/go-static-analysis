#!/bin/bash

CMDNAME=`basename $0`
if [ $# -ne 1 ]; then
  echo "Usage: $CMDNAME file1" 1>&2
  exit 1
fi

# allstr=`cat $1`
allstr=$(cat sd-cmd-master/screwdriver/api/api.go)
# cat sd-cmd-master/screwdriver/api/api.go
# echo "$allstr"

re='func ([^\
]*)'
if [[ "$allstr" =~ $re ]]; then
        echo "${BASH_REMATCH[1]}"
else
        echo no match
fi

