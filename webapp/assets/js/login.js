$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
    evento.preventDefault();
    
    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(function(){
        console.log("oiaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
        window.location = "/home";
    }).fail(function(){
        alert("Usuário ou senha inválidos!");
    });
}