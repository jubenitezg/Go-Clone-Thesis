from math import factorial

import pandas as pd

import common


class Data(object):

    def __init__(self, repo_id: str, topics: dict, loc: dict, analyzed_paths: int, pairs: int, similarities: int,
                 issues: int, license: str, commits: int, created_at: str,
                 contributors: int, stars: int, languages: dict):
        self.repo_id = repo_id
        self.topics = topics
        self.loc = loc
        self.analyzed_paths = analyzed_paths
        self.pairs = pairs
        self.similarities = similarities
        self.issues = issues
        self.license = license
        self.commits = commits
        self.created_at = created_at
        self.contributors = contributors
        self.stars = stars
        self.languages = languages


def nCr(n, r):
    return int((factorial(n) / (factorial(r)
                                * factorial(n - r))))


def compute_similarities(repo_id, threshold=0.8):
    try:
        similarities = common.s3_load_json(f"{repo_id}/{common.SIMILARITIES}")
        count = 0
        for similarity in similarities:
            if similarity['similarity'] >= threshold:
                count += 1
        return count
    except Exception as e:
        return None


def find_topic(topics, repo_id):
    for topic in topics:
        if topic['repo_id'] == repo_id:
            return topic['topic']
    return None


def load_data(repo_id, topics):
    try:
        metadata = common.s3_load_json(f"{repo_id}/{common.METADATA}")
        analyzed_paths = common.s3_load_json(f"{repo_id}/{common.PATHS}")['length']
        pairs = nCr(analyzed_paths, 2)
        similarities = compute_similarities(repo_id)
        issues = metadata['issues']
        loc = metadata['loc']['Go']
        license = metadata['license']
        commits = metadata['commits']
        created_at = metadata['createdAt']
        contributors = metadata['contributors']
        languages = metadata['languages']
        stars = metadata['stars']
        return Data(repo_id, topics[repo_id]['topics'], loc, analyzed_paths, pairs, similarities, issues, license,
                    commits,
                    created_at,
                    contributors, stars, languages)
    except Exception as e:
        return None


def load_topics():
    dtopics = common.s3_load_json(common.DATA_TOPICS)
    full_topics = {}
    for topic in dtopics:
        full_topics[f'{topic["owner"]}/{topic["name"]}'] = topic
        del topic['owner']
        del topic['name']
    return full_topics


def normalize_data():
    repos = common.s3_load_json(common.REPOSITORIES)
    repositories_completed = common.s3_load_json(common.NORM_REPOSITORIES)
    incomplete = common.s3_load_json(common.INCOMPLETE_NORM_REPOSITORIES)
    print(f"Processed: {len(repositories_completed)}/{len(repos)}")
    print(f"Incomplete: {len(incomplete)}")
    process = [r for r in repos if r not in repositories_completed and r not in incomplete]
    tp = load_topics()
    for repo_id in process[:10]:
        data = load_data(repo_id, tp)
        if data is not None:
            common.s3_save_json(f"{repo_id}/{common.DATA}", data.__dict__)
            repositories_completed.append(repo_id)
            common.s3_save_json(common.NORM_REPOSITORIES, repositories_completed)
            print("completed", repo_id)
        else:
            incomplete.append(repo_id)
            common.s3_save_json(common.INCOMPLETE_NORM_REPOSITORIES, incomplete)
            print("incomplete", repo_id)


def panda_approach():
    repos = common.s3_load_json(common.NORM_REPOSITORIES)
    df = pd.DataFrame()
    for repo in repos:
        data = common.s3_load_json(f"{repo}/{common.DATA}")
        df = pd.concat([df, pd.json_normalize(data)])
    print(df.head())


if __name__ == '__main__':
    # common.s3_save_json(common.NORM_REPOSITORIES, [])
    normalize_data()
    # panda_approach()
