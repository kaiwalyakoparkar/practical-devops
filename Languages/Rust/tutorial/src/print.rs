//public function so that it can be accessed from outside
// We will try to access this run() function inside main.rs file
pub fn run() {
    // Prints to console
    println!("Hello from printrs file");

    //println!(1) will give error instead we have to specify the string resolver or something similar to template literal
    println!("{}", 1);
    println!("Hello my name is {}", "Kaiwalya");
    println!("{} is learning {}", "Kaiwalya", "Rust");

    //Multiple placeholders
    println!("{0} is learning {1}, {0} is {2}", "Kaiwalya", "Rust", "student");

    //Named arguments
    println!("{name} is learning {lang}", name="Kaiwalya", lang="rust");

    //Debug traits
    println!("{:?}", (12, true, "Hello"));

    //Basic maths
    println!("10 + 10 is {}", 10+10);
}