
function custom_store() {
  localStorage.for_fun = "hi";
}

function pop_up() {
    var time = new Date();
    var abstime=time.getHours() +':'+time.getMinutes()+':'+time.getSeconds()
    alert(abstime);
}

function background_X() {
    document.body.style.backgroundImage = "url('back2.gif')";
    document.body.style.backgroundPosition = "left";
}

function new_window() {
    window.open("https://www.google.com","","width=200,height=100");
}

function spam() { 
    for (let i = 0; i < 1000; i++){
      console.log(i,"I did it");
    }
}