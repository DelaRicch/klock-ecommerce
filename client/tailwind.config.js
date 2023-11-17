/** @type {import('tailwindcss').Config} */

import colors from "tailwindcss/colors";

export default {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx,vue}",
        "./node_modules/vue-tailwind-datepicker/**/*.js"
    ],
    theme: {
        extend: {
            backgroundImage: {
                'auth-bg': "url('/src/assets/auth-bg.png')",
            },
            colors: {
                "vtd-primary": colors.indigo, // Light mode Datepicker color
                "vtd-secondary": colors.white, // Dark mode Datepicker color

            }
        },
    },
    plugins: [
        require('@tailwindcss/forms'),
    ],
}