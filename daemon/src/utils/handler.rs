use serde_json::Value;

pub fn map(packet: Value, params: Vec<String>) {
    match packet.get("_DATA_") {
        None => {}
        Some(v) => {}
    }
}

pub fn map_helper() {}

pub fn unmap(packet: Value, params: Vec<String>) {}

pub fn gmap(packet: Value, params: Vec<String>) {}

pub fn msg(packet: Value, params: Vec<String>) {
    print!("啊这");
    println!("{:?}xxx", packet);
    println!("{:?}x", params);
    // 根据Params循环拿链接数据，返回{"global":"","local":""}
}

pub fn get_param(packet: Value) -> Vec<String> {
    let params: String = match packet.get("_PARAM_") {
        None => "None".to_string(),
        Some(v) => v.to_string(),
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
