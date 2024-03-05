## Repositories Extraction

This sub-module is responsible for extracting the repositories from the GitHub Miner output. It is a simple script clones the repositories and extracts the go functions from it. The output is a set of files corresponding to each repository containing the functions that can be used as input for the next steps of the pipeline.


## Run

```bash
./extraction.sh -o <output_directory>
```
