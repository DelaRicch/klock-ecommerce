/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx,vue}",
    ],
    theme: {
        extend: {
            backgroundImage: {
                'auth-bg': "url('/src/assets/auth-bg.png')",
            }
        },
    },
    plugins: [],
}