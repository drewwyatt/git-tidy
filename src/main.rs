use structopt::StructOpt;

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

fn main() {
  let args = Cli::from_args();

  println!("force: {}", args.force);
  println!("interactive: {}", args.interactive);
  println!("path: {:?}", args.path.into_os_string());
}
