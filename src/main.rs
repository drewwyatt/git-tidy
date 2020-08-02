mod git;

use indicatif::ProgressBar;
use structopt::StructOpt;

use git::models::GitError;
use git::Git;

#[derive(StructOpt)]
#[structopt(
  about = "Tidy up stale git branches.",
  author = "Drew Wyatt <drew.j.wyatt@gmail.com>",
  name = "git-tidy"
)]
struct Cli {
  #[structopt(
    short,
    long,
    help = "Allow deleting branches irrespective of their apparent merged status"
  )]
  force: bool,

  #[structopt(
    short,
    long,
    help = r#"Present all ": gone" branches in list form, allowing opt-in to deletions"#
  )]
  interactive: bool,

  #[structopt(
    parse(from_os_str),
    default_value = ".",
    help = r#"Path to git repository (defaults to ".")"#
  )]
  path: std::path::PathBuf,
}

fn main() -> Result<(), GitError> {
  let args = Cli::from_args();
  let spinner = ProgressBar::new_spinner();
  spinner.set_message("tidying up...");
  spinner.enable_steady_tick(160);

  let out = Git::from(args.path, |m| spinner.set_message(m))
    .fetch()?
    .prune()?
    .list_branches()?
    .branch_names()?;

  spinner.finish_with_message(&format!("gone branches: {:?}", out));

  Ok(())
}
