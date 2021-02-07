use std::string::ToString;

pub enum Text<'a> {
    BranchesDeleted,
    BranchesToDelete,
    DeletingBranch(&'a str),
    DeletingBranches,
    DryRunEnabled,
    FinishedWithErrors,
    NoBranchesDeleted,
    NothingToDo,
    StartupMessage,
    UnknownErrorEncountered,
}

impl<'a> ToString for Text<'a> {
    fn to_string(&self) -> String {
        match self {
            Text::BranchesDeleted => "Branches deleted:".into(),
            Text::BranchesToDelete => "Branches to delete:".into(),
            Text::DeletingBranch(name) => format!("Deleting ‘{}’...", name),
            Text::DeletingBranches => "Deleting branches...".into(),
            Text::DryRunEnabled => {
                "\n📣 NOTE: --dry-run enabled, no branches will be deleted.\n".into()
            }
            Text::FinishedWithErrors => "Finished with errors:".into(),
            Text::NoBranchesDeleted => "No branches were deleted.".into(),
            Text::NothingToDo => "Nothing to do!".into(),
            Text::StartupMessage => "Tidying up...".into(),
            Text::UnknownErrorEncountered => "An unknown error was encountered".into(),
        }
    }
}
