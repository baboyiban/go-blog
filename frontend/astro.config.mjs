// @ts-check
import { defineConfig } from "astro/config";

// https://astro.build/config
export default defineConfig({
  build: {
    assetsPrefix: `https://cdn.${process.env.DOMAIN || "localhost"}`,
    inlineStylesheets: "auto",
  },
});
