'use strict'
this.service = 'registro_usuario';
this.api = 'http://localhost:3000' + '/' + service;
let registrar_usuario = (usuario) => {
    /* $.post(api, JSON.stringify(usuario), function (response) {
        console.log(response)
    }) */
    $.ajax({
        method: "POST",
        url: api,
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(usuario),
        /* `data: JSON.stringify(usuario)` */
        dataType: "text"
    })
        .done(function (response) {
            var message = JSON.parse(response)
            Swal.fire({
                title: message.Message,
                text: message.Message,
                icon: 'success',
                confirmButtonText: 'Cool'
            })
            sessionStorage.setItem("usuario", JSON.stringify(usuario))
            setTimeout(() => {
                window.location = "http://localhost:3000/usuario.html";
                
            }, 2000);

        })
        .fail(function (jqxhr, textStatus, errowThrown) {
            console.log(jqxhr)
            console.log(textStatus)
            console.log(errowThrown)
        })
    
}
