extern crate serde_yaml;

use std::{
    fs::{create_dir_all, File, OpenOptions},
    path::Path,
};

pub enum FStatus {
    CREATE,
    OPEN,
}

pub fn get_conf() {
    let fp = File::open("application.yml").unwrap();
    let conf: serde_yaml::Value = serde_yaml::from_reader(fp).unwrap();
    println!("{:?}", conf);
}

pub fn fexist(path: &str) -> bool {
    Path::new(path).exists()
}

pub fn create_or_open(path: &str, filename: &str) -> (File, FStatus) {
    if !fexist(path.clone()) {
        _ = create_dir_all(path);
    }
    let target = path.to_owned() + filename;
    if !fexist(&target) {
        let fp = std::fs::File::create(target).expect("Fail to create file");
        return (fp, FStatus::CREATE);
    }
    return (
        OpenOptions::new()
            .read(true)
            .open(target)
            .unwrap(),
        FStatus::OPEN,
    );
}
