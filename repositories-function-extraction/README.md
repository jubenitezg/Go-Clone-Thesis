## Repositories Extraction

This sub-module is responsible for extracting the repositories from the GitHub Miner output. It is a simple script clones the repositories and extracts the go functions from it. The output is a set of files corresponding to each repository containing the functions that can be used as input for the next steps of the pipeline.

> [!IMPORTANT]
> The total size of all functions extracted is around 38GB.
> The total number of files created is 15853.

## Run

```bash
./extraction.sh -o <output_directory>
```

## Output

The [count.txt](count.txt) file contains the total number of functions extracted from all repositories.

> [!IMPORTANT]
> The total number of functions is: 67.204.319

## Exceptions

- Skipped repository 189 due to git authentication required
- Skipped repository 5127 due to git authentication required
- Skipped repository 9045 due to git lfs limit
- Skipped repository 10611 due to git authentication required
- Skipped repository 11639 due to git authentication required
- Skipped repository 13596 due to git lfs missing object
- Skipped repository 13718 due to git authentication required
- Skipped repository 15241 due to git authentication required
- Skipped repository 15740 due to git authentication required
