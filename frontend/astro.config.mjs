// @ts-check
import { defineConfig } from "astro/config";
import { config } from "dotenv";
import path from "path";
import { fileURLToPath } from "url";

config();

// https://astro.build/config
export default defineConfig({
  build: {
    assetsPrefix: `https://cdn.${process.env.DOMAIN || "localhost"}`,
    inlineStylesheets: "auto",
  },
});
