'use strict'

const slt_tipo = document.querySelector('#slt_tipo');
const txt_precio = document.querySelector('#txt_precio');
const txt_cantidad = document.querySelector('#txt_cantidad');
const slt_talla = document.querySelector('#slt_talla');
const btn_agregar_imagen = document.querySelector('#btn_agregar_imagen');
const btn_registrar = document.querySelector('#btn_registrar');
function agregarImagen() {
    cloudinary.openUploadWidget({
        user: 'maxwellab99gmail.com',
        cloud_name: 'drlznypvr',
        upload_preset: 'lonpzlw8',
        sources: ['local', 'url', 'camera', 'instagram', 'facebook']
    }, (error, result) => {
        if (!error && result && result.event === "success") {
            $('#clothing_img').attr("src", result.info.url)
        }
    }
    );

}
function validateAll() {
    let err = false;
    var regexNum = /^\d+$/;
    var regexTxt = "[a-zA-Z'.\\s]";
    $('#form input').each(function () {
        if ($(this).val() == "") {
            $(this).addClass("is-invalid");
            err = true;
        } else {
            $(this).removeClass("is-invalid");
            $(this).addClass("is-valid");
            err = false;
        }
    })
    $('#form input[type=number]').each(function () {
        if (!$(this).val().match(regexNum)) {
            $(this).addClass("is-invalid");
            err = true;
        } else {
            $(this).removeClass("is-invalid");
            $(this).addClass("is-valid");
            err = false;
        }
    })
    return err
}

function registrar() {
    var usuario = JSON.parse(sessionStorage.getItem('usuario'))


    if (validateAll() == false) {
        var ropa = {
            Tipo: slt_tipo.value,
            Precio: txt_precio.value,
            Talla: slt_talla.value,
            Propietario: usuario.Cedula,
            Cantidad: txt_cantidad.value,
            Imagen: $('#clothing_img').attr("src")
        };
        registrar_ropa(ropa, service)


    }
}
btn_registrar.addEventListener('click', registrar);
btn_agregar_imagen.addEventListener('click', agregarImagen);