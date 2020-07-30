pub mod models;

use models::{GitError, GitExec};
use regex::Regex;

pub struct Git {
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
      gone_branch_regex: Regex::new(r"(?m)^(?:\*| ) ([^\s]+)\s+[a-z0-9]+ \[[^:\n]+: gone\].*$")
        .unwrap(),
    }
  }

  pub fn fetch(self: Self) -> Result<Self, GitError> {
    println!("fetching....");
    GitExec::fetch()?;
    Ok(self)
  }

  pub fn list_branches(self: Self) -> Result<Self, GitError> {
    println!("listing branches...");
    let output = GitExec::list_branches()?;
    println!("output ok: {:?}", output);
    for cap in self.gone_branch_regex.captures_iter(&output) {
      println!("branch: {:?}", &cap[0]);
    }

    Ok(self)
  }
}
