from __future__ import print_function
import logging
import os
import shutil
import sys
import tempfile
from subprocess import Popen

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger('gocookiecutter')
thismodule = sys.modules[__name__]

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)
FINAL_DIRECTORY = PROJECT_DIRECTORY

def init_git():
    """
    Initialises git on the new project folder
    """
    GIT_COMMANDS = [
        ['git', 'init'],
        ['git', 'add', '.'],
        ['git', 'commit', '-aq', '-m', 'Initial Commit'],
    ]

    for command in GIT_COMMANDS:
        git = Popen(command, cwd=FINAL_DIRECTORY)
        git.wait()


init_git()

logger.info('Your project is ready to go, to start working:')
logger.info('`cd {0}`'.format(PROJECT_DIRECTORY))
