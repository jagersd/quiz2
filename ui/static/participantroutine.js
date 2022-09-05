const quizSession = {
    quizId : document.getElementById("quiz-id").value,
    stage: document.getElementById("quiz-stage").value -1,
    waitMessage: document.getElementById("waiting-message")
}

let callCounter = 0

let liveResults = setInterval(()=> {
    callCounter > 1000 ? clearInterval(liveResults) : callCounter += 1
    
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
    .catch(error => console.warn(error))
}, 4000)

function parseResults(result) {
    if (result == true){
        quizSession.waitMessage.style.display = "none"
        document.getElementById("options-form").style.display = "block"
    }

}

let submitBtn = document.querySelector('[type=submit]')

let radios = document.querySelectorAll('input[type=radio]')
if (radios){
  let checked = document.querySelectorAll('input[type=radio]:checked')

  if(!checked.length){
    submitBtn.style.opacity = "0.6"
    submitBtn.setAttribute("disabled", "disabled")
  }
  radios.forEach(function(el){
    el.addEventListener('click', function(){
      checked = document.querySelectorAll('input[type=radio]:checked')
      if(checked.length){
        submitBtn.style.opacity = "1"
        submitBtn.removeAttribute("disabled")
      }
    })
  })
}

let openAnswer = document.querySelector('input[type=text]')
if (openAnswer){
  openAnswer.addEventListener('change', function(){
    if (openAnswer.value == ""){
      submitBtn.style.opacity = "0.6"
      submitBtn.setAttribute("disabled", "disabled")
    } else {
      submitBtn.style.opacity = "1"
      submitBtn.removeAttribute("disabled")
    }
  })
}

