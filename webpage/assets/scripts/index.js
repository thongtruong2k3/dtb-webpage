
const navToggle = document.querySelector("nav .wrap-menu")
const navMenu = document.querySelector(".nav__navigation")
const body = document.getElementsByTagName("body")[0]


navToggle.addEventListener("click", () =>{
    navMenu.classList.toggle("nav__navigation_visible");

    body.classList.toggle("body__fixed");

    if (navMenu.classList.contains("nav__navigation_visible")){
        navToggle.setAttribute("aria-label", "Cerrar menú");
    } else {
        navToggle.setAttribute("aria-label", "Abrir menú");
    }
});


const images = document.querySelectorAll('#banner img');
let currentIndex = 0;

setInterval(() => {
  const nextIndex = (currentIndex + 1) % images.length;
  images[currentIndex].style.opacity = 0;
  images[nextIndex].style.opacity = 1;
  currentIndex = nextIndex;
}, 3000);

function search() {
    var searchTerm = document.getElementById('searchInput').value;
    alert('Searching for: ' + searchTerm);
    // You can perform further actions based on the search term
  }

ScrollReveal().reveal('.circle'); 
ScrollReveal().reveal('.footer'); 
ScrollReveal().reveal('.product-card');