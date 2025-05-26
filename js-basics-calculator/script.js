
const display = document.getElementById("display");

function addToDisplay(input){
    display.value += input;
}

function showResult(){
    try{
        display.value = eval(display.value);   
    }
    catch(error){
        display.value = "Error"
    }
}

function clearDisplay(){
    display.value = "";
}