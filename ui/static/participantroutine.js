const quizSession = {
    quizId : document.getElementById("quiz-id").value,
    stage: document.getElementById("quiz-stage").value -1,
    waitMessage: document.getElementById("waiting-message")
}

let callCounter = 0;

let liveResults = setInterval(()=> {
    callCounter > 1000 ? clearInterval(liveResults) : callCounter += 1;
    
    fetch(`/readytoreveal/${quizSession.quizId}/${quizSession.stage}`,
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
      parseResults(responseData.result)
    })
    .catch(error => console.warn(error));
}, 2000)

function parseResults(result) {
    if (result == true){
        quizSession.waitMessage.style.display = "none";
        document.getElementById("options-form").style.display = "block";
    }

}
  