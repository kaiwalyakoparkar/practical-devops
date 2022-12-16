const express = require('express');
const app = express();

app.get('/', (req, res, next) => {
    res.status(200).json({
        status: "Success",
        message: "This is working"
    })
})

app.listen(8080, ()=>{
    console.log("Application running at http://localhost:8080");
})