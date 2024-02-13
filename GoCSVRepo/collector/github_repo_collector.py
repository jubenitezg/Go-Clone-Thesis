import csv
import json
import logging
import os

from github import Auth
from github import Github
from github import PaginatedList
from github import Repository
from tqdm import tqdm

from const.constants import GITHUB_ACCESS_TOKEN, ATTRIBUTES, STATE_PATH, REPOS_PATH

logging.basicConfig()
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


def get_metadata_from_repository(
        repo: Repository.Repository,
) -> dict:
    """
    Get metadata from a repository.
    :param repo: a GitHub repository
    :return: a dictionary with the repository metadata
    """
    metadata = {'full_name': repo.full_name, 'html_url': repo.html_url, 'stargazers_count': repo.stargazers_count,
                'forks_count': repo.forks_count,
                'open_issues_count': repo.open_issues_count, 'description': repo.description,
                'archived': repo.archived,
                'created_at': repo.created_at, 'updated_at': repo.updated_at, 'pushed_at': repo.pushed_at}
    return metadata


class GithubRepositoryCollector(object):

    def __init__(self):
        auth = Auth.Token(os.environ[GITHUB_ACCESS_TOKEN])
        self.current_page = 0
        self._load_previous_state()
        self.g = Github(auth=auth)

    def _load_previous_state(self):
        if not os.path.exists(STATE_PATH):
            logger.info('No previous state found.')
            return
        with open(STATE_PATH, 'r') as f:
            logger.info('Loading previous state.')
            self.status = json.load(f)
            self.current_page = self.status['current_page']
            logger.info(f'Current page: {self.current_page}')

    def _save_state(self, current_page):
        self.status = {'current_page': current_page}
        with open(STATE_PATH, 'w') as f:
            json.dump(self.status, f)
        logger.info('State saved.')

    def save_repositories_csv(
            self,
            repositories: PaginatedList.PaginatedList[Repository.Repository],
            total_repositories: int = 10,
            file_name: str = 'repositories.csv'):
        """
        Save repositories to a CSV file.
        :param repositories: the repositories paginated list from GitHub
        :param total_repositories: the total number of repositories to save
        :param file_name: file name to save the CSV
        :return: nothing
        """
        with open(f'{REPOS_PATH}/{file_name}', 'w') as csvfile:
            writer = csv.DictWriter(csvfile, fieldnames=ATTRIBUTES)
            writer.writeheader()
            current_page = 0
            for repo in tqdm(repositories[self.current_page:total_repositories], total=total_repositories):
                repo_metadata = get_metadata_from_repository(repo)
                writer.writerow(repo_metadata)
                if self.g.get_rate_limit().core.remaining <= 5:
                    logger.info('Rate limit reached. Saving state.')
                    self._save_state(current_page)
                    break
                current_page += 1

    def search_repositories_by_language(
            self,
            language: str,
            sort: str = 'stars',
            order: str = 'desc'
    ) -> PaginatedList.PaginatedList[Repository.Repository]:
        """
        Search repositories by language.
        :param language: the language to search for
        :param sort: the sort field (stars, forks, updated) - default: stars
        :param order: the sort order (asc, desc) - default: desc
        :return: a paginated list of repositories
        """
        return self.g.search_repositories(query=f"language:{language}", sort=sort, order=order)

    def collect(self, total):
        go_repos = self.search_repositories_by_language('go')
        self.save_repositories_csv(go_repos, total)
