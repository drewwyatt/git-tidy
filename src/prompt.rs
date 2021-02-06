use dialoguer::{theme::ColorfulTheme, MultiSelect};

pub struct Prompt {}

impl Prompt {
    pub fn with(branches: Vec<String>) -> Vec<String> {
        let selections = MultiSelect::with_theme(&ColorfulTheme::default())
            .with_prompt("Stale branches")
            .items(&branches)
            .interact()
            .unwrap();

        selections
            .into_iter()
            .map(|index| branches[index])
            .collect()
    }
}
