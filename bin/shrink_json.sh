#!/usr/bin/env bash

for file in pkg/words/data/*; do

  echo "Updating $file"
  sed -i -e 's/"word"/"w"/g' "$file"
  sed -i -e 's/"score"/"s"/g' "$file"
  sed -i -e 's/ //g' "$file"

done
