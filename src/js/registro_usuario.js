'use strict'
this.service = 'usuario';
const imagen = document.querySelector('#user_img');
const identificacion = document.querySelector('#txt_identificacion');
const nombre = document.querySelector('#txt_nombre')
const primer_apellido = document.querySelector('#txt_primer_apellido');
const segundo_apellido = document.querySelector('#txt_segundo_apellido');
const fecha_nacimiento = document.querySelector('#txt_fecha_nacimiento')
const correo = document.querySelector('#txt_correo_electronico');
const contrasena = document.querySelector('#txt_contrasena');
const contrasena2 = document.querySelector('#txt_contrasena_2');

const btn_agregar_imagen = document.querySelector('#btn_agregar_imagen');
const btn_siguiente = document.querySelector('#btn_siguiente');
const btn_registrar = document.querySelector('#btn_registrar');







function validateAll() {
    let error = false;
    var num = /^\d+$/;
    var txt = "[a-zA-Z'.\\s]"
    $('#first_form input').each(function () {
        if ($(this).val() == "") {
            error = true;
            $(this).css("border", "solid 2px #ff0000");
        } else {
            error = false;
            $(this).css("border", "solid 2px #ced4da");

        }
    })
    $('#first_form input[type=number]').each(function () {
        if (!$(this).val().match(num)) {
            error = true;
            $(this).css("border", "solid 2px #ff0000");
        } else {
            error = false;
            $(this).css("border", "solid 2px #ced4da");
        }
    })
    $('#first_form input[type=text]').each(function () {
        if (!$(this).val().match(txt)) {
            error = true;
            $(this).css("border", "solid 2px #ff0000");
        } else {
            error = false;
            $(this).css("border", "solid 2px #ced4da");
        }
    })
    return error;
}

function showSecond() {
    if (validateAll() == false) {
        $('#first_form').hide()
        $('#second_form').show()
    }
}
function validateAll2() {
    let error = false;
    var mail = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
    $('#second_form input').each(function () {
        if ($(this).val() == "") {
            error = true;
            $(this).css("border", "solid 2px #ff0000");
        } else {
            error = false;
            $(this).css("border", "solid 2px #ced4da");
        }
    })
    $('#second_form input[type=text]').each(function () {
        if (!$(this).val().match(mail)) {
            error = true;
            $(this).css("border", "solid 2px #ff0000");
        } else {
            error = false;
            $(this).css("border", "solid 2px #ced4da");
        }
    })
    $('#second_form input[type=password]').each(function () {
        if ($(this).val().length < 9) {
            error = true;
            $(this).css("border", "solid 2px #ff0000");
        } else {
            error = false;
            $(this).css("border", "solid 2px #ced4da");
        }
    })
    if (contrasena.value != contrasena2.value || contrasena2.value != contrasena.value) {
        error = true;
        $(contrasena, contrasena2).css("border", "solid 2px #ff0000");
    } else {
        error = false;
        $(contrasena, contrasena2).css("border", "solid 2px #ced4da");
    }

    return error;
}
function signIn() {
    var usuario = {}
    if (validateAll() == false && validateAll2() == false) {
        usuario.Cedula = identificacion.value;
        usuario.Nombre = nombre.value;
        usuario.PrimerApellido = primer_apellido.value;
        usuario.SegundoApellido = segundo_apellido.value;
        usuario.FechaNacimiento = fecha_nacimiento.value;
        usuario.CorreoElectronico = correo.value;
        usuario.Contrasena = contrasena.value;
        usuario.FotoPerfil = imagen.getAttribute("src");
        console.log(usuario);
        registrar_usuario(usuario,service);
    }
}



btn_agregar_imagen.addEventListener('click', function () {
    var mycloud = cloudinary.openUploadWidget({
        user: 'maxwellab99gmail.com',
        cloud_name: 'drlznypvr',
        upload_preset: 'lonpzlw8',
        sources: ['local', 'url', 'camera', 'instagram', 'facebook']
    }, (error, result) => {
        if (!error && result && result.event === "success") {
            $('#user_img').attr("src", result.info.url)

            $('#user_img').css("border", "none")

        }
    })
    mycloud.open();
})


btn_siguiente.addEventListener("click", showSecond)
btn_registrar.addEventListener("click", signIn)


$(document).ready(function () {
    $('#second_form').hide()
})
