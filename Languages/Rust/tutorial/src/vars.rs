//Variables hold primitive data or references to data
//Variables are immutable by default
//Rust is block-scoped lanugage

pub fn run() {
    let name = "Kaiwalya";
    let age = 19;
    let mut pencil = 3;

    println!("My name is {} and I am {} years old and has {} pencils", name, age, pencil);

    //Generates error coz it's trying the change the variable value
    //age = 30;

    //Will not generate error coz of mut (mutable option added to it)
    pencil = 5;
    println!("But now he has {} pencils", pencil);

    //Contants (We need to provide the type if we are using constants)
    const ID: i32 = 001;
    println!("ID: {}", ID);
}