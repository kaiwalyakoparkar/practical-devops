/*
Primitive Types--
Integers: u8, i8, u16, i16, u32, i32, u64, i64, u128, i128 (number of bits they take in memory)
Floats: f32, f64
Boolean (bool)
Characters (char)
Tuples
Arrays
*/

pub fn run() {
    //Default is i32 for int
    let x = 1;

    //Default is f64 for float
    let y = 2.3;

    //Boolean
    let is_active = true;
    let explicit_boolean: bool = true;
    let eval_boolean: bool = 10<5;

    //Character
    let a1 = 'a';
    let face = '\u{1f600}';//smiley emoji unicode

    println!("{:?}", (is_active, explicit_boolean, eval_boolean, a1, face));
}
