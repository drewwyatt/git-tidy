use dialoguer::{theme::ColorfulTheme, MultiSelect};

pub struct List {
  branches: Vec<String>
}

impl From<Vec<String>> for List {
  fn from(branches: Vec<String>) -> Self {
    List { branches }
  }
}

impl List {
  pub fn present(&self) -> Vec<String> {
    let indexes = MultiSelect::with_theme(&ColorfulTheme::default())
      .with_prompt("Stale branches")
      .items(&self.branches)
      .interact()
      .unwrap();

    let mut selections: Vec<String> = vec!();
    for index in indexes {
      selections.push(String::from(&self.branches[index]));
    }

    selections
  }
}
