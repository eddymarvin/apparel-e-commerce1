document.addEventListener('DOMContentLoaded', function () {
    // Carousel functionality
    const carousel = document.querySelector('.carousel');
    const items = document.querySelectorAll('.carousel-item');
    const prevButton = document.getElementById('prev');
    const nextButton = document.getElementById('next');
    let currentIndex = 0;
    const itemWidth = items[0].clientWidth;
  
    function updateCarouselPosition() {
      carousel.style.transform = `translateX(-${currentIndex * itemWidth}px)`;
    }
  
    nextButton.addEventListener('click', function () {
      if (currentIndex < items.length - 1) {
        currentIndex++;
      } else {
        currentIndex = 0; // Loop back to the first item
      }
      updateCarouselPosition();
    });
  
    prevButton.addEventListener('click', function () {
      if (currentIndex > 0) {
        currentIndex--;
      } else {
        currentIndex = items.length - 1; // Loop to the last item
      }
      updateCarouselPosition();
    });
  
    // Make sure the carousel adapts to screen resizing
    window.addEventListener('resize', function () {
      updateCarouselPosition();
    });
  });
  
  document.addEventListener("DOMContentLoaded", function () {
    const products = document.querySelectorAll(".product-card");

    // Simple hover effect to animate product cards
    products.forEach((product) => {
        product.addEventListener("mouseover", () => {
            product.style.transform = "translateY(-10px)";
            product.style.transition = "transform 0.3s ease";
        });

        product.addEventListener("mouseout", () => {
            product.style.transform = "translateY(0)";
        });
    });

    // Lazy loading effect (images load when scrolling into view)
    const productImages = document.querySelectorAll(".product-image img");

    const observer = new IntersectionObserver((entries, observer) => {
        entries.forEach((entry) => {
            if (entry.isIntersecting) {
                entry.target.src = entry.target.dataset.src;
                observer.unobserve(entry.target);
            }
        });
    });

    productImages.forEach((img) => {
        img.dataset.src = img.src;
        img.src = ""; // Reset the src for lazy load effect
        observer.observe(img);
    });
});
