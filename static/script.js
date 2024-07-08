const themesContainer = document.querySelector(".theme-modes");
const lightThemeBtn = document.querySelector(".theme-modes .sunny")
const darkThemeBtn = document.querySelector(".theme-modes .moon")

const toggleTheme = () => {
    document.querySelector("body").classList.toggle("dark-mode")
    document.querySelector(".container").classList.toggle("dark-mode")

    lightThemeBtn.classList.toggle("active")
    darkThemeBtn.classList.toggle("inactive")
}

themesContainer.addEventListener("click", toggleTheme)