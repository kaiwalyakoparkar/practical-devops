pub fn run() {
    greeting("Hello", "Kaiwalya");

    //Bind functioun values to variables
    let get_sum = add(30, 50);
    println!("Your total bill is ${}", get_sum);
}

fn greeting(greet: &str, name: &str) {
    println!("{} {}, nice to meet you", greet, name);
}

fn add(n1: i32, n2: i32) -> i32 {
    //Remove semicolon if you want to return the value
    // n1+n2;
    n1+n2
}