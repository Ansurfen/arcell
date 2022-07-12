use crate::utils::handler::unmap;

use super::handler::{get_param, gmap, map, gmsg, msg};

extern crate serde_json;

pub fn parser_packet(data: &str) {
    let packet: serde_json::Value = serde_json::from_str(data).unwrap();
    let action = match packet.get("_ACTION_") {
        None => "None",
        Some(v) => v.as_str().unwrap(),
    };
    let params = get_param(packet.clone());
    match action {
        "MAP" => map(packet.clone(), params.clone()),
        "UNMAP" => unmap(packet.clone(), params.clone()),
        "GMAP" => gmap(packet.clone(), params.clone()),
        "MSG" => msg(packet.clone(), params.clone()),
        "GMSG" => gmsg(packet.clone(), params.clone()),
        _ => assert_ne!(action, "None", "Fail to parser action"),
    }
}

pub fn parser_args(data: String) -> String {
    let json: serde_json::Value = serde_json::from_str(&data).unwrap();
    let res = match json.get("args") {
        None => "None".to_string(),
        Some(v) => v.to_string(),
    };
    assert_ne!(res, "None", "Fail to find args");
    let args: Vec<&str> = res.split(" ").collect();
    let mut tmpl = &json;
    for arg in args {
        match tmpl.get(arg) {
            None => break,
            Some(v) => tmpl = v,
        }
    }
    tmpl.to_string()
}

#[test]
fn test_parser_args() {
    let str = "{
        \"args\":\"tell a client xxx\",
        \"tell\":\"\"
    }";
    parser_args(str.to_string());
}

#[test]
fn test_parser_packet() {
    let str = r#"{"_ACTION_":"MSG","_DATA_":"{\"xxx\":\"xx\",\"yyy\":\"yy\"}","_PARAM_":"global"}"#;
    parser_packet(str);
}
