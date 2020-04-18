use clap::{crate_version, App, Arg};

fn main() {
  let matches = App::new("git-tidy")
    .version(crate_version!())
    .author("Drew Wyatt <drew.j.wyatt@gmail.com>")
    .about("Tidy up stale git branches.")
    .arg("-f, --force 'Allow deleting branches irrespective of their apparent merged status'")
    .arg("-i, --interactive 'Present all \": gone\" branches in list form, allowing user to opt-in to deletions'")
    .arg(
      Arg::with_name("path")
        .help("Path to git repository (defaults to \".\")'")
        .required(false),
    )
    .get_matches();

  let force = matches.is_present("force");
  let interactive = matches.is_present("interactive");
  let path = matches.value_of("path").unwrap_or(".");

  println!("Running git-tidy with args:");
  println!("force: {}", force);
  println!("interactive: {}", interactive);
  println!("path: {}", path);
}
