#!/usr/bin/env bash

for file in pkg/words/data_old/*; do

  echo "Updating $file"

  new_file=$(basename "$file" .json) 
  cat "$file" | jq '.top_words | map(.word) | .[]' -r >> "pkg/words/data/$new_file.txt"

done
