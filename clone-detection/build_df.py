import common
import pandas as pd

from tqdm import tqdm


def build_dataframe():
    repos = common.s3_load_json(common.NORM_REPOSITORIES)
    print(f"Length: {len(repos)}")
    df = pd.DataFrame()
    for repo_id in tqdm(repos):
        row = common.s3_load_json(f"{repo_id}/{common.DATA}")
        df = pd.concat([df, pd.json_normalize(row)])

    df.to_csv('full.csv')


if __name__ == '__main__':
    build_dataframe()
