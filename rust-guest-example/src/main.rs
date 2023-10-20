
#[link(wasm_import_module = "log")]
extern "C" {
    #[link_name = "log_count"]
    pub fn add(x: u32, y: u32) -> u32;
    #[link_name = "log_println"]
    pub fn Println(ptr: * const u8, len: usize);
}

fn main() {
    println!("Hello World");
    
    let blah: String = "Hello, World".to_string();
    unsafe {
        println!("{}", add(2, 9934848));
        Println(blah.as_ptr(), blah.len())
    }
}
