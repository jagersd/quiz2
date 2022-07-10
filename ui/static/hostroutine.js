const quizSession = {
    quizId : document.getElementById("quiz-id").value,
    quizSlug: document.getElementById("quiz-code").value,
    stage: document.getElementById("quiz-stage").value -1
}

const ul = document.getElementById("player-list")
const ul2 = document.getElementById("subtotals")
let callCounter = 0;

let liveResults = setInterval(()=> {
    callCounter > 1000 ? clearInterval(liveResults) : callCounter += 1;
    
    fetch(`/liveresults/${quizSession.quizId}/${quizSession.quizSlug}/${quizSession.stage}`,
    {
      method: "GET",
      headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      },
    })
    .then((response) => response.json())
    .then((responseData) => {
      console.log(responseData)
      parseResults(responseData.results)
    })
    .catch(error => console.warn(error));
}, 4000)

function parseResults(results) {
    let readyToMove =[];
    ul.innerHTML = "";
    ul2.innerHTML = "";
    results.forEach(resultLine => {
        readyToMove += resultLine.Result;
        ul.innerHTML += `<li>${resultLine.PlayerName} : ${resultLine.Result != null ? resultLine.Result : ""}</>`;
        ul2.innerHTML += `<li>${resultLine.PlayerName} : ${resultLine.Total}</>`;
    });

    if(!readyToMove.includes(null) || callCounter >= 1000){
        document.getElementById("next-question-btn").style.display = "block";
        clearInterval(liveResults) 
    }

}
  