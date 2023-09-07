import os
import re

from github import Auth
from github import Github
from github import PaginatedList
from github import Repository

from const.constants import GITHUB_ACCESS_TOKEN, GO_MOD_REGEX


def search_repositories_by_language(
        g: Github,
        language: str,
        sort: str = 'stars',
        order: str = 'desc'
) -> PaginatedList.PaginatedList[Repository.Repository]:
    """
    Search repositories by language.
    :param g: an instance of GitHub
    :param language: the language to search for
    :param sort: the sort field (stars, forks, updated) - default: stars
    :param order: the sort order (asc, desc) - default: desc
    :return: nothing
    """
    repositories = g.search_repositories(query=f"language: {language}", sort=sort, order=order)
    return repositories


def get_go_version(
        repositories: PaginatedList.PaginatedList[Repository.Repository],
        page_size: int = 10
) -> [[Repository.Repository, str]]:
    """
    Get the go version from the go.mod file.
    :param repositories: the repositories to search
    :param page_size: the number of repositories to search - default: 10
    :return: a list of tuples containing the repository and the go version
    """
    repos = []
    for repo in repositories[:page_size]:
        try:
            go_mod_content = repo.get_contents('go.mod').decoded_content.decode()
            match = re.search(GO_MOD_REGEX, go_mod_content)
            if match:
                repos.append((repo, match.group(1)))
        except Exception:
            print("Could not find go.mod file in repository: ", repo.full_name)
    return repos


if __name__ == '__main__':
    if GITHUB_ACCESS_TOKEN not in os.environ:
        raise Exception('GITHUB_ACCESS_TOKEN not found in environment variables.')
    auth = Auth.Token(os.environ[GITHUB_ACCESS_TOKEN])
    g = Github(auth=auth)
    go_repos = search_repositories_by_language(g, 'go')
    versions = get_go_version(go_repos)
    print(versions)
