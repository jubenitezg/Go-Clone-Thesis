import json
import logging
import os

import boto3

logging.basicConfig(level=logging.WARN)
logger = logging.getLogger(__name__)

SIMILARITIES = 'similarities.json'
SIMILARITIES_LOC = 'similarities-with-loc.json'
METADATA = 'metadata.json'
DATA = 'data.json'
NORM_REPOSITORIES = 'normalized_repositories.json'
REPOSITORIES = 'repositories.json'
INCOMPLETE_NORM_REPOSITORIES = 'incomplete_normalized_repositories.json'
DATA_TOPICS = 'data_topics.json'
DATA_TOPICS_NORM = 'data_topics_normalized.json'
DATA_TOPICS_READABLE = 'data_topics_readable.json'
BUCKET = os.environ.get('BUCKET')
DUPL = 'dupl_files.json'
FUNCTIONS = 'go-functions.json'
FUNCTIONS_LOC = 'go-functions-with-loc.json'
PATHS = 'go-paths.json'
TOPICS = {
    "topic_0": "Command Line",
    "topic_1": "Client-Server",
    "topic_2": "Cloud",
    "topic_3": "Resource Management",
    "topic_4": "Error Handling & Logging"
}

S3 = boto3.client('s3')


def s3_load_json(key):
    try:
        logger.info(f"Loading {key}")
        response = S3.get_object(Bucket=BUCKET, Key=key)
        return json.loads(response['Body'].read().decode('utf-8'))
    except Exception as e:
        logger.error(f"Error loading {key}: {e}")
        return None


def s3_save_json(key, data):
    try:
        logger.info(f"Saving {key}")
        S3.put_object(Bucket=BUCKET, Key=key, Body=json.dumps(data))
    except Exception as e:
        logger.error(f"Error saving {key}: {e}")


def s3_upload_file(file_name, object_name=None):
    if object_name is None:
        object_name = os.path.basename(file_name)
    try:
        response = S3.upload_file(file_name, BUCKET, object_name)
    except Exception as e:
        logger.error(f"Error saving file {file_name}", e)
