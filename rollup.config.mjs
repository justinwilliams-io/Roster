import { babel } from '@rollup/plugin-babel';

export default {
    input: 'src/js/input.js',
    output: {
        file: 'public/js/bundle.js',
        format: 'iife',
        compact: true,
        name: 'gametime'
    },
    plugins: [
        babel({
            babelHelpers: 'bundled'
        })
    ]
}
