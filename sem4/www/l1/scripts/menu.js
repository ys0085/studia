document.addEventListener("DOMContentLoaded", function () {
    const hamburgerMenu = document.querySelector(".hamburger-menu");
    const nav = document.getElementById("main-nav");
    const menuItems = document.querySelectorAll("#main-nav ul li a");
    function checkScreenWidth() {
        if (window.innerWidth <= 768) {
            nav.classList.remove("active");
        } else {
            nav.classList.remove("active");
            nav.style.display = "block";
            hamburgerMenu.classList.remove("change");
        }
    }
    hamburgerMenu.addEventListener("click", function () {
        this.classList.toggle("change");
        nav.classList.toggle("active");
        if (nav.classList.contains("active")) {
            nav.style.display = "block";
        } else {
            nav.style.display = "none";
        }
    });
    menuItems.forEach(function (item) {
        item.addEventListener("click", function () {
            if (window.innerWidth <= 768) {
                nav.classList.remove("active");
                nav.style.display = "none";
                hamburgerMenu.classList.remove("change");
            }
        });
    });
    window.addEventListener("resize", function () {
        checkScreenWidth();
    });
    checkScreenWidth();
});