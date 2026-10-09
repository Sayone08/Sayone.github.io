const toggle = document.querySelector('.menu-toggle');
const sidebar = document.querySelector('.sidebar');
const links = document.querySelectorAll('.nav-link');

toggle?.addEventListener('click', () => {
  const isOpen = sidebar.classList.toggle('open');
  toggle.setAttribute('aria-expanded', isOpen);
});

links.forEach((link) => {
  link.addEventListener('click', () => {
    sidebar.classList.remove('open');
    toggle?.setAttribute('aria-expanded', 'false');
  });
});

const sections = document.querySelectorAll('main section[id]');
const observer = new IntersectionObserver((entries) => {
  entries.forEach((entry) => {
    if (entry.isIntersecting) {
      entry.target.classList.add('is-visible');
      links.forEach((link) => link.classList.toggle('active', link.getAttribute('href') === `#${entry.target.id}`));
    }
  });
}, { rootMargin: '-35% 0px -55% 0px' });

sections.forEach((section) => observer.observe(section));

window.addEventListener('pointermove', (event) => {
  document.documentElement.style.setProperty('--mouse-x', `${event.clientX}px`);
  document.documentElement.style.setProperty('--mouse-y', `${event.clientY}px`);
});

const contactForm = document.querySelector('#contact-form');
const formStatus = document.querySelector('#form-status');

contactForm?.addEventListener('submit', async (event) => {
  event.preventDefault();
  formStatus.textContent = 'Sending securely through the Go service...';
  const response = await fetch('/api/contact', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(Object.fromEntries(new FormData(contactForm))),
  });
  const result = await response.json();
  formStatus.textContent = result.message;
  if (response.ok) contactForm.reset();
});
