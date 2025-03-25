import { defineConfig } from "astro/config";
import { config } from "dotenv";

config();

export default defineConfig({
  build: {
    assetsPrefix: `https://cdn.${process.env.DOMAIN ?? "localhost"}`,
  },
});
