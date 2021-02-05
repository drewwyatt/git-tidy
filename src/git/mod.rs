pub mod models;

use regex::Regex;

use models::{GitError, GitExec};

pub struct Git<F>
where
  F: Fn(&str) -> (),
{
  gone_branch_regex: Regex,
  #[allow(dead_code)]
  path: std::path::PathBuf,
  output: Option<String>,
  _report_progress: F,
}

impl<F> Git<F>
where
  F: Fn(&str) -> (),
{
  pub fn from(path: std::path::PathBuf, report_progress: F) -> Self {
    Git {
      path,
      _report_progress: report_progress,

      output: None,
      gone_branch_regex: Regex::new(r"(?m)^(?:\*| ) ([^\s]+)\s+[a-z0-9]+ \[[^:\n]+: gone\].*$")
        .unwrap(),
    }
  }

  pub fn delete(&mut self, force: bool, branch_name: String) -> Result<&mut Self, GitError> {
    // TODO: figure out how to prevent getting the delete arg in 2 places
    let delete_arg = if force { "-D" } else { "-d" };
    self.report_progress(&format!(
      "running 'git branch {} {}'",
      delete_arg, branch_name
    ));
    self.output = Some(GitExec::delete(force, branch_name)?);
    Ok(self)
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
