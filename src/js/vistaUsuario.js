this.service = 'ropa';
const logOffButton = document.querySelector('#logOff');

$(document).ready(function () {
    listar(service);

    var usuario = JSON.parse(sessionStorage.getItem('usuario'));
    var user_name = document.querySelector('#txt_user_name');
    var user_img = document.querySelector('#user_img');
    $(user_img).attr("src",usuario.FotoPerfil)
    user_name.innerHTML =usuario.Nombre + ' ' + usuario.PrimerApellido;


})
logOffButton.addEventListener('click', function () {
    sessionStorage.clear()
    setTimeout(() => {
        window.location = "http://localhost:3000/index.html";
    }, 2000);
})



