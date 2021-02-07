use std::string::ToString;
pub enum Text<'a> {
    StartupMessage,
    DeletingBranch(&'a str)
}

impl<'a> ToString for Text<'a> {
    fn to_string(&self) -> String {
        match self {
            Text::StartupMessage => "tidying up...".into(),
            Text::DeletingBranch(name) => format!("Deleting ‘{}’...", name)
        }
    }
}
