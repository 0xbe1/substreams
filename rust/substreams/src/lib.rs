mod externs;
pub mod log;
pub mod memory;
pub mod proto;
pub mod state;
pub mod rpc;
pub mod pb;

pub fn output<M: prost::Message>(msg: M) {
    // Need to return the buffer and forget about it issue occured when trying to write large data
    // wasm was "dropping" the data before we could write to it, which causes us to have garbage
    // value. By forgetting the data we can properly call external output function to write the
    // msg to heap.
    let (ptr, len, _buffer) = proto::encode_to_ptr(&msg).unwrap();
    std::mem::forget(&_buffer);
    unsafe { externs::output(ptr, len as u32) }
}

pub fn output_raw(data: Vec<u8>) {
    unsafe { externs::output(data.as_ptr(), data.len() as u32) }
}

pub fn hook(info: &std::panic::PanicInfo<'_>) {
    let error_msg = info
        .payload()
        .downcast_ref::<String>()
        .map(String::as_str)
        .or_else(|| info.payload().downcast_ref::<&'static str>().copied())
        .unwrap_or("");
    let location = info.location();

    unsafe {
        let _ = match location {
            Some(loc) => {
                let file = loc.file();
                let line = loc.line();
                let column = loc.column();

                externs::register_panic(
                    error_msg.as_ptr(),
                    error_msg.len() as u32,
                    file.as_ptr(),
                    file.len() as u32,
                    line,
                    column,
                )   
            }
            None => externs::register_panic(
                error_msg.as_ptr(),
                error_msg.len() as u32,
                std::ptr::null(),
                0,
                0,
                0,
            ),
        };
    }
}

pub fn register_panic_hook() {
    use std::sync::Once;
    static SET_HOOK: Once = Once::new();
    SET_HOOK.call_once(|| {
        std::panic::set_hook(Box::new(hook));
    });
}
