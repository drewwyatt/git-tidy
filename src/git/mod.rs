pub mod models;

use models::{GitError, GitExec};

pub struct Git {
  path: std::path::PathBuf,
  force: bool,
  interactive: bool,
}

impl Git {
  pub fn from(path: std::path::PathBuf, force: bool, interactive: bool) -> Self {
    Git {
      path,
      force,
      interactive,
    }
  }

  pub fn fetch(self: Self) -> Result<Self, GitError> {
    println!("fetching....");
    GitExec::fetch().map(|_| self)
  }

  pub fn list_branches(self: Self) -> Result<Self, GitError> {
    println!("listing branches...");
    let output = GitExec::list_branches()?;
    println!("output ok: {:?}", output);

    Ok(self)
  }
}
