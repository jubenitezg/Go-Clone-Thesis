import concurrent.futures
import logging

from tqdm import tqdm

import common

logging.basicConfig(level=logging.WARN)

logger = logging.getLogger(__name__)


def is_test(func_path):
    if 'test' in func_path:
        return True
    return False


def is_on_dupl_path(func_path, dupl):
    for path in dupl:
        if path == "":
            continue
        if path in func_path:
            return True
    return False


def filter_functions_paths(func_paths, key, dupl=None):
    if dupl is None:
        logger.info(f"No dupl paths for {key}")
        dupl = []
    no_tests = [func_path for func_path in func_paths if not is_test(func_path)]
    return [func for func in no_tests if is_on_dupl_path(func, dupl)]


def load_function_paths(rkey):
    rfunctions = common.s3_load_json(f"{rkey}/{common.FUNCTIONS}")
    if rfunctions is None:
        logger.error(f"Error loading functions for {rkey}")
        return []
    unique_paths = set()
    for func in rfunctions:
        unique_paths.add(f"{rkey}/{func['path'].removeprefix('/tmp/temp-repo/')}")
    return list(unique_paths)


def do_work(dkey):
    functions = load_function_paths(dkey)
    dupl_paths = common.s3_load_json(f"{dkey}/{common.DUPL}")
    paths = filter_functions_paths(functions, dkey, dupl_paths)
    final = {
        'length': len(paths),
        'paths': paths
    }
    common.s3_save_json(f"{dkey}/{common.PATHS}", final)
    return True


if __name__ == '__main__':
    metadata = common.s3_load_json(common.METADATA)
    keys = [f"{r['owner']}/{r['name']}" for r in metadata]
    with concurrent.futures.ThreadPoolExecutor() as executor:
        list(tqdm(executor.map(do_work, keys), total=len(keys)))
