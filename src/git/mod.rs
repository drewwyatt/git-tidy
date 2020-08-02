pub mod models;

use regex::Regex;

use models::{GitError, GitExec};

pub struct Git<F>
where
  F: Fn(&str) -> (),
{
  gone_branch_regex: Regex,
  path: std::path::PathBuf,
  output: Option<String>,
  _report_progress: F,
}

impl<F> Git<F>
where
  F: Fn(&str) -> (),
{
  pub fn from(
    path: std::path::PathBuf,
    force: bool,
    interactive: bool,
    report_progress: F,
  ) -> Self {
    Git {
      path,
      _report_progress: report_progress,

      output: None,
      gone_branch_regex: Regex::new(r"(?m)^(?:\*| ) ([^\s]+)\s+[a-z0-9]+ \[[^:\n]+: gone\].*$")
        .unwrap(),
    }
  }

  pub fn fetch(&mut self) -> Result<&mut Self, GitError> {
    self.report_progress("running 'git fetch'...");
    self.output = Some(GitExec::fetch()?);
    Ok(self)
  }

  pub fn prune(&mut self) -> Result<&mut Self, GitError> {
    self.report_progress("running 'git remote prune origin'...");
    self.output = Some(GitExec::prune()?);
    Ok(self)
  }

  pub fn list_branches(&mut self) -> Result<&mut Self, GitError> {
    self.report_progress("running 'git branch -vv'...");
    self.output = Some(GitExec::list_branches()?);
    Ok(self)
  }

  pub fn branch_names(&mut self) -> Result<Vec<String>, GitError> {
    self
      .output
      .as_ref()
      .map(|str| {
        self
          .gone_branch_regex
          .captures_iter(&str)
          .map(|cap| String::from(&cap[1]))
          .collect::<Vec<String>>()
      })
      .ok_or(GitError::UnknownError)
  }

  fn report_progress(&mut self, message: &str) {
    let rp = &self._report_progress;
    rp(message);
  }
}
