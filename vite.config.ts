// import prefresh from '@prefresh/vite';
import { resolve } from 'path';
import { UserConfig } from 'vite';

const root = './assets';

// const htmlRefreshPlugin = () => ({
//   name: 'html-refresh',
//   configureServer({ watcher, ws }) {
//     watcher.add(resolve(__dirname, 'views/**/*.html'));
//     watcher.on('change', function (path) {
//       if (path.endsWith('.html')) {
//         ws.send({
//           type: 'full-reload',
//         });
//       }
//     });
//   },
// });

const config: UserConfig = {
  esbuild: {
    jsxFactory: 'h',
    jsxFragment: 'Fragment',
  },
  build: {
    polyfillDynamicImport: false,
    manifest: true,
    outDir: '../public/assets',
    assetsDir: '',
    rollupOptions: {
      input: {
        app: resolve(__dirname, 'assets/app.ts'),
      },
    },
  },
  plugins: [],
  optimizeDeps: { include: ['preact/hooks'] },
  root,
};

module.exports = config;
