pub mod models;

use indicatif::ProgressBar;
use regex::Regex;

use models::{GitError, GitExec};

pub struct Git {
  spinner: ProgressBar,
  path: std::path::PathBuf,
  force: bool,
  interactive: bool,
  gone_branch_regex: Regex,
}

impl Git {
  pub fn from(path: std::path::PathBuf, force: bool, interactive: bool) -> Self {
    Git {
      path,
      force,
      interactive,

      spinner: ProgressBar::new_spinner(),
      gone_branch_regex: Regex::new(r"(?m)^(?:\*| ) ([^\s]+)\s+[a-z0-9]+ \[[^:\n]+: gone\].*$")
        .unwrap(),
    }
  }

  pub fn fetch(self) -> Result<Self, GitError> {
    self.spinner.enable_steady_tick(120);
    self.spinner.set_message("running 'git fetch'...");
    GitExec::fetch()?;
    Ok(self)
  }

  pub fn prune(self) -> Result<Self, GitError> {
    self
      .spinner
      .set_message("running 'git remote prune origin'...");
    GitExec::prune()?;
    Ok(self)
  }

  pub fn list_branches(self) -> Result<Vec<String>, GitError> {
    self.spinner.set_message("running 'git branch -vv'...");
    let output = GitExec::list_branches()?;
    let result = Ok(
      self
        .gone_branch_regex
        .captures_iter(&output)
        .map(|cap| String::from(&cap[1]))
        .collect::<Vec<String>>(),
    );
    let branches = format!("branches: {:?}", result);
    self.spinner.finish_with_message(&branches);
    return result;
  }
}
