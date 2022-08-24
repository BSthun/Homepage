import commonjs from '@rollup/plugin-commonjs';
import image from '@rollup/plugin-image';
import resolve from '@rollup/plugin-node-resolve';
import typescript from '@rollup/plugin-typescript';
import * as fs from 'fs';
import css from 'rollup-plugin-css-only';
import dev from 'rollup-plugin-dev';
import livereload from 'rollup-plugin-livereload';
import svelte from 'rollup-plugin-svelte';
import { terser } from 'rollup-plugin-terser';
import preprocess from 'svelte-preprocess';

const production = !process.env.ROLLUP_WATCH;

export default {
	input: 'src/main.ts',
	output: {
		sourcemap: true,
		format: 'iife',
		name: 'app',
		file: 'public/build/bundle.js',
	},
	plugins: [
		svelte({
			compilerOptions: {
				// * Enable run-time checks when not in production
				dev: !production,
			},
			preprocess: preprocess({
				sourceMap: !production,
				scss: {
					
					// * Removed due to import manually for code suggestions
					// prependData: `@import 'src/styles/_index.scss';`,
				},
			}),
			onwarn: (warning, handler) => {
				const { code, message } = warning;
				if (code === 'css-unused-selector') {
					return;
				}
				handler(warning);
			},
		}),
		
		// we'll extract any component CSS out into
		// a separate file - better for performance
		css({ output: 'bundle.css' }),
		
		// If you have external dependencies installed from
		// npm, you'll most likely need these plugins. In
		// some cases you'll need additional configuration -
		// consult the documentation for details:
		// https://github.com/rollup/plugins/tree/master/packages/commonjs
		resolve({
			browser: true,
			dedupe: ['svelte'],
		}),
		commonjs(),
		typescript({
			sourceMap: !production,
			inlineSources: !production,
		}),
		
		image(),
		
		// In dev mode, call `npm run start` once
		// the bundle has been generated
		!production && dev({
			dirs: ['public'],
			basePath: '/',
			silent: false,
			host: '0.0.0.0',
			port: 3000,
			spa: true,
			server: {
				https: {
					key: fs.readFileSync('./res/key.pem'),
					cert: fs.readFileSync('./res/cert.pem'),
				},
			},
			proxy: [
				{
					from: '/api',
					to: 'http://127.0.0.1:3001/api',
				},
			],
		}),
		
		// Watch the `public` directory and refresh the
		// browser on changes when not in production
		!production && livereload('public'),
		
		// If we're building for production (npm run build
		// instead of npm run dev), minify
		production && terser(),
	],
	watch: {
		clearScreen: false,
	},
};
