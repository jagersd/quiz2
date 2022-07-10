const quizSession = {
    quizId : document.getElementById("quiz-id").value,
    quizSlug: document.getElementById("quiz-code").value,
    stage: document.getElementById("quiz-stage").value -1
}

const ul = document.getElementById("player-list")
const ul2 = document.getElementById("subtotals")

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
      displayResults(responseData.results)
    })
    .catch(error => console.warn(error));
}, 4000)

function displayResults(results) {
    ul.innerHTML = "";
    results.forEach(resultLine => ul.innerHTML += `<li>${resultLine.PlayerName} : ${resultLine.Result != null ? resultLine.Result : ""}</>`);

    ul2.innerHTML = "";
    results.forEach(resultLine => ul2.innerHTML += `<li>${resultLine.PlayerName} : ${resultLine.Total}</>`);
  }
  