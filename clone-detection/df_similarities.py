import pandas as pd
import json


df = pd.read_csv("full_v2.csv")

print(df.head())

# df_lines = df[['repo_id', 'languages.Go', 'similarities']]
df_lines = df[['repo_id', 'similarities']]
df_lines.set_index("repo_id", inplace=True)


topic_assigment = []
with open('topic_assigment.json') as f:
    topic_assigment = json.load(f)

for repo in topic_assigment:
    try:
        df_lines.at[f"{repo['owner']}/{repo['name']}",
                    "topic_assign"] = repo["topic_assign"]
    except Exception as e:
        print("skipped", repo, e)
        break

print(df_lines.head())

avg_values = df_lines.groupby("topic_assign", as_index=False)[
    ["similarities"]].mean()

print(avg_values.head(50))
