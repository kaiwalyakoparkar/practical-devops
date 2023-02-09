use std::env;
pub fn run() {
    // println!("Hello" );

    let args: Vec<String> = env::args().collect();
    let command = args[1].clone();
    let name = "Kaiwalya";
    let status = "100%";

    if command == "hello" {
        println!("Hi, {} how are you!", name);
    } else if command == "status" {
        println!("Current status is {}", status);
    } else {
        println!("That is not a valid command!");
    }
}