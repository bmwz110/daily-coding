var main = function() {
    $(".more-btn").on("click", event => {
        $(event.currentTarget).sibling(".more-menu").toggle();
    });
    $(".share").on("click", event => {
        $(event.currentTarget).next(".share-menu").toggle();
    });
    $(".notification").on("click", event => {
        $(event.currentTarget).toggleClass("active");
    })
};

$(document).ready(main);