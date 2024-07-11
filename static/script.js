const themesContainer = document.querySelector('.theme-modes');

const toggleTheme = () => {
  document.querySelector('body').classList.toggle('dark-mode');
  document.querySelector('.container').classList.toggle('dark-mode');
  themesContainer.classList.toggle('active');
};

themesContainer.addEventListener('click', toggleTheme);
