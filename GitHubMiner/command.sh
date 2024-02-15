#!/bin/bash

node index.js --query "stars:>100 language:Go" --filename "go-repositories" --batchsize 10
