import common

metadata = common.s3_load_json(common.METADATA)

new_data = []

total = 0

for repo in metadata:
    total_functions = 0
    functions = common.s3_load_json(f"{repo['owner']}/{repo['name']}/{common.FUNCTIONS}")
    if functions is not None:
        total_functions = len(functions)
    new_data.append({"repo_id": f"{repo['owner']}/{repo['name']}", "total_functions": total_functions})
    total += 1
    print(f"{total}/{len(metadata)}")

common.s3_save_json("repos_total_functions.json", new_data)
