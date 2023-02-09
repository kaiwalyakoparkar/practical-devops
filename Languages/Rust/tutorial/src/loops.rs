pub fn run() {
    let str = String::from("Hello World from loop");

    //Loop statement
    for word in str.split_whitespace() {
        println!("{}", word);
    }
    println!("{}", str);
}