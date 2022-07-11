include!("lib.rs");

use daemon::utils::{parser::parser_packet};
use tokio::{
    io::{AsyncReadExt, AsyncWriteExt},
    net::{TcpListener, TcpStream},
};

#[tokio::main]
async fn main() {
    let listener = TcpListener::bind("127.0.0.1:1314").await.unwrap();
    loop {
        let (socket, _) = listener.accept().await.unwrap();
        tokio::spawn(async move { handler(socket).await });
    }
}

async fn handler(mut socket: TcpStream) {
    let mut buf = [0u8; 2048];
    _ = socket.read(&mut buf).await;
    let data: String = String::from_utf8_lossy(&buf.to_vec()).to_string();
    let str = r#"{"_ACTION_":"MSG","_PARAM_":"global ","_DATA_":"hello worldsss"}"#;
    parser_packet(data.as_str());
    _ = socket.write_all(b"hello world!").await;
}
