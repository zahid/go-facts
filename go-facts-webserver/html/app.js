$(document).ready(function() {

    $(window).resize(function(){
        let left = ($(window).width() - $('.text').outerWidth()) /2 ;
        let top = ($(window).height() - $('.text').outerHeight())/2;
        console.log(left, top)
        $('.text').css({
            position:'absolute',
            left: left,
            top: top
        });

    });

    $.getJSON('http://localhost/api/fact', function(data) {
        $(".text").text(data.text);
        $("body").css("background-image", "url(" + data.image + ".jpg)");
        $(window).resize();
    });
    
})
