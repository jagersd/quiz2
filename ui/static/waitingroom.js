let quizSlug = document.getElementById("quiz-slug").innerText
let playerSlug = document.getElementById("player-slug").innerText
let quizStarted = false

let quizState = setInterval(()=> {

    quizStarted ? clearInterval(quizState) : quizStarted = false
    
    fetch(`http://localhost:3000/waitingroom/${quizSlug}/${playerSlug}`,
    {
    	method: "GET",
        headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        },
    })
    .then((response) => response.json())
    .then((responseData) => {
      quizStarted = responseData.quiz.Started;
      console.log(quizStarted)
    })
    .then(quizStarted => quizStarted==true ? clearInterval(quizState) : null)
    .catch(error => console.warn(error));
}, 2000)


