import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import { mdsvex } from 'mdsvex';
import mdsvexOptions from './mdsvex.config.js';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  extensions: ['.svelte', '.md', '.svx'],
  kit: {
    adapter: adapter({
      // default options are shown. On some platforms
      // these options are set automatically — see below
      pages: 'build',
      assets: 'build',
      fallback: 'index.html',
      precompress: false,
      strict: true
    }),
    prerender: {
      handleHttpError: ({ path, referrer, message }) => {
        // ignore 404
        if (path === '/not-found' && referrer === '/') {
          return warn;
        }

        // otherwise fail
        throw new Error(message);
      },
    },
  },
  preprocess: [
    vitePreprocess({
      script: true,
    }),
    mdsvex(mdsvexOptions)
  ],
};

export default config;
