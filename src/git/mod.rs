pub mod error;

use error::GitError;
use std::process::Command;

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

  pub fn list_branches(self: Self) -> Result<Self, GitError> {
    let output = Command::new("git").arg("branch").arg("-vv").output()?;

    String::from_utf8(output.stdout)?
      .lines()
      .for_each(|x| println!("{:?}", x));
    Ok(self)
  }
}