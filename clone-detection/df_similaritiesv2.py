import pandas as pd
import common


def create_total_df_locs():
    df = pd.read_json("df_locs.json")

    print(df.head())

    repos = common.s3_load_json(common.REPOSITORIES)

    start = 0
    for repo in repos[start:]:
        try:
            code_loc = common.s3_load_json(f"{repo}/{common.FUNCTIONS_LOC}")
            total = sum(code["loc"] for code in code_loc)
            df.at[repo, "loc_total"] = total
            start += 1
            print("completed", repo, start, len(repos))
            if start % 100 == 0:
                df.to_json("df_locs_total.json")
        except Exception as e:
            print("skipped", repo, e)

    print(df.head())
    df.to_json("df_locs_total.json")


# create_total_df_locs()

df = pd.read_json("df_locs_total.json")
print(df.head())
print()

avg_values = df.groupby("topic_assign", as_index=False)[
    ["languages.Go", "cloned_loc", "loc_total", "similarities"]].mean()

print(avg_values.head())


avg_values["cloned_percentage"] = (
    avg_values["cloned_loc"] / avg_values["loc_total"]) * 100
print(avg_values[["topic_assign", "cloned_percentage"]])
