mod git;
mod prompt;

use indicatif::ProgressBar;
use structopt::StructOpt;

use git::models::GitError;
use git::Git;
use prompt::Prompt;

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

    #[structopt(short, long, help = "Print output, but don't delete any branches")]
    dry_run: bool,
}

fn main() -> Result<(), GitError> {
    let args = Cli::from_args();

    if args.dry_run {
        println!("");
        println!("📣 NOTE: --dry-run enabled, no branches will be deleted.");
        println!("");
    }

    let spinner = ProgressBar::new_spinner();
    spinner.set_message("tidying up...");
    spinner.enable_steady_tick(160);

    let mut git = Git::new(|m| spinner.set_message(m));
    let mut gone_branches = git.fetch()?.prune()?.list_branches()?.branch_names()?;

    if gone_branches.is_empty() {
        spinner.finish_with_message("Nothing to do!");
        return Ok(());
    }

    if args.interactive {
        spinner.finish_and_clear();
        gone_branches = Prompt::with(gone_branches);
        if gone_branches.is_empty() {
            println!("Nothing to do!");
            return Ok(());
        }
    }

    if args.dry_run {
        println!("Branches to delete:");
        for branch in gone_branches {
            println!("  - {}", branch);
        }
        println!("")
    } else {
        let spinner = ProgressBar::new_spinner();
        spinner.set_message("deleting branches...");
        spinner.enable_steady_tick(160);

        let mut deletion_errors: Vec<String> = vec![];
        for branch_name in gone_branches {
            spinner.set_message(&format!("deleting \"{}\"...", branch_name));
            match git.delete(args.force, branch_name) {
                Err(GitError::CommandError(msg)) => deletion_errors.push(msg),
                Err(GitError::ExecError(msg)) => deletion_errors.push(msg),
                Err(GitError::ParseError(msg)) => deletion_errors.push(msg),
                Err(GitError::UnknownError) => {
                    deletion_errors.push(String::from("Unknown error encountered."))
                }
                _ => (),
            }
        }

        if deletion_errors.is_empty() {
            spinner.finish_with_message("All done!");
        } else {
            spinner.finish_with_message("Finished with errors.");
            for error in deletion_errors {
                println!("  - {}", error);
            }
        }
    }

    Ok(())
}
