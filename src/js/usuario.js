$(document).ready(function () {
    var user_img = document.querySelector('#user_img');
    var user_name = document.querySelector('#name');
    var usuario = JSON.parse(sessionStorage.getItem('usuario'));
    user_img.setAttribute("src", usuario.FotoPerfil);
    user_img.setAttribute("border", "none");
    user_name.innerHTML = usuario.Nombre + '\n' + usuario.PrimerApellido + '\n' + usuario.SegundoApellido;
})
    

