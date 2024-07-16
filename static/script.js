const body = document.querySelector('body');
const containerEl = document.querySelector('.container');
const themeToggler = document.querySelector('.theme-modes');

const toggleTheme = () => {
  body.classList.toggle('dark-mode');
  containerEl.classList.toggle('dark-mode');
  themeToggler.classList.toggle('active');
};

themeToggler.addEventListener('click', toggleTheme);
