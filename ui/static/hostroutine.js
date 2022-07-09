const quizSession = {
    quizId : document.getElementById("quiz-id").value,
    quizSlug: document.getElementById("quiz-code").value,
    stage: document.getElementById("quiz-stage").value -1
}

let liveResults = setInterval(()=> {
    
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
    })
    .catch(error => console.warn(error));
}, 4000)