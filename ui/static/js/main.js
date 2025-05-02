// Prevent flash by setting theme immediately before page load
(function() {
    var savedTheme = localStorage.getItem('theme');
    if (savedTheme === 'light') {
        document.documentElement.setAttribute('data-theme', 'light');
    }
    // Dark mode is the default (no attribute needed)
})();

// Set up theme toggle functionality after DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    const themeToggle = document.getElementById('theme-toggle');

    // Set icon based on current theme
    if (localStorage.getItem('theme') === 'light') {
        themeToggle.innerHTML = '🌙'; // Moon for light mode (to switch to dark)
    } else {
        // Default to dark mode
        themeToggle.innerHTML = '☀️'; // Sun for dark mode (to switch to light)
    }

    // Toggle theme on click
    themeToggle.addEventListener('click', () => {
        const currentTheme = document.documentElement.getAttribute('data-theme');

        if (currentTheme === 'light') {
            document.documentElement.removeAttribute('data-theme'); // Default is dark
            localStorage.removeItem('theme'); // Clear preference to default to dark
            themeToggle.innerHTML = '☀️';
        } else {
            document.documentElement.setAttribute('data-theme', 'light');
            localStorage.setItem('theme', 'light');
            themeToggle.innerHTML = '🌙';
        }
    });
});