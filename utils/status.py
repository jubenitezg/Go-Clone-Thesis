import json
import os

import boto3

BUCKET = os.environ.get('BUCKET')
s3 = boto3.client('s3')


def clone_detected():
    response = s3.get_object(Bucket=BUCKET, Key='execution.json')
    urls = json.loads(response['Body'].read().decode('utf-8'))
    return [url for url in urls]


def save_clone_detected(clone_detected):
    s3.put_object(Bucket=BUCKET, Key='execution.json', Body=json.dumps(clone_detected))


def get_skipped():
    response = s3.get_object(Bucket=BUCKET, Key='skipped.json')
    urls = json.loads(response['Body'].read().decode('utf-8'))
    return [url for url in urls]


def save_skipped(skipped):
    s3.put_object(Bucket=BUCKET, Key='skipped.json', Body=json.dumps(skipped))


def get_metadata():
    response = s3.get_object(Bucket=BUCKET, Key='metadata.json')
    return json.loads(response['Body'].read().decode('utf-8'))

def execution_status():
    metadata = get_metadata()
    repositories_completed = clone_detected()
    skipped = get_skipped()
    print(f"Processed: {len(repositories_completed)}/{len(metadata)}")
    print(f"Skipped: {len(skipped)}")

if __name__ == '__main__':
    execution_status()
