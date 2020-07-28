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
    pub fn fetch() -> Result<(), GitError> {
        let output = Command::new("git").arg("fetch").output()?;
        if output.status.success() {
            return Ok(());
        }

        Err(GitError::from(output))
    }
}
