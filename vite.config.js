import { glob, globSync } from "glob";
import { defineConfig } from "vite";
import path from "path";
import { fileURLToPath } from "url";

export default defineConfig({
  root: "./web/views",
  build: {
    watch: {
      include: ["web/views/js/*.js", "web/views/pages/**/*.js"],
    },
    outDir: "../public/js",
    emptyOutDir: true,
    rollupOptions: {
      input: Object.fromEntries(
        globSync(["web/assets/js/*.js", "web/views/pages/**/*.js"]).map(
          (file) => [
            path.basename(file, path.extname(file)),
            fileURLToPath(new URL(file, import.meta.url)),
          ],
        ),
      ),
      output: {
        entryFileNames: "pages/[name]/[name].js",
      },
    },
  },
  server: {
    watch: {
      interval: 300,
    },
  },
});
