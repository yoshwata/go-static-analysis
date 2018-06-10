#!/bin/bash

version="4.10.1"

if [[ ${version} =~ ^([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
  all=${BASH_REMATCH[0]}
  major=${BASH_REMATCH[1]}
  minor=${BASH_REMATCH[2]}
  patch=${BASH_REMATCH[3]}

  echo ${all}    # 4.10.1
  echo ${major}  # 4
  echo ${minor}  # 10
  echo ${patch}  # 1
fi

str='foo = 1
bar = 2
boo = 3
bar = 5
'
re='bar = ([^\
]*)'
if [[ "$str" =~ $re ]]; then
        echo "${BASH_REMATCH[1]}"
else
        echo no match
fi
