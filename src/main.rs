mod blueprint;
mod scaffolder;
use blueprint::Blueprint;

fn main() {
    let config_path = "config.toml";
    let config_contents = std::fs::read(config_path).unwrap();
    let config_str = std::str::from_utf8(&config_contents).unwrap();
    let table: toml::Table = toml::from_str(config_str).unwrap();
    let blueprint = Blueprint::from_toml(&table).unwrap();
    println!("{blueprint:?}");
    scaffolder::scaffold(blueprint).unwrap();
}
