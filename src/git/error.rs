#[derive(Debug)]
pub enum GitError {
  ExecError,
  ParseError,
}

impl From<std::io::Error> for GitError {
  fn from(_: std::io::Error) -> Self {
    Self::ExecError
  }
}

impl From<std::string::FromUtf8Error> for GitError {
  fn from(_: std::string::FromUtf8Error) -> Self {
    Self::ParseError
  }
}
