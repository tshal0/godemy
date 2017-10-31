var OnLoad = function(){

    $('#register').on('submit', function(e){
        e.preventDefault();
        e.stopPropagation();

        var data = $("#register").serialize()
        $.ajax({
            type: "GET", 
            url : "register",
            data: data
        }).done(function(msg)  {
            console.log(msg);
        }).fail(function(msg)  {
            console.log(msg);
        }); 
    })
    
    $('#login').submit(function(){
        console.log("test")
    })

    $('body').bind('beforeunload',function(e){
        console.log('wtf')
     });

}

$(document).ready(OnLoad)