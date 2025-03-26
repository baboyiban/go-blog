import { defineConfig } from "astro/config";
import dotenv from "dotenv";

dotenv.config();

export default defineConfig({
  build: {
    assetsPrefix: process.env.DOMAIN
      ? `https://${process.env.DOMAIN}`
      : "/_astro", // ← 개발 환경 폴백
  },
});
