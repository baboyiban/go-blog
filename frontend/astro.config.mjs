import { defineConfig } from "astro/config";
import { config } from "dotenv";

// Docker 환경 대응
try {
  config({ path: "/app/.env" }); // ← 절대 경로 지정
} catch (e) {
  console.log("Local .env loaded");
}

export default defineConfig({
  build: {
    assetsPrefix:
      process.env.NODE_ENV === "production"
        ? `https://cdn.${process.env.DOMAIN}`
        : undefined,
  },
});
