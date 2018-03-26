#! /bin/bash

# ============================================================================================================== #
# This script will take a file with text and remove all the punctuations, count and display the top unique words #
# ============================================================================================================== #

if [ $# -lt 2 ]; then
  echo "Please provide the filename and number of top words in following order: 
        Argument 1. Filename
        Argument 2. Number of Top words"
  exit 1
fi

cat ${1} | tr -d '[:punct:]' | tr " " "\n" | sort | uniq -c | sort -gr | head -${2}
