import argparse
import os

from collector.github_repo_collector import GithubRepositoryCollector
from const.constants import GITHUB_ACCESS_TOKEN


def parse_args() -> argparse.Namespace:
    """
    Parse arguments.
    :return: a namespace with the arguments
    """
    parser = argparse.ArgumentParser(description='Get repositories from GitHub.')
    parser.add_argument('--total', type=int, default=100, help='the total number of repositories to save')
    return parser.parse_args()


if __name__ == '__main__':
    if GITHUB_ACCESS_TOKEN not in os.environ:
        raise Exception('GITHUB_ACCESS_TOKEN not found in environment variables.')
    args = parse_args()
    g = GithubRepositoryCollector()
    g.collect(args.total)
