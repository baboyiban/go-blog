import { defineConfig } from "astro/config";

export default defineConfig({
  build: {
    assetsPrefix: process.env.DOMAIN
      ? `https://cdn.${process.env.DOMAIN}`
      : "/_astro", // ← 개발 환경 폴백
  },
});
