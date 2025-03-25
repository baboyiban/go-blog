// @ts-check
import { defineConfig } from "astro/config";
import { config } from "dotenv";

config();

// https://astro.build/config
export default defineConfig({
  build: {
    assetsPrefix: `https://cdn.${process.env.DOMAIN}`,
  },
});
