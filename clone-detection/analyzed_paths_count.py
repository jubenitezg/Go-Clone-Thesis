import pandas as pd
from math import factorial
import common
import json

df = pd.read_csv("full_v2.csv")
print(df.head())

print(df[["analyzed_paths", "pairs", "similarities"]].mean())

print(df[df["similarities"] > 0]
      [["analyzed_paths", "pairs", "similarities"]].mean())

print("SUM")
print(df[df["similarities"] > 0]
      [["analyzed_paths", "pairs", "similarities"]].sum())
