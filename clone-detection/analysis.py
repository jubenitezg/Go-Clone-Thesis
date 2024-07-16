import pandas as pd

import common


# Preguntas investigacion
# ¿Existe una relación entre la cantidad de pares de clones detectados y el dominio de aplicación de los proyectos de código abierto en el lenguaje de programación Go?
# ¿Cuál es el nivel de detección de clones expresado como porcentaje del total de código analizado en los proyectos de software en Go?
# ¿Cómo varía la cantidad de clones presentes en un proyecto considerando factores como el número de archivos, líneas de código y autores involucrados en el desarrollo de este?
# ¿Qué disparidades y similitudes se pueden identificar al comparar los resultados obtenidos, en relación con los hallazgos similares en otros lenguajes de programación?


def percentage_of_clones_example():
    import matplotlib.pyplot as plt
    import matplotlib
    matplotlib.use('TkAgg')
    data = {
        'Group': ['topic_0', 'topic_1', 'topic_2', 'topic_3'],
        'Percentage_of_Clones': [23, 45, 67, 12]
    }

    df = pd.DataFrame(data)

    max_clones = df.loc[df['Percentage_of_Clones'].idxmax()]

    plt.figure(figsize=(10, 6))
    plt.bar(df['Group'], df['Percentage_of_Clones'], color='skyblue')

    plt.bar(max_clones['Group'], max_clones['Percentage_of_Clones'], color='orange')

    plt.title('Percentage of Clones by Group')
    plt.xlabel('Group')
    plt.ylabel('Percentage of Clones')

    for i in range(len(df)):
        plt.text(i, df['Percentage_of_Clones'][i] + 1, f"{df['Percentage_of_Clones'][i]}%", ha='center')

    plt.show()


def clone_count(df, thresh=0.8):
    clones = (df['similarity'] >= thresh).sum()
    return clones
    # no clones
    # not_clones = (df['similarity'] < thresh).sum()
    # print("clones", clones)
    # print("not clones", not_clones)
    # return (clones - not_clones) / not_clones * 100


def utility_print(metadata):
    print(f"Metadata for {metadata['owner']}/{metadata['name']}")
    print(f"Commits: {metadata['commits']}")
    print(f"Contributors: {metadata['contributors']}")
    print(f"Stars: {metadata['contributors']}")
    print(f"Go: {metadata['loc']['Go']}")
    print(f"Created At: {metadata['createdAt']}")


def full(repositories):
    similarities = []
    metadatas = []
    topics = common.s3_load_json(f"{common.DATA_TOPICS}")
    for topic in topics:
        topic['repo_id'] = f"{topic['owner']}/{topic['name']}"
    for repo in repositories:
        key = f"{repo['owner']}/{repo['name']}"
        try:
            similarity = common.s3_load_json(f"{key}/{common.SIMILARITIES}")
            data = {'repo_id': key, 'num_pair_clones': clone_count(pd.DataFrame(similarity))}
            similarities.append(data)
            metadata = common.s3_load_json(f"{key}/{common.METADATA}")
            metadata['repo_id'] = key
            metadatas.append(metadata)
        except:
            print("skipp", key)

    simdf = pd.DataFrame(similarities)
    metasdf = pd.DataFrame(metadatas)
    simmeta_data = pd.merge(simdf, metasdf, on='repo_id')
    topics = pd.DataFrame(topics)
    final_data = pd.merge(simmeta_data, topics, on='repo_id')
    print(final_data.head())
    return final_data


def test_analysis(df):
    import seaborn as sns
    import matplotlib.pyplot as plt
    import matplotlib
    matplotlib.use('TkAgg')
    correlation_matrix = df.corr(numeric_only=True)
    plt.figure(figsize=(10, 8))
    sns.heatmap(correlation_matrix, annot=True, cmap='coolwarm')
    plt.show()





if __name__ == '__main__':
    metadata_full = common.s3_load_json(f'{common.METADATA}')
    data_full = full(metadata_full[:1000])
    test_analysis(data_full)


    # percentage_of_clones_example()
    # key = 'mislav/hub'
    # similarities = common.s3_load_json(f"{key}/{common.SIMILARITIES}")
    # metadata = common.s3_load_json(f"{key}/{common.METADATA}")
    # utility_print(metadata)
    # df = pd.DataFrame(similarities)
    # print(df.head())
    # sims = df['similarity']
    # print(sims.describe())
    #
    # mean_similarity = df.groupby(['id1', 'id2'])['similarity'].mean()
    # print(mean_similarity)
    # print(clone_count(df))
