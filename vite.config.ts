import { defineConfig } from "vite";

export default defineConfig({
  root: "./assets",
  build: {
    outDir: "../public",
    rollupOptions: {
      manualChunks: undefined,
      input: {
        main: "./assets/main.ts",
      },
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
        assetFileNames: `assets/[name].[ext]`,
      },
    },
  },
});
