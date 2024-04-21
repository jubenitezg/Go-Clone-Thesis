import boto3
import json

FUNCTIONS_DIR = "/Volumes/JB-DISK/functions-repos"
METADATA = "../metadata-extraction/output/merged.json"
BUCKET = "go-project-functions"
MINER = "../GitHubMiner/output/go-repositories.json"

s3 = boto3.client('s3')

def open_file(file):
    with open(file, 'r') as f:
        data = f.readlines()

    return data

# format
# owner/name
def process(data):
    projects = []
    for line in data:
        r = line.replace("\n", "")
        projects.append(r)
    return projects

def upload_functions(projects):
    for project in projects:
        owner, name = project.split("/")
        try:
            s3.upload_file(f"{FUNCTIONS_DIR}/{owner}-{name}.json", BUCKET, f"{owner}/{name}/go-functions.json")
            print(f"Uploaded functions for {owner}/{name}")
        except FileNotFoundError:
            print(f"No functions found for {owner}/{name}")

def from_missing(file):
    projects = process(open_file(file))
    upload_functions(projects)


def load_miner_ouput():
    with open(MINER, 'r') as f:
        data = json.load(f)
    return data

def load_metadata():
    with open(METADATA, 'r') as f:
        metadata = json.load(f)
    return metadata

def get_missing():
    projects = [f"{item['owner']}/{item['name']}" for item in load_miner_ouput()]
    metadata = [f"{item['owner']}/{item['name']}" for item in load_metadata()]
    return list(set(projects) - set(metadata))

upload_functions(get_missing())
