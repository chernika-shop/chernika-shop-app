/** @type {import('lint-staged').Config} */
const config = {
    // Сначала форматируем всё нужное Prettier'ом
    "*.{js,jsx,ts,tsx,css,scss,md,json}": "prettier --write",

    // Затем чиним JavaScript/TypeScript через ESLint
    "*.{js,jsx,ts,tsx}": "eslint --fix",

    // И отдельно чиним стили (CSS/SCSS модули) через Stylelint
    "**/*.module.{css,scss}": "stylelint --fix",
};

export default config;

