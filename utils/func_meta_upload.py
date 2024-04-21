import boto3
import json
import os
from concurrent.futures import ThreadPoolExecutor

METADATA = "../metadata-extraction/output/merged.json"
FUNCTIONS_DIR = "/Volumes/JB-DISK/functions-repos"
BUCKET = "go-project-functions"

with open(METADATA, 'r') as f:
    metadata = json.load(f)

s3 = boto3.client('s3')

def upload_function(item):
    name = item['name']
    owner = item['owner']
    print(f"Created folder for {owner}/{name}")
    s3.put_object(Bucket=BUCKET, Key=f'{owner}/{name}/')
    meta_json = json.dumps(item)
    s3.put_object(Bucket=BUCKET, Key=f'{owner}/{name}/metadata.json', Body=meta_json)
    print(f"Uploaded metadata for {owner}/{name}")
    try:
        s3.upload_file(f"{FUNCTIONS_DIR}/{owner}-{name}.json", BUCKET, f"{owner}/{name}/go-functions.json")
        print(f"Uploaded functions for {owner}/{name}")
    except FileNotFoundError:
        print(f"No functions found for {owner}/{name}")

def upload_all(metadata):
    with ThreadPoolExecutor() as executor:
        executor.map(upload_function, metadata)

upload_all(metadata)
