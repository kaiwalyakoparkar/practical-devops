pub fn run() {
    let str = String::from("Hello World from loop");
    let mut count = 0;

    //for range loop statement
    for word in str.split_whitespace() {
        println!("{}", word);
    }
    println!("{}", str);

    //Infinite loop
    loop {
        count += 1;
        println!("{}", count);

        if count == 10 {
            break;
        }
    }

    //While loop
    count = 0;
    while count <= 10 {
        if count%3 == 0 {
            println!("{}", count);
        }
        count += 1;
    }
}