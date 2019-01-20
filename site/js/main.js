
(function ($) {
    "use strict";


    /*==================================================================
    [ Focus Contact2 ]*/
    $('.input2').each(function(){
        $(this).on('blur', function(){
            if($(this).val().trim() != "") {
                $(this).addClass('has-val');
            }
            else {
                $(this).removeClass('has-val');
            }
        })    
    })
            
  
    
    /*==================================================================
    [ Validate ]*/
   

    $('.validate-form').on('submit',function(e){
        console.log(this);
        e.preventDefault();

        var zipcode = $('#zipcode').val();

        $.get( "http://localhost:8080/shipments/zipcode/"+zipcode, function() {}) 
        .done(function(data) {
            $('#name').val(data.name);          
            $('#street').val(data.street);
            $('#neighborhood').val(data.neighborhood);
            $('#zipcode').val(data.zipcode);
            $('#city').val(data.city);

            $('#state_short').val(data.state_short);

        })
          .fail(function(data) {
                
                    console.log(data)
                    alert(data.status+' - '+ data.responseJSON.message );
               
            
          });
    
    });

    
    

})(jQuery);