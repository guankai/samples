var socket;

$(document).ready(function () {
    // Create a socket
    socket = new WebSocket('ws://' + window.location.host + '/ws/join?uname=' + $('#uname').text());
    // Message received on the socket
    socket.onmessage = function (event) {
        var data = JSON.parse(event.data);
        console.log(data);
        switch (data.Type) {
        case 0: // JOIN
            // if (data.User == $('#uname').text()) {
            //     $("#chatbox li").first().before("<li>You joined the chat room.</li>");
            // } else {
            //     $("#chatbox li").first().before("<li>" + data.User + " joined the chat room.</li>");
            // }
            break;
        case 1: // LEAVE
            $("#chatbox li").first().before("<li>" + data.User + " left the chat room.</li>");
            break;
        case 2: // MESSAGE
            // $("#chatbox li").first().before("<li><b>" + data.User + "</b>: " + data.Content + "</li>");
            var chat = ''
            if (data.User == 'polly'){
                chat += '<article class="chat-item left">';
                chat += '<a href="#" class="pull-left thumb-sm avatar"><img src="/static/images/a2.png" alt="..."></a>';
                chat += '<section class="chat-body">';
                chat += '<div class="panel b-light text-sm m-b-none">';
                chat += '<div class="panel-body">';
                chat += '<span class="arrow left"></span>';
                chat += '<p class="m-b-none">' + data.Content + '</p>';
                chat += '</div>';
                chat += '</div>';
                chat += '<small class="text-muted"><i class="fa fa-ok text-success"></i> 2 minutes ago</small>';
                chat += '</section>';
                chat += '</article>';
            }else{
                chat += '<article class="chat-item right">';
                chat += '<a href="#" class="pull-right thumb-sm avatar"><img src="/static/images/a3.png" class="img-circle" alt="..."></a>';
                chat += '<section class="chat-body">';
                chat += '<div class="panel bg-light text-sm m-b-none">';
                chat += '<div class="panel-body">';
                chat += '<span class="arrow right"></span>';
                chat += '<p class="m-b-none">' + data.Content+ '</p>';
                chat += '</div>';
                chat += '</div>';
                chat += '<small class="text-muted">1 minutes ago</small>';
                chat += '</section>';
                chat += '</article>';
            }
            $("#polly").prepend(chat);
            break;
        }
    };

    // Send messages.
    var postConecnt = function () {
        var uname = $('#uname').text();
        var content = $('#sendbox').val();
        socket.send(content);
        $('#sendbox').val("");
    }

    $('#sendbtn').click(function () {
        postConecnt();
    });
});