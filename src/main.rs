mod git;
mod i18n;
mod prompt;

use indicatif::ProgressBar;
use structopt::StructOpt;

use git::models::GitError;
use git::Git;
use i18n::Text;
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
        println!("{}", Text::DryRunEnabled.to_string());
    }

    let spinner = ProgressBar::new_spinner();
    spinner.set_message(&Text::StartupMessage.to_string());
    spinner.enable_steady_tick(160);

    let mut git = Git::new(|m| spinner.set_message(m));
    let gone_branches = git.fetch()?.prune()?.list_branches()?.branch_names()?;

    if gone_branches.is_empty() {
        spinner.finish_with_message(&Text::NothingToDo.to_string());
        return Ok(());
    }

    spinner.finish_and_clear();
    let mut stale_branches = gone_branches;

    if args.interactive {
        stale_branches = Prompt::with(stale_branches);
        if stale_branches.is_empty() {
            println!("{}", Text::NothingToDo.to_string());
            return Ok(());
        }
    }

    if args.dry_run {
        println!("{}", Text::BranchesToDelete.to_string());
        for branch in stale_branches {
            println!("  - {}", branch);
        }
        println!("")
    } else {
        let spinner = ProgressBar::new_spinner();
        spinner.set_message(&Text::DeletingBranches.to_string());
        spinner.enable_steady_tick(160);

        let (deleted_branches, deletion_errors) =
            stale_branches
                .into_iter()
                .fold((vec![], vec![]), |(mut del, mut err), branch_name| {
                    spinner.set_message(&Text::DeletingBranch(&branch_name).to_string());
                    match git.delete(args.force, &branch_name) {
                        Err(GitError::CommandError(msg)) => err.push((branch_name, msg)),
                        Err(GitError::ExecError(msg)) => err.push((branch_name, msg)),
                        Err(GitError::ParseError(msg)) => err.push((branch_name, msg)),
                        Err(GitError::UnknownError) => {
                            err.push((branch_name, Text::UnknownErrorEncountered.to_string()))
                        }
                        _ => del.push(branch_name),
                    };

                    (del, err)
                });

        spinner.finish_and_clear();
        if deleted_branches.is_empty() {
            println!("{}", Text::NoBranchesDeleted.to_string());
        } else {
            println!("{}", Text::BranchesDeleted.to_string());
            for branch_name in deleted_branches {
                println!("  - {}", branch_name);
            }
        }

        if !deletion_errors.is_empty() {
            println!("{}", Text::FinishedWithErrors.to_string());
            for (branch_name, error) in deletion_errors {
                println!("  - {}: {}", branch_name, error);
            }
        }
    }

    Ok(())
}
