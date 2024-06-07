import concurrent.futures
import json
import os
import shutil
import threading

import boto3
from git import Repo, RemoteProgress

BUCKET = os.environ.get('BUCKET')
s3 = boto3.client('s3')

cpu_count = os.cpu_count()
# lock = threading.Lock()
from tqdm import tqdm


class CloneProgress(RemoteProgress):
    def __init__(self):
        super().__init__()
        self.pbar = tqdm()

    def update(self, op_code, cur_count, max_count=None, message=''):
        self.pbar.total = max_count
        self.pbar.n = cur_count
        self.pbar.refresh()


def get_metadata():
    response = s3.get_object(Bucket=BUCKET, Key='metadata.json')
    return json.loads(response['Body'].read().decode('utf-8'))


def get_processed_repositories():
    response = s3.get_object(Bucket=BUCKET, Key='processed_dupl.json')
    return json.loads(response['Body'].read().decode('utf-8'))


def save_processed_repositories(urls):
    s3.put_object(Bucket=BUCKET, Key='processed_dupl.json', Body=json.dumps(urls))


def execute_dupl(path):
    os.system(f"dupl -plumbing {path} > {path}/dupl.txt")


def save_dupl_output(path, key):
    with open(f"{path}/dupl.txt", 'r') as f:
        output = f.read()
    s3.put_object(Bucket=BUCKET, Key=f"{key}/dupl.txt", Body=output)


def process(repository):
    url = repository['url']
    owner = repository['owner']
    name = repository['name']
    path = f"/tmp/{owner}--{name}"
    print(f"Cloning {owner}/{name}")
    Repo.clone_from(url, path, depth=1, progress=CloneProgress())
    # Repo.clone_from(url, path, depth=1)
    print(f"Executing dupl {owner}/{name}")
    execute_dupl(path)
    save_dupl_output(path, f"{owner}/{name}")
    print(f"Saved dupl {owner}/{name}")
    shutil.rmtree(path)
    # with lock:
    processed.append(url)
    save_processed_repositories(processed)
    print(f"Processed {owner}/{name}")


if __name__ == '__main__':
    # save_processed_repositories([])
    metadata = get_metadata()
    processed = get_processed_repositories()
    print(f"Processed: {len(processed)}/{len(metadata)}")
    missing = [item for item in metadata if item['url'] not in processed]
    # with concurrent.futures.ThreadPoolExecutor(max_workers=10) as executor:
    #     executor.map(process, missing)
    for item in missing:
        process(item)
