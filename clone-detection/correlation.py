import pandas as pd
from scipy.stats import pearsonr, linregress
import seaborn as sns
import matplotlib.pyplot as plt

df = pd.read_csv('full.csv')

SAMPLE = 500

def remove_outliers(df, column):
    df_cp = df.copy()

    Q1 = df_cp[column].quantile(0.25)
    Q3 = df_cp[column].quantile(0.75)
    IQR = Q3 - Q1

    lower_bound = Q1 - 1.5 * IQR
    upper_bound = Q3 + 1.5 * IQR

    return df_cp[(df_cp[column] >= lower_bound) & (df_cp[column] <= upper_bound)]

def pearson_corr(df, column1, column2):
    return pearsonr(df[column1], df[column2])

def graph(df, column1, column2, sample_size=1000):
    df_sample = df.sample(n=min(sample_size, len(df)), random_state=42)

    slope, intercept, r_value, p_value, std_err = linregress(
        df_sample[column1], df_sample[column2]
    )
    print(f"Pendiente (slope): {slope}")
    print(f"Intercepto (intercept): {intercept}")
    print(f"Coeficiente de correlación (r_value): {r_value}")
    print(f"Coeficiente de determinación (R^2): {r_value ** 2}")
    print(f"p-value: {p_value}")
    print(f"Error estándar (std_err): {std_err}")
    print()

    plt.figure(figsize=(8,5))
    sns.regplot(x=df_sample[column1], y=df_sample[column2], scatter_kws={'alpha':0.5}, line_kws={'color':'red'})
    plt.xlabel(column1)
    plt.ylabel(column2)
    plt.title(f"Regresión Lineal: {column1} vs {column2}")
    plt.show()


def commitsVSsimilarities(df):
    df_clean = df[['similarities', 'commits']].dropna()

    df_clean = remove_outliers(df_clean, 'similarities')
    df_clean = remove_outliers(df_clean, 'commits')

    similarities = df_clean['similarities']
    commits = df_clean['commits']


    corr, p_value = pearsonr(similarities, commits)
    print('Pearsons correlation: %.3f' % corr)
    print("p-value: %f" % p_value)

    graph(df_clean, 'commits', 'similarities', SAMPLE)

def contributorsVSsimilarities(df):
    df_clean = df[['similarities', 'contributors']].dropna()

    df_clean = remove_outliers(df_clean, 'similarities')
    df_clean = remove_outliers(df_clean, 'contributors')

    similarities = df_clean['similarities']
    contributors = df_clean['contributors']

    corr, p_value = pearsonr(similarities, contributors)
    print('Pearsons correlation: %.3f' % corr)
    print("p-value: %f" % p_value)

    graph(df_clean, 'contributors', 'similarities', SAMPLE)

def starsVSsimilarities(df):
    df_clean = df[['similarities', 'stars']].dropna()

    df_clean = remove_outliers(df_clean, 'similarities')
    df_clean = remove_outliers(df_clean, 'stars')
    similarities = df_clean['similarities']
    stars = df_clean['stars']

    corr, p_value = pearsonr(similarities, stars)
    print('Pearsons correlation: %.3f' % corr)
    print("p-value: %f" % p_value)

    graph(df_clean, 'stars', 'similarities', SAMPLE)

def filesVSsimilarities(df):
    df_clean = df[['similarities', 'analyzed_paths']].dropna()

    df_clean = remove_outliers(df_clean, 'similarities')
    df_clean = remove_outliers(df_clean, 'analyzed_paths')
    similarities = df_clean['similarities']
    files = df_clean['analyzed_paths']

    corr, p_value = pearsonr(similarities, files)
    print('Pearsons correlation: %.3f' % corr)
    print("p-value: %f" % p_value)

    graph(df_clean, 'analyzed_paths', 'similarities', SAMPLE)

print("filesVSsimilarities")
filesVSsimilarities(df)
print("starsVSsimilarities")
starsVSsimilarities(df)
print("contributorsVSsimilarities")
contributorsVSsimilarities(df)
print("commitsVSsimilarities")
commitsVSsimilarities(df)
