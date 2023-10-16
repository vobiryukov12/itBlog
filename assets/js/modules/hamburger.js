const ELEMENTS_SELECTOR = {
  hamburger: '[data-hamburger]',
  nav: '[data-nav]'
};

document.querySelector(ELEMENTS_SELECTOR.hamburger).addEventListener('click', () => {
  document.querySelector(ELEMENTS_SELECTOR.nav).classList.toggle('nav--active');
});
