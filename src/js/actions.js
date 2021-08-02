'use strict'
this.api = 'http://localhost:3000';
let registrar_usuario = (usuario, service) => {
    /* $.post(api, JSON.stringify(usuario), function (response) {
        console.log(response)
    }) */
    $.ajax({
        method: "POST",
        url: api + '/' + service,
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
let registrar_ropa = (ropa, service) => {
    $.ajax({
        type: "POST",
        url: api + '/' + service,
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(ropa),
        dataType: "text",
    })
}
let listarUsuario = (user, service) => {
   
    $.ajax({
        type: "GET",
        url: api + '/' + service,
        data: user,
        dataType: "json",
        success: function (response) {
            sessionStorage.setItem('usuario',JSON.stringify(response));
            setTimeout(() => {
                window.location = "http://localhost:3000/usuario.html";

            }, 2000);

        }
    });
}

let listar = (service) => {
    var usuario = JSON.parse(sessionStorage.getItem('usuario'));
    $.ajax({
        method: "GET",
        url: api + '/' + service,
        data: usuario.Cedula,
        dataType: "json",
        success: function (data) {
            for (let i = 0; i < data.length; i++) {
                var ropa = data[i];
                /*creates column*/
                var col = document.createElement('div');
                /*creates card*/
                var card = document.createElement('div');
                /*creates card body*/
                var card_body = document.createElement('div');
                /*creates image*/
                var img = document.createElement('img');
                /*creates h3*/
                var h3 = document.createElement('h3');
                var h2 = document.createElement('h2');

                $(card).addClass('card');
                $(card_body).addClass('card-body bg-primary text-light');
                $(col).addClass('col-md-4  mt-4 mb-5');
                $(h2).addClass('card-title');
                $(h3).addClass('card-text');


                $(img).addClass('card-img-top bg-light');

                $(img).attr("src", ropa.Imagen);
                $(img).attr("height", "230px;");


                $(card).addClass('mr-auto ml-auto w-50');


                $('#contenido').append(col);
                $(col).append(card);
                $(card).append(img);
                $(card).append(card_body);
                $(card_body).append(h2);
                $(card_body).append(h3);
                $(h2).append(ropa.Tipo);
                $(h3).append(ropa.Precio);




            }

        }
    });



}


