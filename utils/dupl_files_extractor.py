import json
import os

import boto3
from tqdm.contrib.concurrent import thread_map

"""
Extracts the paths of the duplicate files from the dupl.txt files and saves them in a json file.
"""

SPLIT_ON = ' duplicate of '
BUCKET = os.environ.get('BUCKET')
DUPL = 'dupl.txt'

s3 = boto3.client('s3')

cpu_count = os.cpu_count()


def get_repos():
    response = s3.get_object(Bucket=BUCKET, Key='metadata.json')
    metadata = json.loads(response['Body'].read().decode('utf-8'))
    return [
        {'name': i['name'], 'owner': i['owner']} for i in metadata
    ]


def save_paths(key, paths):
    s3.put_object(Bucket=BUCKET, Key=f'{key}/dupl_files.json', Body=json.dumps(list(paths)))


def get_paths(content, path_name):
    paths = set()
    for line in content.split('\n'):
        for path in line.split(SPLIT_ON):
            start_pos = path.rfind(path_name)
            end_pos = path.rfind(".go") + 3
            file_path = path[start_pos + len(path_name):end_pos]
            if "\\" in file_path:
                file_path = file_path.replace('\\', '/')
            paths.add(file_path)
    return paths


def process(repository):
    owner = repository['owner']
    name = repository['name']
    key = f"{owner}/{name}"
    try:
        response = s3.get_object(Bucket=BUCKET, Key=f"{key}/{DUPL}")
        content = response['Body'].read().decode('utf-8')
        path_name = f"{owner}--{name}"
        paths = get_paths(content, path_name)
        save_paths(key, paths)
    except:
        print(f"No dupl.txt available for {key}")


if __name__ == '__main__':
    print("Get repos")
    repositories = get_repos()
    r = thread_map(process, repositories, max_workers=cpu_count)
