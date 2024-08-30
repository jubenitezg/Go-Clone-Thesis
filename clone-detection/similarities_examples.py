import random

from common import s3_load_json, REPOSITORIES, SIMILARITIES, FUNCTIONS


def get_code(repo, sims):
    data = s3_load_json(f"{repo}/{FUNCTIONS}")
    print(f"Repo: {repo}")
    print(f"Similarity: {sims['similarity']}")
    for func in data:
        if func['id'] == sims['id1']:
            print(func['path'])
            print(func['code'])
            break
    for func in data:
        if func['id'] == sims['id2']:
            print(func['path'])
            print(func['code'])
            break

    print("=" * 50)


def similarity_pairs():
    repos = s3_load_json(REPOSITORIES)
    chosen = random.choices(repos, k=10)
    print(f"Chosen: {chosen}")
    for repo in chosen:
        similarities = s3_load_json(f"{repo}/{SIMILARITIES}")
        if similarities is None:
            continue
        print(f"Repo: {repo} - Similarities: {len(similarities)}")
        for sims in similarities:
            if sims['similarity'] >= 0.8:
                get_code(repo, sims)


if __name__ == '__main__':
    similarity_pairs()
