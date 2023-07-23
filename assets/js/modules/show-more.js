const ELEMENTS_SELECTOR = {
  buttonArticle: '[data-button-more-article]',
  buttonNews: '[data-button-more-news]'
};

const buttonArticle = document.querySelector(ELEMENTS_SELECTOR.buttonArticle);
const buttonNews = document.querySelector(ELEMENTS_SELECTOR.buttonNews);

buttonArticle.addEventListener('click', () => {
  const hidden = document.querySelectorAll('.cards__hidden-article');
  for (let i = 0; i < 3 && i < hidden.length; i++) {
    hidden[i].classList.remove('cards__hidden-article');
  }

  if (hidden.length < 4) {
    buttonArticle.remove();
  }
});

buttonNews.addEventListener('click', () => {
  const hidden = document.querySelectorAll('.cards__hidden-news');
  for (let i = 0; i < 1 && i < hidden.length; i++) {
    hidden[i].classList.remove('cards__hidden-news');
  }

  if (hidden.length < 2) {
    buttonNews.remove();
  }
});
