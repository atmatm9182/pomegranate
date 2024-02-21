use std::{fs, io, path::Path};

use crate::blueprint::{Blueprint, Entry, FileSpec};

pub fn scaffold(blueprint: Blueprint<'_>) -> io::Result<()> {
    for entry in blueprint.dir_tree {
        create_entry("./", entry)?;
    }

    Ok(())
}

fn create_entry(parent_dir: impl AsRef<Path>, entry: Entry<'_>) -> io::Result<()> {
    match entry {
        Entry::File { spec, name } => match spec {
            FileSpec::Raw(contents) => {
                let full_path = parent_dir.as_ref().join(name);
                fs::write(full_path, contents)
            }
            FileSpec::File(path) => {
                let contents = fs::read(path)?;
                let full_path = parent_dir.as_ref().join(name);
                fs::write(full_path, contents)
            }
        },
        Entry::Directory { name, entries } => {
            fs::create_dir(name)?;
            
            let full_path = parent_dir.as_ref().join(name);
            for entry in entries {
                create_entry(&full_path, entry)?;
            }
            Ok(())
        }
    }
}
