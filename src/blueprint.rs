use toml::Value;

#[derive(Debug)]
pub enum FileSpec<'a> {
    Raw(&'a str),
    File(&'a str),
}

impl<'a> FileSpec<'a> {
    fn parse(name: &str, table: &'a toml::Table) -> Result<Self, Error> {
        match table.get("raw") {
            Some(contents) => {
                let contents = contents.as_str().ok_or(Error::wrong_key_type(
                    name,
                    "raw",
                    vec!["string"],
                    contents.type_str(),
                ))?;
                Ok(FileSpec::Raw(contents))
            }
            None => match table.get("file") {
                Some(path) => {
                    let path = path.as_str().ok_or(Error::wrong_key_type(
                        name,
                        "raw",
                        vec!["string"],
                        path.type_str(),
                    ))?;
                    Ok(FileSpec::File(path))
                }
                None => Err(Error::missing_keys(name, vec!["raw", "file"])),
            },
        }
    }
}

#[derive(Debug)]
pub enum Entry<'a> {
    File {
        name: &'a str,
        spec: FileSpec<'a>,
    },
    Directory {
        name: &'a str,
        entries: Vec<Entry<'a>>,
    },
}

impl<'a> Entry<'a> {
    fn parse(name: &'a str, table: &'a toml::Value) -> Result<Self, Error> {
        match table {
            Value::Array(arr) => Self::parse_dir(name, arr),
            Value::Table(table) => Self::parse_file(name, table),
            e => Err(Error::wrong_key_type(
                "files",
                name,
                vec!["array", "table"],
                e.type_str(),
            )),
        }
    }

    fn parse_file(name: &'a str, table: &'a toml::Table) -> Result<Self, Error> {
        let spec = FileSpec::parse(name, table)?;
        Ok(Entry::File { name, spec })
    }

    fn parse_dir(name: &'a str, entries: &'a Vec<toml::Value>) -> Result<Self, Error> {
        let mut entries_vec = vec![];
        for spec in entries {
            let entry = Entry::parse(name, spec)?;
            entries_vec.push(entry);
        }

        Ok(Self::Directory {
            name,
            entries: entries_vec,
        })
    }
}

impl Default for Entry<'_> {
    fn default() -> Self {
        Self::File {
            name: "<empty>",
            spec: FileSpec::Raw("<empty>"),
        }
    }
}

#[derive(Debug, Default)]
pub struct Blueprint<'a> {
    name: &'a str,
    dir_tree: Vec<Entry<'a>>,
}

#[derive(Debug)]
enum BlueprintItem {
    Section(String),
    Key(String, String),
}

#[derive(Debug)]
pub enum Error {
    MissingSection(String),
    MissingKeys {
        section: String,
        keys: Vec<String>,
    },
    WrongType {
        item: BlueprintItem,
        expected: Vec<String>,
        got: String,
    },
    WrongKeyValue {
        section: String,
        name: String,
        value: String,
    },
}

impl Error {
    fn wrong_section_type(
        name: impl Into<String>,
        expected: Vec<impl Into<String>>,
        got: impl Into<String>,
    ) -> Self {
        Self::WrongType {
            item: BlueprintItem::Section(name.into()),
            expected: expected.into_iter().map(|t| t.into()).collect(),
            got: got.into(),
        }
    }

    fn wrong_key_type(
        section: impl Into<String>,
        name: impl Into<String>,
        expected: Vec<impl Into<String>>,
        got: impl Into<String>,
    ) -> Self {
        Self::WrongType {
            item: BlueprintItem::Key(section.into(), name.into()),
            expected: expected.into_iter().map(|t| t.into()).collect(),
            got: got.into(),
        }
    }

    fn missing_keys(section: impl Into<String>, keys: Vec<impl Into<String>>) -> Self {
        Self::MissingKeys {
            section: section.into(),
            keys: keys.into_iter().map(|k| k.into()).collect(),
        }
    }

    fn wrong_key_value(
        section: impl Into<String>,
        name: impl Into<String>,
        value: impl Into<String>,
    ) -> Self {
        Self::WrongKeyValue {
            section: section.into(),
            name: name.into(),
            value: value.into(),
        }
    }
}

impl<'a> Blueprint<'a> {
    pub fn from_toml(table: &'a toml::Table) -> Result<Self, Error> {
        let blueprint_section = table
            .get("blueprint")
            .ok_or(Error::MissingSection("blueprint".into()))?;
        let blueprint_section = blueprint_section
            .as_table()
            .ok_or(Error::wrong_section_type(
                "blueprint",
                vec!["table"],
                blueprint_section.type_str(),
            ))?;

        let name = blueprint_section
            .get("name")
            .ok_or(Error::missing_keys("blueprint", vec!["name"]))?;
        let name = name.as_str().ok_or(Error::wrong_key_type(
            "blueprint",
            "name",
            vec!["string"],
            name.type_str(),
        ))?;

        let files_section = table
            .get("files")
            .ok_or(Error::MissingSection("files".into()))?;
        let files_section = files_section.as_table().ok_or(Error::wrong_section_type(
            "files",
            vec!["table"],
            files_section.type_str(),
        ))?;

        let mut dir_tree = vec![];
        for (file_name, file_spec) in files_section {
            let entry = Entry::parse(file_name, file_spec)?;
            dir_tree.push(entry);
        }

        Ok(Blueprint { name, dir_tree })
    }
}
