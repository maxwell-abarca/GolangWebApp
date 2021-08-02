'use strict'
this.service = 'usuario';
const correo = document.querySelector('#txt_correo_electronico');
const contrasena = document.querySelector('#txt_contrasena');
const btn_inicio_sesion = document.querySelector('#btn_iniciar_sesion');


function validarFormulario() {
    let err = false;
    $('#formulario_inicio input').each(function () {
        if ($(this).val() == "") {
            $(this).addClass("is-invalid");
            err = true;
        } else {
            $(this).removeClass("is-invalid");
            $(this).addClass("is-valid");
            err = false;
        }
    })
    return err;
}
function sendData() {
    if (validarFormulario() == false) {
        var usuario = {
            CorreoElectronico: correo.value,
            Contrasena: contrasena.value
        }
        listarUsuario(usuario, service)
    }

}
btn_inicio_sesion.addEventListener('click', sendData)