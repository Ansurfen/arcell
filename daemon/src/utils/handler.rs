use std::{collections::HashMap, fs::File, io::Write, string};

use serde_json::{json, Value};

use crate::utils::exec::home_dir;

use super::io::{create_or_open, FStatus};

pub fn map(packet: Value, params: Vec<String>) {
    assert_ne!(params.len(), 2, "Params amount must be one");
    // "_DATA_": "k v"
    let data: Vec<&str> = match packet.get("_DATA_") {
        None => panic!("Data must not be empty"),
        Some(v) => v.as_str().unwrap().split(" ").collect(),
    };
    assert_eq!(data.len(), 2, "Data amount must be two");
    let pf = get_path_and_filename(params.clone(), "mapper");
    let mut map = HashMap::new();
    map.insert(data[0].to_string(), data[1].to_string());
    push_json(&pf.0.as_str(), &pf.1, map);
}

pub fn unmap(packet: Value, params: Vec<String>) {}

pub fn gmap(packet: Value, params: Vec<String>) {}

pub fn reg(packet: Value, params: Vec<String>) {}

pub fn unreg(packet: Value, params: Vec<String>) {}

pub fn msg(packet: Value, params: Vec<String>) {
    assert_ne!(params.len(), 2, "Params amount must be one");
    /* "_DATA_" : {
        {"k":"v"}, {"k2":"v2"}, ...
    }*/
    let datas_set = match packet.get("_DATA_") {
        None => panic!("Data must not be empty"),
        Some(v) => v.as_str().unwrap(),
    };
    let datas: serde_json::Value = serde_json::from_str(datas_set).unwrap();
    let obj = datas.as_object().unwrap().clone();
    let mut map: HashMap<String, String> = HashMap::new();
    for (k, v) in obj {
        map.insert(k, v.as_str().unwrap().to_string());
    }
    let pf = get_path_and_filename(params.clone(), "msg");
    push_json(&pf.0.as_str(), &pf.1, map);
}

pub fn gmsg(packet: Value, params: Vec<String>) {
    println!("{:?}", packet);
    println!("{:?}", params);
    // 根据Params循环拿链接数据，返回{"global":"","local":""}
}

pub fn get_param(packet: Value) -> Vec<String> {
    let params = match packet.get("_PARAM_") {
        None => "None",
        Some(v) => v.as_str().unwrap(),
    };
    assert_ne!(params, "None", "Fail to parser param");
    let param: Vec<&str> = params.split(" ").collect();
    let mut ret: Vec<String> = Vec::new();
    for p in param {
        if p != "" {
            ret.push(p.to_string());
        }
    }
    println!("{:?}", ret);
    return ret;
}

pub fn push_json(path: &str, filename: &str, mapper: HashMap<String, String>) {
    let mut fp = create_or_open(path, filename);
    match fp.1 {
        FStatus::CREATE => {
            let mut json: serde_json::Value = json!({});
            for (k, v) in mapper {
                json[k] = json!(v);
            }
            let res = serde_json::to_string_pretty(&json).expect("Fail to serialize json");
            fp.0.write_all(res.as_bytes()).expect("Fail to write file");
        }
        FStatus::OPEN => {
            let mut json: serde_json::Value = serde_json::from_reader(&fp.0).unwrap();
            for (k, v) in mapper {
                json[k] = json!(v);
            }
            let res = serde_json::to_string_pretty(&json).expect("Fail to serialize json");
            let mut wfp = File::create(path.to_owned() + filename).unwrap();
            wfp.write_all(res.as_bytes()).expect("Fail to write file");
        }
    }
}

pub fn pop_json(path: &str, key: Vec<String>) {}

pub fn get_path_and_filename(params: Vec<String>, dir: &str) -> (String, String) {
    let path = home_dir() + "\\.arcell\\" + dir + "\\";
    let mut target = "".to_string();
    if params[0] == "global" {
        target = "global.json".to_string();
    } else {
        target = params[0].clone() + ".json";
    }
    return (path, target);
}
