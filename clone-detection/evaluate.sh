#!/bin/bash

python evaluator/evaluator.py -a dataset/test.txt -p saved_models/predictions.txt 2>&1| tee saved_models/score.log