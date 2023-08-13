const ELEMENTS_SELECTOR = {
  card: '[data-card]',
  button: '[data-favourites]'
};

document.querySelectorAll(ELEMENTS_SELECTOR.button).forEach((elem, index) => {
  elem.addEventListener('click', () => {
    const container = elem.closest(ELEMENTS_SELECTOR.card);

    container.querySelector(ELEMENTS_SELECTOR.button).classList.toggle('favourites--fill');
  });
});
