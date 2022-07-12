include!("lib.rs");

use daemon::utils::parser::parser_packet;

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
    let mut buf = vec![0; 2048];
    let n = socket.read(&mut buf).await.unwrap();
    let data: String = String::from_utf8_lossy(&buf).to_string();
    println!("{}", &data[..n]);
    parser_packet(&data[..n]);
    _ = socket.write_all(b"Execute successfully").await;
}
