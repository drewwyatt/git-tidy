pub mod models;

use indicatif::ProgressBar;
use regex::Regex;
use std::collections::VecDeque;

use models::{GitError, GitExec};

pub struct Git {
  force: bool,
  gone_branch_regex: Regex,
  interactive: bool,
  path: std::path::PathBuf,
  queue: VecDeque<GitProcess>,
}

enum GitProcess {
  Fetch,
  Prune,
  ListBranches,
}

impl Git {
  pub fn from(path: std::path::PathBuf, force: bool, interactive: bool) -> Self {
    Git {
      path,
      force,
      interactive,

      queue: VecDeque::new(),
      gone_branch_regex: Regex::new(r"(?m)^(?:\*| ) ([^\s]+)\s+[a-z0-9]+ \[[^:\n]+: gone\].*$")
        .unwrap(),
    }
  }

  pub fn fetch(&mut self) -> &mut Self {
    self.enqueue(GitProcess::Fetch);
    self
  }

  pub fn prune(&mut self) -> &mut Self {
    self.enqueue(GitProcess::Prune);
    self
  }

  pub fn list_branches(&mut self) -> &mut Self {
    self.enqueue(GitProcess::ListBranches);
    self
  }

  pub fn output(&mut self) -> Result<String, GitError> {
    let spinner = ProgressBar::new_spinner();
    spinner.enable_steady_tick(120);
    let mut output: Option<String> = None;

    for process in self.queue.iter() {
      match process {
        GitProcess::Fetch => {
          spinner.set_message("running 'git fetch'...");
          output = Some(GitExec::fetch()?);
        }
        GitProcess::Prune => {
          spinner.set_message("running 'git remote prune origin'...");
          output = Some(GitExec::prune()?);
        }
        GitProcess::ListBranches => {
          spinner.set_message("running 'git branch -vv'...");
          output = Some(GitExec::list_branches()?);
        }
      }
    }

    spinner.finish_with_message("Done!");
    output.map(Ok).unwrap_or(Err(GitError::UnknownError))
  }

  pub fn to_branch_names(self, output: String) -> Result<Vec<String>, GitError> {
    Ok(
      self
        .gone_branch_regex
        .captures_iter(&output)
        .map(|cap| String::from(&cap[1]))
        .collect::<Vec<String>>(),
    )
  }

  fn enqueue(&mut self, process: GitProcess) {
    self.queue.push_front(process);
  }
}
