// Set up theme toggle functionality after DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    const elements = {
        themeToggle: document.getElementById('theme-toggle'),
        themeToggleMobile: document.getElementById('theme-toggle-mobile'),
        hamburger: document.getElementById('hamburger'),
        mobileNav: document.querySelector('.mobile-nav')
    };

    // Theme management
    function isLightTheme() {
        return document.documentElement.getAttribute('data-theme') === 'light';
    }

    function updateThemeUI() {
        const isLight = isLightTheme();
        elements.themeToggle.innerHTML = isLight ? '🌙' : '☀️';
        elements.themeToggleMobile.textContent = isLight ? 'DARK' : 'LIGHT';
    }

    function toggleTheme() {
        if (isLightTheme()) {
            document.documentElement.removeAttribute('data-theme');
            localStorage.removeItem('theme');
        } else {
            document.documentElement.setAttribute('data-theme', 'light');
            localStorage.setItem('theme', 'light');
        }
        updateThemeUI();
    }

    // Mobile navigation
    function toggleMenu() {
        const isOpen = elements.mobileNav.classList.contains('active');

        elements.hamburger.classList.toggle('active', !isOpen);
        elements.mobileNav.classList.toggle('active', !isOpen);
        document.body.style.overflow = isOpen ? '' : 'hidden';
    }

    // Event listeners
    elements.themeToggle.addEventListener('click', toggleTheme);
    elements.themeToggleMobile.addEventListener('click', toggleTheme);
    elements.hamburger.addEventListener('click', toggleMenu);

    // Close menu when clicking navigation links
    elements.mobileNav.querySelectorAll('.mobile-menu a').forEach(link => {
        link.addEventListener('click', toggleMenu);
    });

    // Initialize theme UI
    updateThemeUI();
});