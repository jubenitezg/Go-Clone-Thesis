from common import s3_load_json, REPOSITORIES, SIMILARITIES, FUNCTIONS


def get_code(repo, sims, paths_exclude=None):
    data = s3_load_json(f"{repo}/{FUNCTIONS}")
    print(f"Repo: {repo}")
    print(f"Similarity: {sims['similarity']}")
    if paths_exclude is None:
        paths_exclude = []
    for func in data:
        for path in paths_exclude:
            if path in func['path']:
                break
        else:
            if func['id'] == sims['id1']:
                print(func['path'])
                print(func['code'])
                break
    for func in data:
        for path in paths_exclude:
            if path in func['path']:
                break
        else:
            if func['id'] == sims['id2']:
                print(func['path'])
                print(func['code'])
                break

    print("=" * 50)


def similarity_pairs():
    repos = s3_load_json(REPOSITORIES)
    # chosen = random.choices(repos, k=10)
    chosen = ['projectdiscovery/nuclei']
    # chosen = ['hashicorp/vault']
    # chosen = ['fatedier/frp']
    print(f"Chosen: {chosen}")
    for repo in chosen:
        similarities = s3_load_json(f"{repo}/{SIMILARITIES}")
        if similarities is None:
            continue
        print(f"Repo: {repo} - Similarities: {len(similarities)}")
        for sims in similarities:
            if sims['similarity'] >= 0.8:
                get_code(repo, sims, ['examples', 'tutorial'])


if __name__ == '__main__':
    similarity_pairs()
