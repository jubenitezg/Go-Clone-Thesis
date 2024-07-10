import pandas as pd

import common

# Preguntas investigacion
# ¿Existe una relación entre la cantidad de pares de clones detectados y el dominio de aplicación de los proyectos de código abierto en el lenguaje de programación Go?
# ¿Cuál es el nivel de detección de clones expresado como porcentaje del total de código analizado en los proyectos de software en Go?
# ¿Cómo varía la cantidad de clones presentes en un proyecto considerando factores como el número de archivos, líneas de código y autores involucrados en el desarrollo de este?
# ¿Qué disparidades y similitudes se pueden identificar al comparar los resultados obtenidos, en relación con los hallazgos similares en otros lenguajes de programación?


if __name__ == '__main__':
    key = 'mislav/hub'
    similarities = common.s3_load_json(f"{key}/{common.SIMILARITIES}")
    metadata = common.s3_load_json(common.METADATA)
    df = pd.DataFrame(similarities)
    print(df.head())
    sims = df['similarity']
    print(sims.describe())

    mean_similarity = df.groupby(['id1', 'id2'])['similarity'].mean()
    print(mean_similarity)
