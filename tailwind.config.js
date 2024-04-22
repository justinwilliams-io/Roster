/** @type {import('tailwindcss').Config} */
module.exports = {
    mode: 'jit',
    content: ["./view/**/*.templ", "./node_modules/flowbite/**/*.js"],
    theme: {
        container: {
            center: true,
        }
    },
    extend: {},
    plugins: [
        require('flowbite/plugin')
    ],
}

