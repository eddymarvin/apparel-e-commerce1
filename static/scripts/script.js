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
  