# Topic Analysis
Repository topic analysis classification

## Structure

```bash
.
├── README.md
├── description-extraction
│   ├── README.md
│   ├── extractor
│   │   └── extractor.go
│   ├── go.mod
│   ├── main.go
│   └── model
│       └── repository.go
└── lda
    ├── README.md
    ├── assets
    │   ├── data_topics.json
    │   ├── index.html
    │   ├── output.json
    │   └── topics.json
    ├── poetry.lock
    ├── pyproject.toml
    └── repositories-topic-modeling.ipynb

6 directories, 14 files
```

The `description-extraction` directory has a command line tool for extracting the description, topics and readme of all the repositories from the `GitHub Miner Tool`.

The `lda` directory has an interactive jupyter notebook with 5 topics, the output is in the `assets` directory.

Check the topic distribution in [github.io](https://julianbenitez99.github.io/)
