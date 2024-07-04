import json
import os

import boto3

SIMILARITIES = 'similarities.json'
METADATA = 'metadata.json'
DATA_TOPICS = 'data_topics.json'
DATA_TOPICS_READABLE = 'data_topics_readable.json'
BUCKET = os.environ.get('BUCKET')
TOPICS = {
    "topic_0": "Command Line",
    "topic_1": "Client-Server",
    "topic_2": "Cloud",
    "topic_3": "Resource Management",
    "topic_4": "Error Handling & Logging"
}


# Preguntas investigacion
# ¿Existe una relación entre la cantidad de pares de clones detectados y el dominio de aplicación de los proyectos de código abierto en el lenguaje de programación Go?
# ¿Cuál es el nivel de detección de clones expresado como porcentaje del total de código analizado en los proyectos de software en Go?
# ¿Cómo varía la cantidad de clones presentes en un proyecto considerando factores como el número de archivos, líneas de código y autores involucrados en el desarrollo de este?
# ¿Qué disparidades y similitudes se pueden identificar al comparar los resultados obtenidos, en relación con los hallazgos similares en otros lenguajes de programación?

# TODO:

def load_data_topics():
    s3 = boto3.client('s3')
    response = s3.get_object(Bucket=BUCKET, Key=DATA_TOPICS)
    return json.loads(response['Body'].read().decode('utf-8'))


def load_data_topics_readable():
    s3 = boto3.client('s3')
    response = s3.get_object(Bucket=BUCKET, Key=DATA_TOPICS_READABLE)
    return json.loads(response['Body'].read().decode('utf-8'))


def load_metadata(key):
    s3 = boto3.client('s3')
    response = s3.get_object(Bucket=BUCKET, Key=f"{key}/{METADATA}")
    return json.loads(response['Body'].read().decode('utf-8'))


def load_similarities(key):
    s3 = boto3.client('s3')
    response = s3.get_object(Bucket=BUCKET, Key=f"{key}/{SIMILARITIES}")
    return json.loads(response['Body'].read().decode('utf-8'))


if __name__ == '__main__':
    pass
    # df = pd.DataFrame(load_similarities('mislav/hub'))
    # mets = load_metadata('mislav/hub')
    # print(mets)
    # print(df.head())
    # sims = df['similarity']
    # print(sims.describe())
    #
    # mean_similarity = df.groupby(['id1', 'id2'])['similarity'].mean()
    # print(mean_similarity)
