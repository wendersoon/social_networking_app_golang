$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

function criarPublicacao(evento){
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function(){
        window.location = "/home";
    }).fail(function(){
        Swal.fire({
            title: "Ops...!",
            text: "Erro ao criar a publicação",
            icon: "error"
          });
    })

}

function curtirPublicacao(evento) {
    evento.preventDefault();
    
    const elementoClicado = $(evento.currentTarget);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function(){
        const contadorDeCurtidas = elementoClicado.find('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');
    }).fail(function(){
        Swal.fire({
            title: "Ops...!",
            text: "Erro ao curtir a publicação",
            icon: "error"
        });
    }).always(function(){
        elementoClicado.prop('disabled', false);
    })
    
}

function descurtirPublicacao(evento){
    evento.preventDefault();
    
    const elementoClicado = $(evento.currentTarget);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function(){
        const contadorDeCurtidas = elementoClicado.find('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);
        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');
    }).fail(function(){
        Swal.fire({
            title: "Ops...!",
            text: "Erro ao descurtir a publicação",
            icon: "error"
        });
    }).always(function(){
        elementoClicado.prop('disabled', false);
    })
    
}

function atualizarPublicacao(evento) {
    $(this).prop('disable', true);

    const publicacaoId = $(this).data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function(){
        Swal.fire({
            title: "Sucesso!",
            text: "Publicação atualizada com sucesso",
            icon: "success"
          }).then(function(){
            window.location = "/home";
          })
    }).fail(function(){
        Swal.fire({
            title: "Ops...!",
            text: "Falha em atualizar publicação",
            icon: "error"
        });
    }).always(function(){
        $('#atualizar-publicacao').prop('disable', false);
    })
}

function deletarPublicacao(evento){
    evento.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir a publicação? É uma ação irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"

      }).then(function(confirmacao){
        if (!confirmacao.value) return;

        const elementoClicado = $(evento.target);
        const publicacao = elementoClicado.closest('div')
        const publicacaoId = publicacao.data('publicacao-id');
    
        elementoClicado.prop('disabled', true);
        console.log("aqui");
        $.ajax({
            url: `publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function(){
            publicacao.fadeOut("slow", function(){
                $(this).remove();
    
            })
        }).fail(function(){
            Swal.fire({
                title: "Ops...!",
                text: "Erro ao excluir a publicação!",
                icon: "error"
            });
        })
      })
    
}
