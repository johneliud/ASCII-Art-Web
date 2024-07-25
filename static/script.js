const body = document.querySelector('body');
const containerEl = document.querySelector('.container');
const themeToggler = document.querySelector('.theme-modes');

const setTheme = (isDark) => {
  if (isDark) {
    body.classList.add('dark-mode');
    containerEl.classList.add('dark-mode');
    themeToggler.classList.add('active');
  } else {
    body.classList.remove('dark-mode');
    containerEl.classList.remove('dark-mode');
    themeToggler.classList.remove('active');
  }
  localStorage.setItem('darkMode', isDark);
};

const toggleTheme = () => {
  const isDark = !body.classList.contains('dark-mode');
  setTheme(isDark);
};

document.addEventListener('DOMContentLoaded', () => {
  const isDark = localStorage.getItem('darkMode') == 'true';
  setTheme(isDark);
});

themeToggler.addEventListener('click', toggleTheme);
