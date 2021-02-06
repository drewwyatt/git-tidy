use dialoguer::{theme::ColorfulTheme, MultiSelect};

pub struct Prompt {}

impl Prompt {
    pub fn with(branches: Vec<String>) -> Vec<String> {
        let mut branches = branches;
        let mut selections = MultiSelect::with_theme(&ColorfulTheme::default())
            .with_prompt("Stale branches")
            .items(&branches)
            .interact()
            .unwrap();

        // Sort by descending order
        selections.sort_by(|a, b| b.cmp(a));
        selections
            .into_iter()
            .map(|idx| branches.swap_remove(idx))
            .rev() // sort back to ascending order
            .collect()
    }
}
