import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

const config = {
  kit: {
    adapter: adapter({
      pages: 'build',
      assets: 'build',
      fallback: 'index.html', // 👈 ensures SPA routing works
      precompress: false,
      strict: false
    })
  },
  preprocess: vitePreprocess()
};

export default config;
