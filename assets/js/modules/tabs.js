const ELEMENTS_SELECTOR = {
  tabs: '[data-tabs]',
  tab: '[data-tab]',
  section: 'data-tab-section'
};

export default function () {
  document.querySelectorAll(ELEMENTS_SELECTOR.tab).forEach((elem) => {
    elem.addEventListener('click', () => {
      const id = elem.dataset.tab;

      const container = elem.closest(ELEMENTS_SELECTOR.tabs);

      container.querySelectorAll(ELEMENTS_SELECTOR.tab).forEach((elem) => {
        elem.classList.remove('tabs__item--active');
      });

      container.querySelectorAll(`[${ELEMENTS_SELECTOR.section}]`).forEach((elem) => {
        elem.classList.remove('tabs__block--active');
      });

      elem.classList.add('tabs__item--active');
      container.querySelector(`[${ELEMENTS_SELECTOR.section}='${id}']`).classList.add('tabs__block--active');
    });
  });
}
