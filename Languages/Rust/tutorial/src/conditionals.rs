pub fn run() {
    let age = 21;
    let check_id = false;
    let knows_person_age = false;

    //If-else
    if age >= 21 && check_id || knows_person_age{
        println!("You are above legal drinking age");
    } else {
        println!("You are not above legal drinking age");
    }

    //A short hand technique
    let is_of_age = if age >= 21 {true} else {false};
    println!("{}", is_of_age);
}