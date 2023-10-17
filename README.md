# itBlog
### Ссылка на прототип:

https://vobiryukov12.github.io/itBlog/

## Preview

<img src='./assets/images/it-blog.gif' width='450'>

## Описание

Блог о технологиях и программировании со статьями и новостями.

Проект находиться на стадии разработки. Сейчас реализована только базовая верстка основных страниц и некоторый функционал. Сайт разрабатывается совместно с  [backend-разработчиком Stepan](https://github.com/Zemavong), который в настоящее время делает API, к которому клиентская часть будет делать запросы.

В дальнейшем планируется сделать административную панель для сайта, через которую можно будет добавлять, удалять и редактировать посты.

### Технологии используемые на проекте:

Frontend: <br><br>
[![Skills](https://skillicons.dev/icons?i=js,html,scss,webpack)](https://skillicons.dev)

Backend: <br><br> 
[![Skills](https://skillicons.dev/icons?i=go,postgres)](https://skillicons.dev)

### В проекте реализовано:

- Каркас для всех страниц:
  - Хедер содержит логотип, а также кнопки для перехода на страницы "Посты", "Избранное" и блок "Подписаться"
  - Футер 

- Главная страница:
  - Табы для переключения между категориями 
  - Карточки двух видов. Наполнение карточек меняется в зависимости от категории и разрешения экрана. Карточки содержат: заголовок, описание, дату публикации, кнопку для добавления в избранное, кол-во лайков, комментариев и просмотров, название категории
  - Кнопка для загрузки дополнительных карточек
  - Блок с формой подписки 

- Страница поста:
  - Каждая страница поста содержит заголовок, дату публикации, фотографию и содержание

### Этапы работы над проектом
- Первичные настройки, необходимые для начала работы:
  - Анализ дизайн-макета и вынесение в базовые scss-переменные значения цветов, размеров шрифта и т.д., которые будут переиспользоваться на сайте в различных местах
  - Приведение браузерных стилей к единому виду с помощью пакета 'normalize'
  - Подключение необходимых функций и шрифтов
  - Настройка webpack

- Базовая стилизация на тэги (заголовки, параграфы, нумерованные и ненумерованные списки, cсылки, изображения). 
  - Для начала была проведена стилизация основных часто используемых тэгов. Стилизация, описанная на тэги, может, например, применяться на тех страницах и в тех блоках, контент которых будет наполняться в административной части сайта через визуальный редактор. Соответственно, в разметке для такого контента не будут использоваться классы.
  - Стилизация на тэги описана в файле scss/base/_generic.scss.

- Верстка GUI
  - [GUI](https://vobiryukov12.github.io/itBlog/gui.html) - это набор всех блоков, элементов, которые встречаются на всем сайте. Начиная с самых простых - иконки, списки, заканчивая сложными - карточки и т.д. Правильное создание и ведение GUI существенно облегчает верстку.

- Верстка базовых блоков (иконки, заголовки, ссылки, кнопки, элементы форм).
  - После того, как произведена стилизация на тэги, создали БЭМ-блоки и описали стилизацию для самых основных и простых компонентов сайта. Для каждого БЭМ-блока завели свой отдельный файл. Название файла в точности совпадает с названием блока. Все файлы БЭМ-блоков лежат в scss/blocks/.
  - Далее настроили сборку символьного спрайта. Все иконки для символьного спрайта кладутся в assets/icons/svg/ и подключаются в файл assets/icons/icons.js. Внутри иконок стили, отвечающие за цвет заменены на currentColor (кроме тех иконок, у которых цвет должен быть неизменным).

- Верстка шаблонов сайта (каркас под сквозные элементы сайта и рабочую область, хедер, футер)
  - Для реализации каркаса разбили страницу таким образом, чтобы у всех страниц были сквозные элементы, а в рабочей области содержимое менялось в зависимости от конкретной страницы.

- Верстка блоков для содержимого страниц (табы, карточки)
  - После выполнения предыдущих задач полностью были готовы базовые блоки и весь лэйаут для сайта. После этого приступили к верстке тех блоков, которые потребуются для сборки конкретных страниц.

- Сборка страниц (главная страница, cтраница поста)
  - Когда уже сверстаны все необходимые блоки, приступили к сборке и верстке уже конкретных страниц.

### Особенности проекта:

- Дизайн сайта имеет 3 версии: мобильную: 375-767px, планшетную: 768-1279px, десктопную: 1440-1920px.
- Верстка осуществлялась по методологии БЭМ. 
- Для иконок собирается общий символьный спрайт.
- Стилизация компонентов производится по принципу mobile-first.
- Сайт сверстан с использованием специальных семантических тегов.
- В проекте используется модульная система.
- Вся верстка выполняется в относительных единицах измерения, поэтому была реализована функция, которая будут переводить значения из px в rem.

## Структура проекта
### Исходники
Иконки, стили, и скрипты находятся в /assets/. Включает поддиректории: fonts, icons, images, js, и scss.

### Готовые файлы 
Исходники из /assets/ компилируются в /public/.

### Сборщик
На проекте используется сборщик Webpack.

### Препроцессор
Для написания стилей используется препроцессор SASS (синтаксис SCSS).

### БЭМ
Стили пишутся по методологии БЭМ. Внутри блока описывается внутренняя геометрия и стилизация. Для указания внешней геометрии (позиционирования, отступов и т.д.) используем миксы. Ни один блок не влияет на другие блоки и их элементы. 

### Структура стилей
Исходники стилей находятся в поддиректории scss/, структура данной поддиректории выглядит следующим образом:
  - abstract/ (функции, миксины, переменные)
  - base/ (шрифты и базовая стилизация)
  - blocks/ (стилизация компонентов)
  - style.scss (главный файл, импортирующий остальные).

Сначала в style.scss подключаются стили плагинов, затем стили "абстракций", после - базовые стили, а в завершение всего стили компонентов.

### CSS-свойства
В файлах scss свойства объединяются по смыслу, порядок следования соответствует [htmlacademy/codeguide](https://codeguide.academy/html-css.html#css). 
