let score = JSON.parse(localStorage.getItem('score'));
        
if(score === null){
    score = {
        wins : 0,
        loses : 0,
        tie : 0
    }
}
updateScore();

function playGame(playerMove){
    let result = '';
    let computerMove = pickComputerMove();
    if(playerMove === 'rock'){
        if(computerMove === 'rock'){
            result = 'tie';
        }
        else if(computerMove === 'paper'){
            result = 'lose';
        }
        else{
            result = 'win';
        }
   }
   else if(playerMove === 'paper'){
    if(computerMove === 'rock'){
        result = 'win';
        }
        else if(computerMove === 'paper'){
            result = 'tie';
        }
        else{
            result = 'lose';
        }
    }
    else{
        if(computerMove === 'rock'){
            result = 'lose';
        }
        else if(computerMove === 'paper'){
            result = 'win';
        }
        else{
            result = 'tie';
        }
    }
    if(result === 'win'){
        score.wins++;
    }
    else if(result === 'lose'){
        score.loses++;
    }
    else{
        score.tie++;
    }

    localStorage.setItem('score', JSON.stringify(score));
    updateScore();
    document.querySelector('.js-result').innerHTML = `You ${result}.`;
    document.querySelector('.js-moves').innerHTML = `You <img src="images/${playerMove}-emoji.png" class="move-icon">
    <img src="images/${computerMove}-emoji.png" class="move-icon"> Computer`;
}


function pickComputerMove(){
    const val = Math.random();
    let computerMove = '';
    if(val < 1/3){
    computerMove = 'rock';
    }
    else if(val < 2/3){
        computerMove = 'paper';
    }
    else{
        computerMove = 'scissors';
    }

    return computerMove;
}

function updateScore(){
    document.querySelector('.js-score').innerHTML = 
`Wins : ${score.wins}, Loses : ${score.loses}, Tie : ${score.tie}`;
}
