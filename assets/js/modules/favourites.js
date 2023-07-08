const ELEMENTS_SELECTOR = {
  card: '[data-card]',
  button: '[data-favourites]'
};

let arr;
const favourites = document.querySelectorAll(ELEMENTS_SELECTOR.button);

if (localStorage.getItem('favourites')) {
  arr = JSON.parse(localStorage.getItem('favourites'));
} else {
  arr = [];
  for (let i = 0; i < favourites.length; i++) {
    arr[i] = false;
  }
}

document.querySelectorAll(ELEMENTS_SELECTOR.button).forEach((elem, index) => {
  elem.addEventListener('click', () => {
    const container = elem.closest(ELEMENTS_SELECTOR.card);

    container.querySelector(ELEMENTS_SELECTOR.button).classList.toggle('favourites--fill');

    if (container.querySelector(ELEMENTS_SELECTOR.button).classList.contains('favourites--fill')) {
      arr[index] = true;
    } else {
      arr[index] = false;
    }

    localStorage.setItem('favourites', JSON.stringify(arr));
  });
});

for (let i = 0; i < arr.length; i++) {
  if (arr[i] === true) {
    favourites[i].classList.add('favourites--fill');
  }
}
