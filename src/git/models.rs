#[derive(Debug)]
pub enum GitError {
    CommandError(String),
    ExecError(String),
    ParseError(String),
    UnknownError,
}

impl From<std::io::Error> for GitError {
    fn from(err: std::io::Error) -> Self {
        Self::ExecError(err.to_string())
    }
}

impl From<std::string::FromUtf8Error> for GitError {
    fn from(err: std::string::FromUtf8Error) -> Self {
        Self::ParseError(err.to_string())
    }
}

impl From<std::process::Output> for GitError {
    fn from(output: std::process::Output) -> Self {
        String::from_utf8(output.stderr)
            .map(Self::CommandError)
            .unwrap_or(Self::UnknownError)
    }
}

use std::process::Command;

pub struct GitExec {}

impl GitExec {
    pub fn delete(force: bool, branch_name: &str) -> Result<String, GitError> {
        let delete_arg = if force { "-D" } else { "-d" };
        let output = Command::new("git")
            .arg("branch")
            .arg(delete_arg)
            .arg(branch_name)
            .output()?;

        if output.status.success() {
            return Ok(String::from_utf8(output.stdout)?);
        }

        Err(GitError::from(output))
    }

    pub fn fetch() -> Result<String, GitError> {
        let output = Command::new("git").arg("fetch").output()?;
        if output.status.success() {
            return Ok(String::from_utf8(output.stdout)?);
        }

        Err(GitError::from(output))
    }

    pub fn prune() -> Result<String, GitError> {
        let output = Command::new("git")
            .arg("remote")
            .arg("prune")
            .arg("origin")
            .output()?;
        if output.status.success() {
            return Ok(String::from_utf8(output.stdout)?);
        }

        Err(GitError::from(output))
    }

    pub fn list_branches() -> Result<String, GitError> {
        let output = Command::new("git").arg("branch").arg("-vv").output()?;
        if output.status.success() {
            return Ok(String::from_utf8(output.stdout)?);
        }

        Err(GitError::from(output))
    }
}
