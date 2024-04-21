import boto3
import json

FILE = "../GitHubMiner/output/go-repositories.json"
BUCKET = "go-project-functions"

with open(FILE, 'r') as f:
    data = json.load(f)

s3 = boto3.client('s3')
for item in data:
    owner = item['owner']
    s3.put_object(Bucket=BUCKET, Key=f'{owner}/')
    print(f"Created folder for {owner}")
