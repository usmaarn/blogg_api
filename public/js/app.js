document.addEventListener("DOMContentLoaded", (e) => {
  const siteYearEl = document.getElementById("site-year");
  if (siteYearEl) {
    siteYearEl.textContent = new Date().getFullYear();
  }
});
