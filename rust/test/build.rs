use std::io::Result;
fn main() -> Result<()> {
    let mut prost_build = prost_build::Config::new();
    prost_build.out_dir("./src/pb");
    prost_build.compile_protos(
        &["../../proto/sf/ethereum/substreams/v1/rpc.proto"],
        &["../../proto"],
    )
}
