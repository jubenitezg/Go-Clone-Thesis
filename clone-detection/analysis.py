from io import StringIO
from sys import platform

import matplotlib
import matplotlib.pyplot as plt
import pandas as pd
import seaborn as sns

import common

if platform in ('linux' or 'linux2'):
    matplotlib.use('TkAgg')


# Preguntas investigacion
# ¿Existe una relación entre la cantidad de pares de clones detectados y el dominio de aplicación de los proyectos de código abierto en el lenguaje de programación Go?
# ¿Cuál es el nivel de detección de clones expresado como porcentaje del total de código analizado en los proyectos de software en Go?
# ¿Cómo varía la cantidad de clones presentes en un proyecto considerando factores como el número de archivos, líneas de código y autores involucrados en el desarrollo de este?
# ¿Qué disparidades y similitudes se pueden identificar al comparar los resultados obtenidos, en relación con los hallazgos similares en otros lenguajes de programación?


def percentage_of_clones_example():
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


def topics_distribution():
    topics = common.s3_load_json(f"{common.DATA_TOPICS_NORM}")
    df = pd.read_json(StringIO(topics))
    topic_columns = [col for col in df.columns if 'topic' in col]
    df['max_topic'] = df[topic_columns].idxmax(axis=1)
    grouped_data = df.groupby('max_topic').size().reset_index(name='count')
    grouped_data.set_index('max_topic', inplace=True)
    print(grouped_data)

    plt.figure(figsize=(10, 6))
    ax = sns.barplot(x=grouped_data.index, y='count', data=grouped_data, color='skyblue')
    ax.bar_label(ax.containers[0], label_type='edge', padding=3)
    plt.title('Data Grouped by Max Topic')
    plt.xlabel('Topic')
    plt.ylabel('Count')
    plt.xticks(rotation=45, ha='right')
    plt.show()


def quick_analysis():
    df = pd.read_csv('full.csv')
    df_melted = df.melt(id_vars=['repo_id', 'similarities'],
                        value_vars=[col for col in df.columns if col.startswith('topics.topic_')],
                        var_name='topic_id', value_name='topic_value')
    df_melted = df_melted.dropna(subset=['topic_value'])
    df_melted = df_melted.dropna(subset=['similarities'])

    domain_clones_relationship = df_melted.groupby('topic_value')['similarities'].mean().reset_index()
    print("Relationship between application domain and number of clones:")
    print(domain_clones_relationship)

    plt.figure(figsize=(12, 6))
    plt.bar(domain_clones_relationship['topic_value'], domain_clones_relationship['similarities'])
    plt.xlabel('Application Domain')
    plt.ylabel('Average Number of Clones')
    plt.title('Relationship Between Application Domain and Number of Clones')
    plt.xticks(rotation=45)
    plt.show()


if __name__ == '__main__':
    quick_analysis()
