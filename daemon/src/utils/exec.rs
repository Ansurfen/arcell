// use std::process::Command;
use std::env::{self};

pub fn home_dir() -> String {
    match env::home_dir() {
        Some(path) => path.display().to_string(),
        None => panic!("Impossible to get your home dir!"),
    }
}

// pub fn home_dir() -> String {
//     let cmd: String;
//     if cfg!(target_os = "windows") {
//         cmd = "echo %username%".to_string();
//     } else {
//         cmd = "echo ${HOME}".to_string();
//     }
//     if cfg!(target_os = "windows") {
//         let output = Command::new("cmd")
//             .arg("/c")
//             .arg(cmd)
//             .output()
//             .expect("Fail to get home dir")
//             .stdout;
//         let mut dir = "C:\\Users\\".to_string() + &String::from_utf8_lossy(&output).to_string();
//         dir.pop();
//         dir.pop();
//         return dir;
//     } else {
//         let output = Command::new("sh")
//             .arg("-c")
//             .arg(cmd)
//             .output()
//             .expect("Fail to get home dir")
//             .stdout;
//         return String::from_utf8_lossy(&output).to_string();
//     };
// }
