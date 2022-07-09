const quizSession = {
  quizSlug : document.getElementById("quiz-slug").innerText,
  playerSlug: document.getElementById("player-slug").innerText,
  stage:1
}

window.localStorage.clear();
window.localStorage.setItem('quizSession', JSON.stringify(quizSession));

let quizStarted = false
let ul = document.getElementById("player-list")
let callCounter = 0;

let quizState = setInterval(()=> {

    quizStarted ? clearInterval(quizState) : quizStarted = false;
    callCounter > 100 ? clearInterval(quizState) : callCounter += 1;
    
    fetch(`/waitingroom/${quizSession.quizSlug}/${quizSession.playerSlug}`,
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
      createItems(responseData.players)
    })
    .catch(error => console.warn(error));
}, 4000)


function createItems(players) {
  ul.innerHTML = "";
  players.forEach(player => ul.innerHTML += `<li>${player}</>`);
}

async function startQuiz(){
  const response = await fetch("/startquiz",{
    method: "POST",
    headers: {'Content-Type':'application/json'},
    body:JSON.stringify(quizSession)
  })
  return response.json();
}
